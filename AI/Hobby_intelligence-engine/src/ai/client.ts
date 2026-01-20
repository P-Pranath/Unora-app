/**
 * AI Client Factory
 * 
 * Single entry point for creating AI clients.
 * Supports Hugging Face (default) and OpenAI providers.
 * 
 * Usage:
 *   const ai = createAIClient('huggingface');
 *   const response = await ai.complete({ ... });
 */

import { AIClient } from './types.js';
import { createHuggingFaceClient } from './providers/huggingface.js';

export type AIProvider = 'huggingface' | 'openai';

/**
 * Creates an AI client for the specified provider
 * 
 * @param provider - The AI provider to use (default: 'huggingface')
 * @returns An AIClient instance
 * 
 * To add a new provider:
 * 1. Create providers/newprovider.ts implementing AIClient
 * 2. Add case here to instantiate it
 */
export function createAIClient(provider: AIProvider = 'huggingface'): AIClient {
    switch (provider) {
        case 'huggingface':
            return createHuggingFaceClient();
        case 'openai':
            // Dynamic import to avoid loading OpenAI SDK when not needed
            // eslint-disable-next-line @typescript-eslint/no-require-imports
            const { createOpenAIClient } = require('./providers/openai.js');
            return createOpenAIClient();
        default:
            throw new Error(`Unknown AI provider: ${provider}`);
    }
}

// Re-export types for convenience
export { AIClient, AIRequest, AIResponse, AIError } from './types.js';

