import {create} from 'zustand';
import {persist, PersistStorage} from 'zustand/middleware';
import {mmkvStorage} from '../services/storage/mmkv';
import {queryClient} from '../services/http/queryClient';
import {STORAGE_KEYS} from '../constants';

/**
 * Auth user type
 */
export type AuthUser = {
  id: string;
  email: string;
  firstName?: string;
  lastName?: string;
  avatar?: string;
};

/**
 * Auth tokens type
 */
export type AuthTokens = {
  accessToken: string;
  refreshToken?: string;
  expiresAt?: number;
};

/**
 * Auth state interface
 */
export type AuthState = {
  user: AuthUser | null;
  tokens: AuthTokens | null;
  isAuthenticated: boolean;
  isLoading: boolean;

  // Actions
  setSession: (payload: {user: AuthUser; tokens: AuthTokens}) => void;
  setUser: (user: AuthUser) => void;
  setTokens: (tokens: AuthTokens) => void;
  setLoading: (loading: boolean) => void;
  clearSession: () => void;
};

/**
 * Persisted state type (excludes actions and transient state)
 */
type PersistedAuthState = Pick<
  AuthState,
  'user' | 'tokens' | 'isAuthenticated'
>;

/**
 * Auth store with MMKV persistence
 *
 * Usage:
 * - After login: authStore.getState().setSession({user, tokens})
 * - On logout: authStore.getState().clearSession()
 * - In components: const user = useAuthStore(state => state.user)
 */
export const useAuthStore = create<AuthState>()(
  persist(
    set => ({
      user: null,
      tokens: null,
      isAuthenticated: false,
      isLoading: false,

      setSession: ({user, tokens}) => {
        set({
          user,
          tokens,
          isAuthenticated: true,
          isLoading: false,
        });
      },

      setUser: user => {
        set({user});
      },

      setTokens: tokens => {
        set({tokens});
      },

      setLoading: isLoading => {
        set({isLoading});
      },

      clearSession: () => {
        set({
          user: null,
          tokens: null,
          isAuthenticated: false,
          isLoading: false,
        });
        // Clear React Query cache on logout to prevent data leakage
        queryClient.clear();
      },
    }),
    {
      name: STORAGE_KEYS.AUTH,
      storage: mmkvStorage as unknown as PersistStorage<PersistedAuthState>,
      // Only persist user and tokens (exclude actions and loading state)
      partialize: (state): PersistedAuthState => ({
        user: state.user,
        tokens: state.tokens,
        isAuthenticated: state.isAuthenticated,
      }),
    },
  ),
);

export default useAuthStore;
