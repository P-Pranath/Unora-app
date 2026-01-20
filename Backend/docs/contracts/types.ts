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
 * Error codes returned by the API
 * Use these for conditional error handling in FE
 */
export type ErrorCode =
  // Client errors (4xx)
  | 'VALIDATION_ERROR'    // 400 - Request validation failed
  | 'INVALID_REQUEST'     // 400 - Malformed request body
  | 'UNAUTHORIZED'        // 401 - Missing or invalid authentication
  | 'FORBIDDEN'           // 403 - Access denied
  | 'NOT_FOUND'           // 404 - Resource not found
  | 'CONFLICT'            // 409 - Resource conflict (e.g., already exists)
  | 'RATE_LIMITED'        // 429 - Too many requests / at capacity
  // Server errors (5xx)
  | 'INTERNAL_ERROR'      // 500 - Unexpected server error
  | 'DATABASE_ERROR'      // 500 - Database operation failed
  | 'EXTERNAL_SERVICE_ERROR' // 502 - Third-party service failed
  | 'SERVICE_UNAVAILABLE'; // 503 - Service temporarily unavailable

/**
 * Standard API error response
 */
export interface ApiError {
  /** Always false for error responses */
  success: false;
  /** Error details */
  error: {
    /** Error code - use for programmatic error handling */
    code: ErrorCode;
    /** Human-readable error message - display to user */
    message: string;
    /** Optional additional context (e.g., field validation errors) */
    details?: Record<string, unknown>;
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
