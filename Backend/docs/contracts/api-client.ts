/**
 * Unora API Client Template
 * Copy and adapt for your mobile app (React Native, Flutter, etc.)
 */

import {
  API_BASE_URL,
  API_ENDPOINTS,
  ApiResponse,
  ApiError,
  AuthResponse,
  GoogleOAuthRequest,
  RefreshTokenRequest,
  RefreshTokenResponse,
  TokenPayload,
} from './types';

// Storage keys - adapt for your platform (AsyncStorage, SecureStore, etc.)
const TOKEN_STORAGE_KEY = 'auth_tokens';

interface StoredTokens {
  accessToken: string;
  refreshToken: string;
  expiresAt: number;
}

/**
 * Base API client class
 * Extend or adapt for your specific framework
 */
class ApiClient {
  private baseUrl: string;
  private tokens: StoredTokens | null = null;

  constructor(baseUrl = API_BASE_URL) {
    this.baseUrl = baseUrl;
  }

  /**
   * Make authenticated API request with automatic token refresh
   */
  async request<T>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<ApiResponse<T> | ApiError> {
    const headers: Record<string, string> = {
      'Content-Type': 'application/json',
      ...(options.headers as Record<string, string>),
    };

    // Add auth header if we have tokens
    if (this.tokens?.accessToken) {
      headers['Authorization'] = `Bearer ${this.tokens.accessToken}`;
    }

    try {
      const response = await fetch(`${this.baseUrl}${endpoint}`, {
        ...options,
        headers,
      });

      const data = await response.json();

      // If 401 and we have refresh token, try to refresh
      if (response.status === 401 && this.tokens?.refreshToken) {
        const refreshed = await this.refreshAccessToken();
        if (refreshed) {
          // Retry original request with new token
          headers['Authorization'] = `Bearer ${this.tokens!.accessToken}`;
          const retryResponse = await fetch(`${this.baseUrl}${endpoint}`, {
            ...options,
            headers,
          });
          return retryResponse.json();
        }
      }

      return data;
    } catch (error) {
      return {
        success: false,
        error: {
          code: 'NETWORK_ERROR',
          message: error instanceof Error ? error.message : 'Network request failed',
        },
      };
    }
  }

  /**
   * Login with Google ID token
   */
  async loginWithGoogle(idToken: string): Promise<ApiResponse<AuthResponse> | ApiError> {
    const result = await this.request<AuthResponse>(API_ENDPOINTS.auth.google, {
      method: 'POST',
      body: JSON.stringify({ idToken } as GoogleOAuthRequest),
    });

    if (result.success) {
      // Store tokens
      this.tokens = {
        accessToken: result.data.accessToken,
        refreshToken: result.data.refreshToken,
        expiresAt: Date.now() + result.data.expiresIn * 1000,
      };
      // TODO: Persist to secure storage
    }

    return result;
  }

  /**
   * Refresh access token
   */
  async refreshAccessToken(): Promise<boolean> {
    if (!this.tokens?.refreshToken) return false;

    const result = await this.request<RefreshTokenResponse>(API_ENDPOINTS.auth.refresh, {
      method: 'POST',
      body: JSON.stringify({ refreshToken: this.tokens.refreshToken } as RefreshTokenRequest),
    });

    if (result.success) {
      this.tokens = {
        ...this.tokens,
        accessToken: result.data.accessToken,
        expiresAt: Date.now() + result.data.expiresIn * 1000,
      };
      return true;
    }

    // Refresh failed, clear tokens
    this.tokens = null;
    return false;
  }

  /**
   * Logout
   */
  async logout(): Promise<void> {
    if (this.tokens) {
      await this.request(API_ENDPOINTS.auth.logout, { method: 'POST' });
    }
    this.tokens = null;
    // TODO: Clear from secure storage
  }

  /**
   * Get current user info
   */
  async getCurrentUser(): Promise<ApiResponse<TokenPayload> | ApiError> {
    return this.request<TokenPayload>(API_ENDPOINTS.auth.me);
  }

  /**
   * Check if user is authenticated
   */
  isAuthenticated(): boolean {
    return !!this.tokens?.accessToken;
  }
}

// Export singleton instance
export const apiClient = new ApiClient();
export default apiClient;
