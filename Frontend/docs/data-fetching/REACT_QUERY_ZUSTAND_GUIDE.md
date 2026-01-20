# React Query + Zustand Integration Guide for React Native

This guide adapts the React Query + Zustand separation-of-concerns patterns for **React Native mobile apps**, including MMKV persistence, axios interceptors, and mobile-specific considerations.

## Core Principles

- **One source of truth per kind of state:**
  - **Server-derived data** (fetching, caching, revalidation, retries, background refresh) → React Query.
  - **App/session/UI state** (auth tokens, user selections, feature flags, theme, modal states) → Zustand with MMKV persistence.
- **Do not duplicate server data in Zustand:**
  - If non-React code (e.g., axios interceptors, push notification handlers) needs a synchronous snapshot, store minimal projections only (e.g., `userId`, `accessToken`); keep the canonical dataset in React Query.
- **Centralize cross-cutting concerns:**
  - One axios instance with interceptors.
  - One QueryClient instance shared everywhere.
  - One error handler used by both React Query and axios.
- **Prefer invalidation over manual cache writes** unless UX needs optimistic updates.

---

## Decision Guide: Where Should This State Live?

Use this checklist to decide between React Query and Zustand:

| Question                                                                                              | Answer                                                                                                                                                                                                                   |
| ----------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Can this be fetched/refetched from the backend?                                                       | **React Query**                                                                                                                                                                                                          |
| Does this need to be available synchronously everywhere (including non-React code like interceptors)? | **Zustand** (store minimal identifiers only)                                                                                                                                                                             |
| Is this ephemeral, component-local UI (open/closed, input text)?                                      | **Local state** (`useState`, `useReducer`)                                                                                                                                                                               |
| Is this cross-screen UI flag or selection (theme, selected tenant, modal state)?                      | **Zustand** (optionally with MMKV `persist`)                                                                                                                                                                             |
| Is this derived from server data (e.g., ability checks, permissions)?                                 | Derive from **React Query** in components; if needed globally or outside React, store a minimal denormalized snapshot in Zustand (e.g., `string[]` of permission slugs), but keep the server data canonical in the cache |

---

## High-Level Mobile Architecture

```
┌─────────────────────────────────────────────────────────────┐
│  React Native App (SafeAreaProvider, NavigationContainer)  │
├─────────────────────────────────────────────────────────────┤
│  QueryClientProvider (React Query)                          │
│    ├─ queries: server data (users, orders, products)        │
│    ├─ mutations: create/update/delete                       │
│    └─ invalidation after mutations                          │
├─────────────────────────────────────────────────────────────┤
│  Zustand Stores (MMKV-persisted)                            │
│    ├─ authStore: { token, userId, refreshToken }           │
│    ├─ uiStore: { theme, selectedTenant, isOnboarded }      │
│    └─ settingsStore: { pushNotificationsEnabled, ... }     │
├─────────────────────────────────────────────────────────────┤
│  Axios Instance                                              │
│    ├─ Request interceptor: reads token from authStore      │
│    ├─ Response interceptor: handles 401 (token refresh)    │
│    └─ Error handler: unified error toasts/alerts           │
└─────────────────────────────────────────────────────────────┘
```

---

## Reference Implementation (Copy into Your Project)

### 1. `src/services/http/queryClient.ts`

Central QueryClient with global defaults and unified error handling.

```ts
import {QueryCache, QueryClient} from '@tanstack/react-query';
import {handleApiError} from './error';

export const queryClient = new QueryClient({
  queryCache: new QueryCache({
    onError: handleApiError,
  }),
  defaultOptions: {
    queries: {
      retry: 1, // Retry once on mobile (avoid excessive retries on flaky networks)
      staleTime: 5 * 60 * 1000, // 5 minutes
      gcTime: 10 * 60 * 1000, // 10 minutes (formerly cacheTime in v4)
      refetchOnWindowFocus: false,
      refetchOnReconnect: true, // Important for mobile
    },
    mutations: {
      retry: 0,
    },
  },
});

export default queryClient;
```

