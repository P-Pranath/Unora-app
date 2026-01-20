# Unora API Error Contract

## Standard Error Response Format

All API endpoints return errors in this consistent format:

```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE",
    "message": "Human-readable message",
    "details": { }  // Optional additional context
  }
}
```

## Error Codes

### Client Errors (4xx)

| Code | HTTP Status | Description |
|------|-------------|-------------|
| `VALIDATION_ERROR` | 400 | Request validation failed |
| `INVALID_REQUEST` | 400 | Malformed request body |
| `UNAUTHORIZED` | 401 | Missing or invalid authentication |
| `FORBIDDEN` | 403 | Access denied (valid auth but no permission) |
| `NOT_FOUND` | 404 | Resource not found |
| `CONFLICT` | 409 | Resource conflict (e.g., already exists) |
| `RATE_LIMITED` | 429 | Too many requests / at capacity |

### Server Errors (5xx)

| Code | HTTP Status | Description |
|------|-------------|-------------|
| `INTERNAL_ERROR` | 500 | Unexpected server error |
| `DATABASE_ERROR` | 500 | Database operation failed |
| `EXTERNAL_SERVICE_ERROR` | 502 | Third-party service failed |
| `SERVICE_UNAVAILABLE` | 503 | Service temporarily unavailable |

## TypeScript Types

```typescript
interface APIErrorResponse {
  success: false;
  error: {
    code: ErrorCode;
    message: string;
    details?: Record<string, unknown>;
  };
}

type ErrorCode =
  // Client errors
  | 'VALIDATION_ERROR'
  | 'INVALID_REQUEST'
  | 'UNAUTHORIZED'
  | 'FORBIDDEN'
  | 'NOT_FOUND'
  | 'CONFLICT'
  | 'RATE_LIMITED'
  // Server errors
  | 'INTERNAL_ERROR'
  | 'DATABASE_ERROR'
  | 'EXTERNAL_SERVICE_ERROR'
  | 'SERVICE_UNAVAILABLE';
```

## FE Error Handling Example

```typescript
async function apiCall<T>(url: string, options?: RequestInit): Promise<T> {
  const response = await fetch(url, options);
  const data = await response.json();
  
  if (!data.success) {
    const { code, message } = data.error;
    
    switch (code) {
      case 'UNAUTHORIZED':
        // Redirect to login
        break;
      case 'RATE_LIMITED':
        // Show "slow down" message
        break;
      case 'NOT_FOUND':
        // Show 404 UI
        break;
      default:
        // Show generic error toast
    }
    
    throw new APIError(code, message);
  }
  
  return data.data;
}
```

## Success Response Format

```json
{
  "success": true,
  "data": { ... }
}
```
