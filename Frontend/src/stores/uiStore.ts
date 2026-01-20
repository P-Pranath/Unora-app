import {create} from 'zustand';
import {persist, PersistStorage} from 'zustand/middleware';
import {mmkvStorage} from '../services/storage/mmkv';
import {STORAGE_KEYS} from '../constants';

/**
 * Theme options
 */
export type ThemeMode = 'light' | 'dark' | 'system';

/**
 * UI state interface
 */
export type UiState = {
  // Theme
  theme: ThemeMode;

  // Onboarding
  isOnboarded: boolean;

  // Selected tenant/workspace (if applicable)
  selectedTenantId: string | null;

  // Actions
  setTheme: (theme: ThemeMode) => void;
  setOnboarded: (value: boolean) => void;
  setTenant: (id: string | null) => void;
  resetUiState: () => void;
};

const initialState = {
  theme: 'system' as ThemeMode,
  isOnboarded: false,
  selectedTenantId: null,
};

/**
 * UI store for app-wide UI state with MMKV persistence
 *
 * Usage:
 * - const theme = useUiStore(state => state.theme)
 * - const setTheme = useUiStore(state => state.setTheme)
 */
export const useUiStore = create<UiState>()(
  persist(
    set => ({
      ...initialState,

      setTheme: theme => set({theme}),

      setOnboarded: value => set({isOnboarded: value}),

      setTenant: id => set({selectedTenantId: id}),

      resetUiState: () => set(initialState),
    }),
    {
      name: STORAGE_KEYS.UI,
      storage: mmkvStorage as unknown as PersistStorage<UiState>,
    },
  ),
);

export default useUiStore;