**Mobile-specific notes:**

- `retry: 1` – avoids aggressive retries on slow/flaky mobile networks.
- `refetchOnReconnect: true` – refetches stale queries when network reconnects.
- `staleTime` and `gcTime` tuned for mobile (shorter than web to avoid stale data lingering).

---

### 2. `src/services/http/client.ts`

Axios instance with interceptors reading from the auth store.

```ts
import axios, {AxiosError, InternalAxiosRequestConfig} from 'axios';
import Config from 'react-native-config';
import {handleApiError} from './error';
import {useAuthStore} from '../../features/auth/store/authStore';

export const api = axios.create({
  baseURL: Config.API_BASE_URL || 'https://api.example.com',
  timeout: 15000, // 15s timeout for mobile
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor: add auth token
api.interceptors.request.use(
  async (config: InternalAxiosRequestConfig) => {
    const tokens = useAuthStore.getState().tokens;

    if (tokens?.accessToken) {
      config.headers.Authorization = `Bearer ${tokens.accessToken}`;
    }

    return config;
  },
  (error: AxiosError) => {
    return Promise.reject(error);
  },
);

// Response interceptor: handle errors
api.interceptors.response.use(
  response => response,
  (error: AxiosError) => {
    handleApiError(error);
    return Promise.reject(error);
  },
);

export default api;
```

**Mobile-specific notes:**

- `react-native-config` for environment variables (vs `process.env` on web).
- Shorter `timeout` (15s) to avoid long waits on poor connections.
- Token is read synchronously via `useAuthStore.getState()` (works in interceptors, which are not React hooks).

---

### 3. `src/services/http/error.ts`

Unified error handler for both React Query and axios.

```ts
import {AxiosError} from 'axios';
import Toast from 'react-native-simple-toast';
import {useAuthStore} from '../../features/auth/store/authStore';

export function handleApiError(error: unknown) {
  if (error instanceof AxiosError) {
    const status = error.response?.status;

    if (status === 401) {
      // Unauthorized: clear session and redirect to login
      const clearSession = useAuthStore.getState().clearSession;
      clearSession();
      Toast.show('Session expired. Please log in again.', Toast.LONG);
      return;
    }

    if (status === 403) {
      Toast.show('Access denied. Contact support.', Toast.LONG);
      return;
    }

    if (status === 500) {
      Toast.show('Server error. Try again later.', Toast.LONG);
      return;
    }

    const message =
      error.response?.data?.message || error.message || 'Network error';
    Toast.show(message, Toast.SHORT);
    return;
  }

  // Non-axios errors
  const message = (error as Error)?.message || 'Unexpected error';
  Toast.show(message, Toast.SHORT);
}
```

**Mobile-specific notes:**

- Uses `react-native-simple-toast` (or your preferred toast/alert library) instead of web toasts.
- On 401, calls `clearSession` from the auth store (clears MMKV-persisted tokens).

---

### 4. `src/services/http/queryKeys.ts`

Query key factories for stable, parameterized keys.

```ts
export const queryKeys = {
  auth: {
    me: () => ['auth', 'me'] as const,
  },
  orders: {
    all: () => ['orders'] as const,
    list: (filters: {status?: string; page?: number}) =>
      ['orders', 'list', filters] as const,
    detail: (id: string) => ['orders', 'detail', id] as const,
  },
  products: {
    all: () => ['products'] as const,
    list: (categoryId?: string) => ['products', 'list', categoryId] as const,
    detail: (id: string) => ['products', 'detail', id] as const,
  },
  users: {
    all: () => ['users'] as const,
    list: () => ['users', 'list'] as const,
    detail: (id: string) => ['users', 'detail', id] as const,
  },
};
```

**Usage:**

- Always use these factories for consistency.
- Makes invalidation precise and predictable.

---

### 5. `src/features/auth/store/authStore.ts`

