/**
 * Query key factories for stable, parameterized keys
 * Always use these factories for consistency in cache management
 */
export const queryKeys = {
  // Auth related queries
  auth: {
    all: () => ['auth'] as const,
    me: () => ['auth', 'me'] as const,
    profile: () => ['auth', 'profile'] as const,
  },

  // Users
  users: {
    all: () => ['users'] as const,
    list: (filters?: {page?: number; search?: string}) =>
      ['users', 'list', filters] as const,
    detail: (id: string) => ['users', 'detail', id] as const,
  },

  // Example: Orders (template for your features)
  orders: {
    all: () => ['orders'] as const,
    list: (filters?: {status?: string; page?: number}) =>
      ['orders', 'list', filters] as const,
    detail: (id: string) => ['orders', 'detail', id] as const,
  },

  // Example: Products (template for your features)
  products: {
    all: () => ['products'] as const,
    list: (categoryId?: string) => ['products', 'list', categoryId] as const,
    detail: (id: string) => ['products', 'detail', id] as const,
    search: (query: string) => ['products', 'search', query] as const,
  },

  // Settings
  settings: {
    all: () => ['settings'] as const,
    notifications: () => ['settings', 'notifications'] as const,
    preferences: () => ['settings', 'preferences'] as const,
  },
} as const;

export type QueryKeys = typeof queryKeys;
