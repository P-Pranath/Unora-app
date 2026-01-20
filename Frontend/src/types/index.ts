/**
 * Common utility types
 */

/**
 * Make specific keys optional
 */
export type PartialBy<T, K extends keyof T> = Omit<T, K> & Partial<Pick<T, K>>;

/**
 * Make specific keys required
 */
export type RequiredBy<T, K extends keyof T> = Omit<T, K> &
  Required<Pick<T, K>>;

/**
 * Extract keys of type from object
 */
export type KeysOfType<T, U> = {
  [K in keyof T]: T[K] extends U ? K : never;
}[keyof T];

/**
 * Nullable type
 */
export type Nullable<T> = T | null;

/**
 * API response wrapper
 */
export type ApiResponse<T> = {
  data: T;
  message?: string;
  success: boolean;
};

/**
 * Paginated API response
 */
export type PaginatedResponse<T> = {
  data: T[];
  pagination: {
    page: number;
    pageSize: number;
    totalPages: number;
    totalItems: number;
  };
};

/**
 * API error response
 */
export type ApiError = {
  message: string;
  code?: string;
  field?: string;
  details?: Record<string, unknown>;
};