Zustand auth store with MMKV persistence (minimal identifiers only).

```ts
import {create} from 'zustand';
import {persist} from 'zustand/middleware';
import {mmkvStorage} from '../../../services/storage/mmkv';
import queryClient from '../../../services/http/queryClient';
import type {AuthUser, AuthTokens} from '../types';

export type AuthState = {
  user: AuthUser | null;
  tokens: AuthTokens | null;
  setSession: (payload: {user: AuthUser; tokens: AuthTokens}) => void;
  clearSession: () => void;
};

export const useAuthStore = create<AuthState>()(
  persist(
    (set, get) => ({
      user: null,
      tokens: null,

      setSession: ({user, tokens}) => {
        set({user, tokens});
      },

      clearSession: () => {
        set({user: null, tokens: null});
        queryClient.clear(); // Clear React Query cache on logout
      },
    }),
    {
      name: 'auth', // MMKV key
      storage: mmkvStorage,
      partialize: state => ({
        user: state.user,
        tokens: state.tokens,
      }),
    },
  ),
);
```

**Mobile-specific notes:**

- Uses `mmkvStorage` (see `src/services/storage/mmkv.ts` from earlier).
- On `clearSession`, also calls `queryClient.clear()` to remove cached server data (prevents leaking data across sessions).

---

### 6. `src/stores/uiStore.ts`

UI/preference store with MMKV persistence.

```ts
import {create} from 'zustand';
import {persist} from 'zustand/middleware';
import {mmkvStorage} from '../services/storage/mmkv';

export type UiState = {
  theme: 'light' | 'dark' | 'system';
  isOnboarded: boolean;
  selectedTenantId: string | null;
  setTheme: (theme: 'light' | 'dark' | 'system') => void;
  setOnboarded: (value: boolean) => void;
  setTenant: (id: string | null) => void;
};

export const useUiStore = create<UiState>()(
  persist(
    set => ({
      theme: 'system',
      isOnboarded: false,
      selectedTenantId: null,
      setTheme: theme => set({theme}),
      setOnboarded: value => set({isOnboarded: value}),
      setTenant: id => set({selectedTenantId: id}),
    }),
    {
      name: 'ui',
      storage: mmkvStorage,
    },
  ),
);
```

---

### 7. Example React Query Hooks

#### `src/features/orders/api/useOrdersQuery.ts`

```ts
import {useQuery} from '@tanstack/react-query';
import {api} from '../../../services/http/client';
import {queryKeys} from '../../../services/http/queryKeys';
import type {Order} from '../types';

export type OrderFilters = {
  status?: string;
  page?: number;
};

async function fetchOrders(filters: OrderFilters): Promise<Order[]> {
  const {data} = await api.get('/orders', {params: filters});
  return data;
}

export const useOrdersQuery = (filters: OrderFilters) => {
  return useQuery({
    queryKey: queryKeys.orders.list(filters),
    queryFn: () => fetchOrders(filters),
  });
};
```

#### `src/features/orders/api/useCreateOrderMutation.ts`

```ts
import {useMutation, useQueryClient} from '@tanstack/react-query';
import Toast from 'react-native-simple-toast';
import {api} from '../../../services/http/client';
import {queryKeys} from '../../../services/http/queryKeys';
import type {Order} from '../types';

async function createOrder(payload: Partial<Order>): Promise<Order> {
  const {data} = await api.post('/orders', payload);
  return data;
}

export const useCreateOrderMutation = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: createOrder,
    onSuccess: () => {
      queryClient.invalidateQueries({queryKey: queryKeys.orders.all()});
      Toast.show('Order created successfully', Toast.SHORT);
    },
  });
};
```

---

### 8. Auth Flow: Login + Bootstrap

Combines Zustand (for token) and React Query (for priming user data).

#### `src/features/auth/api/useLoginMutation.ts`

