import {AxiosError} from 'axios';
import Toast from 'react-native-simple-toast';
import {useAuthStore} from '../../stores/authStore';

/**
 * Unified error handler for both React Query and axios
 * Handles different HTTP status codes with appropriate user feedback
 */
export function handleApiError(error: unknown): void {
  if (error instanceof AxiosError) {
    const status = error.response?.status;

    switch (status) {
      case 401:
        // Unauthorized: clear session and redirect to login
        const clearSession = useAuthStore.getState().clearSession;
        clearSession();
        Toast.show('Session expired. Please log in again.', Toast.LONG);
        return;

      case 403:
        Toast.show('Access denied. Contact support.', Toast.LONG);
        return;

      case 404:
        Toast.show('Resource not found.', Toast.SHORT);
        return;

      case 422:
        // Validation error - typically handled at form level
        const validationMessage =
          error.response?.data?.message || 'Validation failed';
        Toast.show(validationMessage, Toast.SHORT);
        return;

      case 429:
        Toast.show('Too many requests. Please wait a moment.', Toast.LONG);
        return;

      case 500:
      case 502:
      case 503:
        Toast.show('Server error. Try again later.', Toast.LONG);
        return;

      default:
        // Generic error handling
        const message =
          error.response?.data?.message || error.message || 'Network error';
        Toast.show(message, Toast.SHORT);
        return;
    }
  }

  // Non-axios errors
  const message = (error as Error)?.message || 'Unexpected error';
  Toast.show(message, Toast.SHORT);
}

/**
 * Extract error message from an error object
 */
export function getErrorMessage(error: unknown): string {
  if (error instanceof AxiosError) {
    return (
      error.response?.data?.message || error.message || 'An error occurred'
    );
  }

  if (error instanceof Error) {
    return error.message;
  }

  return 'An unexpected error occurred';
}
