import { AIValidationResult } from '../schemas/validation.schema';

/**
 * AI Client Interface
 * 
 * Defines the contract for AI vision validation clients.
 * Implementations can be swapped between OpenAI and Mock for testing.
 * 
 * IMPORTANT: AI validation results are EPHEMERAL and must NOT be stored.
 * Only the final usable/has_face flags should be persisted.
 */
export interface IAIClient {
    /**
     * Validate a photo for face presence and usability
     * 
     * @param imageBase64 - Base64 encoded image data (with or without data URL prefix)
     * @returns Promise resolving to validation result
     * 
     * Response must contain ONLY:
     * - usable: boolean
     * - issues: string[] (from allowed tags)
     * - confidence: 'high' | 'medium' | 'low'
     */
    validatePhoto(imageBase64: string): Promise<AIValidationResult>;

    /**
     * Optional: Check if the client is ready/configured
     */
    isReady(): boolean;

    /**
     * Get client name/type for logging
     */
    getName(): string;
}

/**
 * Default fallback result when AI validation fails
 * Conservative: marks as unusable with low confidence
 */
export const FALLBACK_AI_RESULT: AIValidationResult = {
    usable: false,
    issues: [],
    confidence: 'low'
};

/**
 * Factory type for creating AI clients
 */
export type AIClientFactory = () => IAIClient;
