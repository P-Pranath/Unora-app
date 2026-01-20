/**
 * Logger Utility
 * 
 * Structured logging for debugging and monitoring.
 * All logs include timestamps, levels, and context.
 */

export enum LogLevel {
    DEBUG = 'DEBUG',
    INFO = 'INFO',
    WARN = 'WARN',
    ERROR = 'ERROR'
}

interface LogContext {
    [key: string]: string | number | boolean | undefined;
}

interface LogEntry {
    timestamp: string;
    level: LogLevel;
    message: string;
    context?: LogContext;
    error?: {
        name: string;
        message: string;
        stack?: string;
    };
}

/**
 * Format log entry for console output
 */
function formatLog(entry: LogEntry): string {
    const levelColors: Record<LogLevel, string> = {
        [LogLevel.DEBUG]: '\x1b[90m', // Gray
        [LogLevel.INFO]: '\x1b[36m',  // Cyan
        [LogLevel.WARN]: '\x1b[33m',  // Yellow
        [LogLevel.ERROR]: '\x1b[31m'  // Red
    };
    const reset = '\x1b[0m';
    const color = levelColors[entry.level];

    let output = `${color}[${entry.timestamp}] [${entry.level}]${reset} ${entry.message}`;

    if (entry.context && Object.keys(entry.context).length > 0) {
        output += ` ${JSON.stringify(entry.context)}`;
    }

    if (entry.error) {
        output += `\n  ${color}Error: ${entry.error.name}: ${entry.error.message}${reset}`;
        if (entry.error.stack && process.env.NODE_ENV !== 'production') {
            output += `\n  ${entry.error.stack.split('\n').slice(1, 4).join('\n  ')}`;
        }
    }

    return output;
}

/**
 * Create a log entry and output to console
 */
function log(level: LogLevel, message: string, context?: LogContext, error?: Error): void {
    const entry: LogEntry = {
        timestamp: new Date().toISOString(),
        level,
        message,
        context
    };

    if (error) {
        entry.error = {
            name: error.name,
            message: error.message,
            stack: error.stack
        };
    }

    const formatted = formatLog(entry);

    switch (level) {
        case LogLevel.ERROR:
            console.error(formatted);
            break;
        case LogLevel.WARN:
            console.warn(formatted);
            break;
        default:
            console.log(formatted);
    }
}

/**
 * Logger functions
 */
export const logger = {
    debug: (message: string, context?: LogContext) => {
        if (process.env.NODE_ENV !== 'production') {
            log(LogLevel.DEBUG, message, context);
        }
    },

    info: (message: string, context?: LogContext) => {
        log(LogLevel.INFO, message, context);
    },

    warn: (message: string, context?: LogContext) => {
        log(LogLevel.WARN, message, context);
    },

    error: (message: string, error?: Error, context?: LogContext) => {
        log(LogLevel.ERROR, message, context, error);
    },

    // Specific error loggers for common scenarios

    uploadError: (error: Error, context?: { filename?: string; size?: number }) => {
        log(LogLevel.ERROR, 'UPLOAD_ERROR: File upload failed', {
            errorCode: 'UPLOAD_FAILED',
            ...context
        }, error);
    },

    cvValidationError: (error: Error, context?: { photoId?: string; check?: string }) => {
        log(LogLevel.ERROR, 'CV_VALIDATION_ERROR: Image analysis failed', {
            errorCode: 'CV_FAILED',
            ...context
        }, error);
    },

    aiValidationError: (error: Error, context?: { photoId?: string; client?: string }) => {
        log(LogLevel.ERROR, 'AI_VALIDATION_ERROR: AI analysis failed', {
            errorCode: 'AI_FAILED',
            ...context
        }, error);
    },

    storageError: (error: Error, context?: { operation?: string; path?: string }) => {
        log(LogLevel.ERROR, 'STORAGE_ERROR: Storage operation failed', {
            errorCode: 'STORAGE_FAILED',
            ...context
        }, error);
    },

    apiError: (error: Error, context?: { endpoint?: string; method?: string; statusCode?: number }) => {
        log(LogLevel.ERROR, 'API_ERROR: Request processing failed', {
            errorCode: 'API_FAILED',
            ...context
        }, error);
    },

    configError: (message: string, context?: { key?: string; expected?: string }) => {
        log(LogLevel.ERROR, `CONFIG_ERROR: ${message}`, {
            errorCode: 'CONFIG_INVALID',
            ...context
        });
    }
};

/**
 * Error codes reference for integration debugging
 */
export const ERROR_CODES = {
    // Upload errors
    UPLOAD_FAILED: 'File upload processing failed',
    NO_FILES: 'No files provided in request',
    INSUFFICIENT_PHOTOS: 'Less than 3 photos uploaded',
    TOO_MANY_FILES: 'More than 6 photos uploaded',
    FILE_TOO_LARGE: 'File exceeds 10MB limit',
    INVALID_FILE_TYPE: 'Unsupported file format',

    // CV Validation errors
    CV_FAILED: 'Computer vision analysis failed',
    CV_BLUR_CHECK_FAILED: 'Blur detection failed',
    CV_BRIGHTNESS_CHECK_FAILED: 'Brightness analysis failed',
    CV_RESOLUTION_CHECK_FAILED: 'Resolution check failed',

    // AI Validation errors
    AI_FAILED: 'AI validation failed',
    AI_CLIENT_NOT_READY: 'AI client not initialized',
    AI_RESPONSE_INVALID: 'AI response failed schema validation',
    AI_TIMEOUT: 'AI request timed out',

    // Storage errors
    STORAGE_FAILED: 'Storage operation failed',
    STORAGE_WRITE_FAILED: 'Failed to write file',
    STORAGE_READ_FAILED: 'Failed to read file',
    STORAGE_DELETE_FAILED: 'Failed to delete file',

    // Config errors
    CONFIG_INVALID: 'Configuration error',
    MISSING_API_KEY: 'OpenAI API key not configured',

    // General
    INTERNAL_ERROR: 'Internal server error',
    VALIDATION_ERROR: 'Request validation failed'
} as const;

export type ErrorCode = keyof typeof ERROR_CODES;
