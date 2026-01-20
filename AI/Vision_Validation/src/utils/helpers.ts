import { v4 as uuidv4 } from 'uuid';

/**
 * Generate a new UUID v4
 */
export function generateId(): string {
    return uuidv4();
}

/**
 * Generate ISO timestamp for current time
 */
export function getCurrentTimestamp(): string {
    return new Date().toISOString();
}

/**
 * Safely extract file extension from filename
 */
export function getFileExtension(filename: string): string {
    const lastDot = filename.lastIndexOf('.');
    return lastDot !== -1 ? filename.substring(lastDot).toLowerCase() : '';
}

/**
 * Generate a unique filename for storage
 */
export function generateStorageFilename(originalFilename: string): string {
    const ext = getFileExtension(originalFilename);
    return `${generateId()}${ext}`;
}

/**
 * Normalize MIME type for HEIC/HEIF consistency
 */
export function normalizeMimetype(mimetype: string): string {
    if (mimetype === 'image/heif') {
        return 'image/heic';
    }
    return mimetype;
}

/**
 * Calculate base64 size in bytes
 */
export function calculateBase64Size(base64: string): number {
    // Remove data URL prefix if present
    const base64Data = base64.replace(/^data:image\/\w+;base64,/, '');
    return Math.ceil((base64Data.length * 3) / 4);
}

/**
 * Convert buffer to base64 data URL
 */
export function bufferToBase64DataUrl(buffer: Buffer, mimetype: string): string {
    const base64 = buffer.toString('base64');
    return `data:${mimetype};base64,${base64}`;
}

/**
 * Sleep for specified milliseconds (useful for testing/rate limiting)
 */
export function sleep(ms: number): Promise<void> {
    return new Promise(resolve => setTimeout(resolve, ms));
}

/**
 * Clamp a number between min and max values
 */
export function clamp(value: number, min: number, max: number): number {
    return Math.max(min, Math.min(max, value));
}

/**
 * Safe JSON parse with fallback
 */
export function safeJsonParse<T>(json: string, fallback: T): T {
    try {
        return JSON.parse(json);
    } catch {
        return fallback;
    }
}

/**
 * Extract JSON from a string that may contain markdown code blocks
 */
export function extractJsonFromString(text: string): string {
    // Try to extract JSON from markdown code block
    const jsonBlockMatch = text.match(/```(?:json)?\s*([\s\S]*?)```/);
    if (jsonBlockMatch) {
        return jsonBlockMatch[1].trim();
    }

    // Try to find raw JSON object
    const jsonMatch = text.match(/\{[\s\S]*\}/);
    if (jsonMatch) {
        return jsonMatch[0];
    }

    return text;
}