```ts
import {useMutation} from '@tanstack/react-query';
import Toast from 'react-native-simple-toast';
import {api} from '../../../services/http/client';
import queryClient from '../../../services/http/queryClient';
import {queryKeys} from '../../../services/http/queryKeys';
import {useAuthStore} from '../store/authStore';
import type {AuthUser, AuthTokens} from '../types';

type LoginPayload = {
  email: string;
  password: string;
};

type LoginResponse = {
  user: AuthUser;
  tokens: AuthTokens;
};

async function login(payload: LoginPayload): Promise<LoginResponse> {
  const {data} = await api.post('/auth/login', payload);
  return data;
}

export const useLoginMutation = () => {
  const setSession = useAuthStore(state => state.setSession);

  return useMutation({
    mutationFn: login,
    onSuccess: async ({user, tokens}) => {
      // 1) Write minimal identifiers to Zustand (for axios interceptors)
      setSession({user, tokens});

      // 2) Prime critical server data into React Query cache
      await queryClient.prefetchQuery({
        queryKey: queryKeys.auth.me(),
        queryFn: async () => {
          const {data} = await api.get('/auth/me');
          return data as AuthUser;
        },
      });

      Toast.show('Logged in successfully', Toast.SHORT);
    },
    onError: () => {
      Toast.show('Login failed. Check credentials.', Toast.LONG);
    },
  });
};
```

**Flow:**

1. Call backend login endpoint.
2. Write `user` + `tokens` to Zustand (persisted to MMKV).
3. Prefetch user profile (`/auth/me`) into React Query cache so it's available immediately on authenticated screens.

---

## Interaction Patterns

### 1. Post-Auth Bootstrap

```ts
// In useLoginMutation or app bootstrap logic
setSession({user, tokens});

await queryClient.prefetchQuery({
  queryKey: queryKeys.auth.me(),
  queryFn: () => fetchMe(),
});

// Optionally mirror minimal permissions snapshot if needed by non-React code
// e.g., store permissionSlugs: string[] in a separate permissionsStore
```

### 2. List/Detail Pages

- Use React Query for list and detail queries.
- Derive query keys from route params and filters.
- On create/update/delete, invalidate the smallest affected keys.
- Keep table filters in local state; elevate to Zustand only if filters are cross-screen or need persistence.

**Example: Orders list screen**

```tsx
const {data, isLoading} = useOrdersQuery({status: 'pending', page: 1});
const {mutateAsync: createOrder} = useCreateOrderMutation();

// On create, mutation auto-invalidates queryKeys.orders.all()
```

### 3. Feature Flags & Permissions

- Treat as **server state** (React Query).
- If checks are needed outside React (e.g., in navigation guards or push notification handlers), mirror a minimal, denormalized snapshot in Zustand (e.g., `string[]` of permission slugs).

**Example:**

```ts
// Query for full permissions object
const {data: permissions} = useQuery({
  queryKey: ['permissions'],
  queryFn: fetchPermissions,
});

// In your app bootstrap or onSuccess, optionally store minimal snapshot
usePermissionsStore.getState().setPermissionSlugs(permissions.map(p => p.slug));
```

### 4. UI State

- Use Zustand for cross-screen UI flags (modal visibility, theme, onboarding status).
- Component-local UI stays local (`useState`).

**Example:**

```ts
const theme = useUiStore(state => state.theme);
const setTheme = useUiStore(state => state.setTheme);
```

### 5. Error Handling

- Centralize to a single handler used by both React Query (`QueryCache.onError`) and axios interceptors.
- Use mobile-friendly toasts or alerts based on severity.

---

## Mobile-Specific Considerations

### 1. Network State Awareness

React Query automatically refetches on reconnect (`refetchOnReconnect: true`), but you can enhance this with manual network monitoring:

```ts
import NetInfo from '@react-native-community/netinfo';
import {onlineManager} from '@tanstack/react-query';

NetInfo.addEventListener(state => {
  onlineManager.setOnline(state.isConnected ?? false);
});
```

### 2. Background App State

