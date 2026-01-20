import fs from 'fs';
import path from 'path';
import { IAIClient, FALLBACK_AI_RESULT } from '../clients/ai.client.interface';
import { createOpenAIClient } from '../clients/openai.client';
import { createMockAIClient, MockAIClient } from '../clients/mock.client';
import { AIValidationResult } from '../schemas/validation.schema';
import { config } from '../config/env';
import { bufferToBase64DataUrl } from '../utils/helpers';
import { logger } from '../utils/logger';

/**
 * AI Validation Service
 * 
 * Orchestrates AI-powered photo validation.
 * Handles client selection, image preparation, and result processing.
 * 
 * CRITICAL: AI analysis results are EPHEMERAL and must NOT be stored.
 * Only the final usable/has_face flags should be persisted.
 */

// Singleton AI client instance
let aiClient: IAIClient | null = null;

/**
 * Get or create AI client based on configuration
 */
export function getAIClient(): IAIClient {
    if (aiClient) {
        return aiClient;
    }

    if (config.useMockAI) {
        logger.info('Initializing Mock AI Client', { mode: 'mock' });
        aiClient = createMockAIClient();
    } else if (config.openaiApiKey) {
        logger.info('Initializing OpenAI Client', { mode: 'openai' });
        aiClient = createOpenAIClient();
    } else {
        logger.warn('No AI configuration found, defaulting to Mock Client', { reason: 'missing_api_key' });
        aiClient = createMockAIClient();
    }

    return aiClient;
}

/**
 * Reset AI client (useful for testing)
 */
export function resetAIClient(): void {
    aiClient = null;
}

/**
 * Read image file and convert to base64
 */
async function imageToBase64(imagePath: string): Promise<string> {
    const buffer = await fs.promises.readFile(imagePath);

    // Determine MIME type from extension
    const ext = path.extname(imagePath).toLowerCase();
    let mimetype = 'image/jpeg';

    switch (ext) {
        case '.png':
            mimetype = 'image/png';
            break;
        case '.heic':
        case '.heif':
            mimetype = 'image/heic';
            break;
        case '.webp':
            mimetype = 'image/webp';
            break;
        default:
            mimetype = 'image/jpeg';
    }

    return bufferToBase64DataUrl(buffer, mimetype);
}

/**
 * Validate a photo using AI vision
 * 
 * @param imagePath - Path to the image file
 * @param originalFilename - Original filename (for mock client hints)
 * @returns AI validation result
 * 
 * IMPORTANT: This result is ephemeral and should NOT be stored.
 */
export async function validateWithAI(
    imagePath: string,
    originalFilename?: string
): Promise<AIValidationResult> {
    const client = getAIClient();

    if (!client.isReady()) {
        logger.warn('AI client not ready, returning fallback result', { client: client.getName() });
        return FALLBACK_AI_RESULT;
    }

    try {
        // Set filename hint for mock client
        if (originalFilename && client instanceof MockAIClient) {
            (client as MockAIClient).setFilenameHint(originalFilename);
        }

        // Convert image to base64
        const imageBase64 = await imageToBase64(imagePath);

        // Validate with AI
        const result = await client.validatePhoto(imageBase64);

        logger.info('AI validation completed', {
            client: client.getName(),
            usable: result.usable,
            issues: result.issues.join(','),
            confidence: result.confidence
        });

        return result;

    } catch (error) {
        logger.aiValidationError(error as Error, { client: client.getName() });
        return FALLBACK_AI_RESULT;
    }
}

/**
 * Determine if a photo has a face based on AI result
 */
export function hasFaceFromAIResult(result: AIValidationResult): boolean {
    // If no_face issue is present, there's no face
    if (result.issues.includes('no_face')) {
        return false;
    }

    // If usable and no no_face issue, assume face is present
    // Conservative: if not usable for other reasons, still might have face
    return result.usable || !result.issues.includes('no_face');
}

/**
 * Get AI client info for debugging
 */
export function getAIClientInfo(): { name: string; ready: boolean } {
    const client = getAIClient();
    return {
        name: client.getName(),
        ready: client.isReady()
    };
}
