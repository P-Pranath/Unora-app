#!/bin/bash

# Generate TypeScript types from Swagger spec
# Works with both Swagger 2.0 and OpenAPI 3.0

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

SWAGGER_FILE="docs/swagger/swagger.json"
OUTPUT_DIR="docs/contracts"

echo -e "${BLUE}=== Generating TypeScript Types ===${NC}"

# Check if swagger file exists
if [ ! -f "$SWAGGER_FILE" ]; then
    echo -e "${RED}Error: $SWAGGER_FILE not found${NC}"
    echo -e "${YELLOW}Run 'make swagger' first to generate Swagger docs${NC}"
    exit 1
fi

# Create output directory
mkdir -p "$OUTPUT_DIR"

# Generate comprehensive TypeScript types file
cat > "$OUTPUT_DIR/types.ts" << 'EOF'
/**
 * Unora API TypeScript Types
 * Auto-generated from Swagger spec - DO NOT EDIT MANUALLY
 * Run 'make types' to regenerate
 * 
 * Usage:
 * import { AuthResponse, GoogleOAuthRequest, ApiResponse } from '@/api/types';
 */

// ============================================
// Base Configuration
// ============================================

export const API_BASE_URL = process.env.API_URL || 'http://localhost:8080';
export const API_VERSION = 'v1';

// ============================================
// Auth Types
// ============================================

/**
 * Request body for Google OAuth login
 * Send the ID token received from Google Sign-In SDK
 */
export interface GoogleOAuthRequest {
  /** Google ID token obtained from Google Sign-In SDK on mobile */
  idToken: string;
}

/**
 * Successful authentication response
 * Contains tokens and user information
 */
export interface AuthResponse {
  /** JWT access token - use in Authorization header for protected routes */
  accessToken: string;
  /** Refresh token - use to get new access token when it expires */
  refreshToken: string;
  /** Access token expiry time in seconds (typically 86400 = 24 hours) */
  expiresIn: number;
  /** Token type - always "Bearer" */
  tokenType: 'Bearer';
  /** Authenticated user information */
  user: UserResponse;
}

/**
 * User profile information
 */
export interface UserResponse {
  /** Unique user identifier (UUID) */
  id: string;
  /** User's email address */
  email: string;
  /** User's full name */
  name: string;
  /** User's first name */
  firstName?: string;
  /** User's last name */
  lastName?: string;
  /** URL to user's profile picture */
  picture?: string;
  /** Authentication provider ('google', 'apple') */
  provider: 'google' | 'apple';
  /** ISO timestamp of account creation */
  createdAt: string;
}

/**
 * Token refresh request
 */
export interface RefreshTokenRequest {
  /** Refresh token received from login response */
  refreshToken: string;
}

/**
 * Token refresh response
 */
export interface RefreshTokenResponse {
  /** New JWT access token */
  accessToken: string;
  /** Token expiry time in seconds */
  expiresIn: number;
  /** Token type - always "Bearer" */
  tokenType: 'Bearer';
}

/**
 * JWT token payload (decoded from access token)
 */
export interface TokenPayload {
  /** User's unique identifier */
  userId: string;
  /** User's email address */
  email: string;
  /** Authentication provider */
  provider: 'google' | 'apple';
}

// ============================================
// Storage Types
// ============================================

/**
 * Request for presigned upload URL
 */
export interface PresignedUploadRequest {
  /** Name of the file to upload */
  fileName: string;
  /** MIME type of the file (e.g., 'image/jpeg', 'application/pdf') */
  contentType: string;
}

/**
 * Response with presigned upload URL
 */
export interface PresignedUploadResponse {
  /** Presigned URL to upload the file to (PUT request) */
  presignedUrl: string;
  /** Filename as stored */
  fileName: string;
}

/**
 * Response with presigned download/preview URL
 */
export interface PresignedUrlResponse {
  /** Presigned URL to download/preview the file */
  presignedUrl: string;
  /** Filename */
  fileName: string;
}

// ============================================
// API Response Types
// ============================================

/**
 * Standard API success response wrapper
 */
export interface ApiResponse<T = unknown> {
  /** Always true for successful responses */
  success: true;
  /** Response payload */
  data: T;
}

/**
 * Standard API error response
 */
export interface ApiError {
  /** Always false for error responses */
  success: false;
  /** Error details */
  error: {
    /** Error code (e.g., 'UNAUTHORIZED', 'NOT_FOUND') */
    code: string;
    /** Human-readable error message */
    message: string;
  };
}

/**
 * Paginated response wrapper
 */
export interface PaginatedResponse<T> {
  success: true;
  data: T[];
  pagination: Pagination;
}

/**
 * Pagination metadata
 */
export interface Pagination {
  /** Current page number (1-indexed) */
  page: number;
  /** Items per page */
  limit: number;
  /** Total number of items */
  total: number;
  /** Total number of pages */
  totalPages: number;
}

// ============================================
// API Endpoints
// ============================================

/**
 * API endpoint paths
 */
export const API_ENDPOINTS = {
  auth: {
    /** POST - Login with Google ID token */
    google: '/api/v1/auth/google',
    /** POST - Refresh access token */
    refresh: '/api/v1/auth/refresh',
    /** POST - Logout (requires auth) */
    logout: '/api/v1/auth/logout',
    /** GET - Get current user info (requires auth) */
    me: '/api/v1/auth/me',
  },
  storage: {
    /** POST - Get presigned URL for file upload */
    presignedUpload: '/api/v1/storage/presigned-upload',
    /** GET - Get presigned URL for file download */
    presignedDownload: (fileName: string) => `/api/v1/storage/presigned-download/${encodeURIComponent(fileName)}`,
    /** GET - Get presigned URL for file preview */
    presignedPreview: (fileName: string) => `/api/v1/storage/presigned-preview/${encodeURIComponent(fileName)}`,
  },
} as const;

// ============================================
// HTTP Request Helper Types
// ============================================

/**
 * HTTP request configuration
 */
export interface RequestConfig {
  /** Request headers */
  headers?: Record<string, string>;
  /** Query parameters */
  params?: Record<string, string | number | boolean>;
  /** Request timeout in milliseconds */
  timeout?: number;
}

/**
 * Authenticated request configuration
 */
export interface AuthenticatedRequestConfig extends RequestConfig {
  /** Access token for Authorization header */
  accessToken: string;
}

// ============================================
// Utility Types
// ============================================

/** Union type for all possible API responses */
export type ApiResult<T> = ApiResponse<T> | ApiError;

/** Extract success data type from ApiResponse */
export type ExtractData<T> = T extends ApiResponse<infer D> ? D : never;

/** Make specific keys optional */
export type PartialBy<T, K extends keyof T> = Omit<T, K> & Partial<Pick<T, K>>;

/** Make specific keys required */
export type RequiredBy<T, K extends keyof T> = Omit<T, K> & Required<Pick<T, K>>;

/** Nullable type helper */
export type Nullable<T> = T | null;
EOF

echo -e "${GREEN}✓ TypeScript types generated at: $OUTPUT_DIR/types.ts${NC}"

# Generate API client template
cat > "$OUTPUT_DIR/api-client.ts" << 'EOF'
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
EOF

echo -e "${GREEN}✓ API client template generated at: $OUTPUT_DIR/api-client.ts${NC}"
echo ""
echo -e "${YELLOW}Copy the docs/contracts/ folder to your mobile app project.${NC}"
