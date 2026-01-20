import axios, {AxiosError, InternalAxiosRequestConfig} from 'axios';
import Config from 'react-native-config';
import {handleApiError} from './error';
import {useAuthStore} from '../../stores/authStore';
import {API_CONFIG} from '../../constants';

/**
 * Create axios instance with mobile-optimized defaults
 */
export const api = axios.create({
  baseURL: Config.API_BASE_URL || 'https://api.example.com',
  timeout: API_CONFIG.TIMEOUT,
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json',
  },
});

/**
 * Request interceptor: add auth token to requests
 */
api.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
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

/**
 * Response interceptor: handle token refresh and errors
 */
api.interceptors.response.use(
  response => response,
  async error => {
    const originalRequest = error.config;

    // Handle 401 with token refresh
    if (error.response?.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;

      const refreshToken = useAuthStore.getState().tokens?.refreshToken;

      if (!refreshToken) {
        useAuthStore.getState().clearSession();
        return Promise.reject(error);
      }

      try {
        // Attempt to refresh the token
        const {data} = await axios.post(
          `${Config.API_BASE_URL || 'https://api.example.com'}/auth/refresh`,
          {refreshToken},
        );

        // Update tokens in store
        const currentUser = useAuthStore.getState().user;
        if (currentUser) {
          useAuthStore.getState().setSession({
            user: currentUser,
            tokens: data.tokens,
          });
        }

        // Retry the original request with new token
        originalRequest.headers.Authorization = `Bearer ${data.tokens.accessToken}`;
        return api(originalRequest);
      } catch {
        // Refresh failed - clear session
        useAuthStore.getState().clearSession();
        return Promise.reject(error);
      }
    }

    // Handle other errors
    handleApiError(error);
    return Promise.reject(error);
  },
);

export default api;
