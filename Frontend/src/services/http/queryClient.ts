import {QueryCache, QueryClient} from '@tanstack/react-query';
import {handleApiError} from './error';

/**
 * Centralized React Query client with global defaults
 *
 * Mobile-specific configuration:
 * - retry: 1 - Avoids aggressive retries on flaky mobile networks
 * - staleTime: 5 minutes - Keeps data fresh for reasonable time
 * - gcTime: 10 minutes - Garbage collection time (formerly cacheTime)
 * - refetchOnReconnect: true - Important for mobile connectivity
 */
export const queryClient = new QueryClient({
  queryCache: new QueryCache({
    onError: handleApiError,
  }),
  defaultOptions: {
    queries: {
      retry: 1,
      staleTime: 5 * 60 * 1000, // 5 minutes
      gcTime: 10 * 60 * 1000, // 10 minutes
      refetchOnWindowFocus: false,
      refetchOnReconnect: true,
    },
    mutations: {
      retry: 0,
    },
  },
});

export default queryClient;
