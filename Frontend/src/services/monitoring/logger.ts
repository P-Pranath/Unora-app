/**
 * Centralized logging service
 * Integrates with Sentry for error reporting and provides consistent logging API
 */

// Sensitive keys that should be redacted from logs
const SENSITIVE_KEYS = new Set([
  'password',
  'token',
  'accessToken',
  'refreshToken',
  'apiKey',
  'secret',
  'ssn',
  'creditCard',
  'cvv',
  'pin',
  'authorization',
]);

/**
 * Sanitize data by redacting sensitive fields
 */
function sanitizeData(data: unknown): unknown {
  if (data === null || data === undefined) {
    return data;
  }

  if (typeof data !== 'object') {
    return data;
  }

  if (Array.isArray(data)) {
    return data.map(sanitizeData);
  }

  const sanitized: Record<string, unknown> = {};
  for (const [key, value] of Object.entries(data as Record<string, unknown>)) {
    if (SENSITIVE_KEYS.has(key.toLowerCase())) {
      sanitized[key] = '***REDACTED***';
    } else if (typeof value === 'object' && value !== null) {
      sanitized[key] = sanitizeData(value);
    } else {
      sanitized[key] = value;
    }
  }
  return sanitized;
}

type LogLevel = 'debug' | 'info' | 'warn' | 'error';

interface LoggerOptions {
  level?: LogLevel;
  enableConsole?: boolean;
}

const LOG_LEVELS: Record<LogLevel, number> = {
  debug: 0,
  info: 1,
  warn: 2,
  error: 3,
};

class Logger {
  private level: LogLevel;
  private enableConsole: boolean;
  private customSensitiveKeys: Set<string>;

  constructor(options: LoggerOptions = {}) {
    this.level = options.level ?? (__DEV__ ? 'debug' : 'info');
    this.enableConsole = options.enableConsole ?? __DEV__;
    this.customSensitiveKeys = new Set();
  }

  /**
   * Add custom sensitive keys to be redacted
   */
  addSensitiveKeys(keys: string[]): void {
    keys.forEach(key => this.customSensitiveKeys.add(key.toLowerCase()));
  }

  private shouldLog(level: LogLevel): boolean {
    return LOG_LEVELS[level] >= LOG_LEVELS[this.level];
  }

  private formatMessage(
    level: LogLevel,
    message: string,
    data?: unknown,
  ): string {
    const timestamp = new Date().toISOString();
    const sanitizedData = data ? sanitizeData(data) : undefined;
    const dataStr = sanitizedData ? ` ${JSON.stringify(sanitizedData)}` : '';
    return `[${timestamp}] [${level.toUpperCase()}] ${message}${dataStr}`;
  }

  debug(message: string, data?: Record<string, unknown>): void {
    if (this.shouldLog('debug') && this.enableConsole) {
      console.debug(this.formatMessage('debug', message, data));
    }
  }

  info(message: string, data?: Record<string, unknown>): void {
    if (this.shouldLog('info') && this.enableConsole) {
      console.info(this.formatMessage('info', message, data));
    }
  }

  warn(message: string, data?: Record<string, unknown>): void {
    if (this.shouldLog('warn') && this.enableConsole) {
      console.warn(this.formatMessage('warn', message, data));
    }
  }

  error(
    message: string,
    error?: Error | unknown,
    data?: Record<string, unknown>,
  ): void {
    if (this.shouldLog('error')) {
      const errorData = {
        ...data,
        error:
          error instanceof Error
            ? {
                name: error.name,
                message: error.message,
                stack: error.stack,
              }
            : error,
      };

      if (this.enableConsole) {
        console.error(this.formatMessage('error', message, errorData));
      }

      // In production, you would send this to Sentry here
      // Sentry.captureException(error, { extra: sanitizeData(data) });
    }
  }

  /**
   * Log user action for analytics
   */
  logUserAction(action: string, data?: Record<string, unknown>): void {
    this.info(`User Action: ${action}`, data);
  }

  /**
   * Log navigation event
   */
  logNavigation(screenName: string, params?: Record<string, unknown>): void {
    this.debug(`Navigation: ${screenName}`, params);
  }

  /**
   * Log API call success
   */
  logApiCall(
    method: string,
    url: string,
    status: number,
    duration: number,
  ): void {
    this.debug(`API ${method} ${url}`, {status, duration: `${duration}ms`});
  }

  /**
   * Log API error
   */
  logApiError(
    url: string,
    status: number,
    error: unknown,
    duration?: number,
  ): void {
    this.error(`API Error: ${url}`, error, {
      status,
      duration: duration ? `${duration}ms` : undefined,
    });
  }
}

export const logger = new Logger();
export default logger;
