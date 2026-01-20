import { Request, Response, NextFunction } from 'express';
import { ZodError } from 'zod';
import { MulterError } from 'multer';
import { FILE_CONSTRAINTS } from '../schemas/photo.schema';
import { logger } from '../utils/logger';

/**
 * Error response structure
 */
interface ErrorResponse {
    error: string;
    message: string;
    details?: string[];
}

/**
 * Custom application error class
 */
export class AppError extends Error {
    public statusCode: number;
    public code: string;
    public details?: string[];

    constructor(
        message: string,
        statusCode: number = 400,
        code: string = 'bad_request',
        details?: string[]
    ) {
        super(message);
        this.statusCode = statusCode;
        this.code = code;
        this.details = details;
        this.name = 'AppError';
    }
}

/**
 * Convert error to user-safe response
 * Ensures no sensitive information is leaked
 */
function toErrorResponse(error: Error): { status: number; body: ErrorResponse } {
    // Handle Multer errors
    if (error instanceof MulterError) {
        switch (error.code) {
            case 'LIMIT_FILE_SIZE':
                return {
                    status: 400,
                    body: {
                        error: 'file_too_large',
                        message: `File size exceeds the maximum limit of ${FILE_CONSTRAINTS.maxFileSize / (1024 * 1024)}MB`
                    }
                };
            case 'LIMIT_FILE_COUNT':
                return {
                    status: 400,
                    body: {
                        error: 'too_many_files',
                        message: `Maximum ${FILE_CONSTRAINTS.maxPhotos} photos allowed per upload`
                    }
                };
            case 'LIMIT_UNEXPECTED_FILE':
                return {
                    status: 400,
                    body: {
                        error: 'unexpected_field',
                        message: 'Please upload photos using the "photos" field'
                    }
                };
            default:
                return {
                    status: 400,
                    body: {
                        error: 'upload_error',
                        message: 'Failed to process file upload'
                    }
                };
        }
    }

    // Handle Zod validation errors
    if (error instanceof ZodError) {
        return {
            status: 400,
            body: {
                error: 'validation_error',
                message: 'Request validation failed',
                details: error.errors.map(e => `${e.path.join('.')}: ${e.message}`)
            }
        };
    }

    // Handle custom application errors
    if (error instanceof AppError) {
        return {
            status: error.statusCode,
            body: {
                error: error.code,
                message: error.message,
                details: error.details
            }
        };
    }

    // Handle file type errors from multer filter
    if (error.message.includes('Invalid file type')) {
        return {
            status: 400,
            body: {
                error: 'invalid_file_type',
                message: error.message
            }
        };
    }

    // Default: log internal errors but return generic message
    logger.error('Unhandled internal error', error, { type: error.name });
    return {
        status: 500,
        body: {
            error: 'internal_error',
            message: 'An unexpected error occurred. Please try again.'
        }
    };
}

/**
 * Global error handling middleware
 */
export function errorHandler(
    error: Error,
    req: Request,
    res: Response,
    next: NextFunction
): void {
    const { status, body } = toErrorResponse(error);

    // Log error for debugging
    if (process.env.NODE_ENV !== 'test') {
        logger.apiError(error, {
            endpoint: req.path,
            method: req.method,
            statusCode: status
        });
    }

    res.status(status).json(body);
}

/**
 * Handle 404 Not Found
 */
export function notFoundHandler(
    req: Request,
    res: Response,
    next: NextFunction
): void {
    res.status(404).json({
        error: 'not_found',
        message: `Route ${req.method} ${req.path} not found`
    });
}