When the app returns to foreground, React Query refetches stale queries. You can customize this:

```ts
import {AppState} from 'react-native';
import {focusManager} from '@tanstack/react-query';

AppState.addEventListener('change', state => {
  focusManager.setFocused(state === 'active');
});
```

### 3. Optimistic Updates for Better UX

On mobile, network can be slow. Use optimistic updates for immediate feedback:

```ts
const {mutate: likePost} = useMutation({
  mutationFn: likePostApi,
  onMutate: async postId => {
    await queryClient.cancelQueries({queryKey: ['posts', postId]});
    const previous = queryClient.getQueryData(['posts', postId]);

    queryClient.setQueryData(['posts', postId], (old: any) => ({
      ...old,
      isLiked: true,
      likeCount: old.likeCount + 1,
    }));

    return {previous};
  },
  onError: (_err, _postId, context) => {
    queryClient.setQueryData(['posts', _postId], context?.previous);
  },
  onSettled: postId => {
    queryClient.invalidateQueries({queryKey: ['posts', postId]});
  },
});
```

### 4. MMKV vs AsyncStorage

- **MMKV** is synchronous and faster than AsyncStorage.
- Use MMKV for all Zustand persistence (auth, ui, settings).
- Avoid AsyncStorage unless you have a specific reason (e.g., existing legacy code).

### 5. Token Refresh Flow

If your API uses refresh tokens, implement token refresh in the axios response interceptor:

```ts
api.interceptors.response.use(
  response => response,
  async error => {
    const originalRequest = error.config;

    if (error.response?.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;

      const refreshToken = useAuthStore.getState().tokens?.refreshToken;
      if (!refreshToken) {
        useAuthStore.getState().clearSession();
        return Promise.reject(error);
      }

      try {
        const {data} = await axios.post('/auth/refresh', {refreshToken});
        useAuthStore.getState().setSession({
          user: useAuthStore.getState().user!,
          tokens: data.tokens,
        });

        originalRequest.headers.Authorization = `Bearer ${data.tokens.accessToken}`;
        return api(originalRequest);
      } catch {
        useAuthStore.getState().clearSession();
        return Promise.reject(error);
      }
    }

    handleApiError(error);
    return Promise.reject(error);
  },
);
```

---

## Anti-Patterns to Avoid

- **Duplicating entire server responses in Zustand** (two sources of truth).
- **Calling `useQueryClient()` inside stores or non-React modules**; instead, import the shared `queryClient`.
- **Using React Query for purely local UI state** (use local `useState` or Zustand).
- **Using Zustand to cache lists from the server** (reinventing caching/invalidations).
- **Storing large objects in MMKV** (MMKV is optimized for small key-value pairs; large objects can degrade performance).

---

## TL;DR

- **React Query** for server state; **Zustand (+ MMKV)** for app/session/UI state.
- Keep each kind of state in a single place; only mirror minimal projections when absolutely necessary.
- Centralize client, axios, and error handling.
- Invalidate after mutations; add optimistic updates only when it improves UX.
- On mobile: tune retry/staleTime, handle network reconnects, monitor app state, and use MMKV for persistence.

---

## Next Steps

1. Install dependencies (already done in this project):

   - `@tanstack/react-query`
   - `zustand`
   - `react-native-mmkv`
   - `axios`
   - `react-native-config` (for env vars)
   - `@react-native-community/netinfo` (for network monitoring)
   - `react-native-simple-toast` (or your preferred toast library)

2. Copy the reference implementations above into your project.

3. Build your first feature following the patterns:

   - Create query/mutation hooks in `features/<feature>/api`.
   - Use Zustand stores for minimal session identifiers.
   - Use React Query for all server data.

4. Test the full flow:

   - Login → token persisted to MMKV.
   - Fetch data → cached in React Query.
   - Mutation → invalidates relevant queries.
   - Logout → clears both Zustand and React Query cache.

5. Refer to this guide when adding new features or refactoring existing code.
