// ============================================================================
// AI Client - Unified LLM Provider Interface
// ============================================================================
// Central abstraction for all LLM calls. All business logic should use this
// client, never raw providers directly.
// ============================================================================

import { AIProvider } from './types';
import { OpenAIProvider } from './providers/openai_provider';

// ═══════════════════════════════════════════════════════════════════════════
// CONFIGURATION
// ═══════════════════════════════════════════════════════════════════════════
// Ensure OPENAI_API_KEY is set in your .env file
// ═══════════════════════════════════════════════════════════════════════════

const provider: AIProvider = new OpenAIProvider();

// ═══════════════════════════════════════════════════════════════════════════

// Log active provider
console.log('━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━');
console.log('🤖 AI CLIENT CONFIGURATION');
console.log('━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━');
console.log(`   Active Provider: ${provider.name}`);
console.log(`   Model: ${provider.model}`);
console.log('━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n');

/**
 * AI Client - Unified interface for LLM operations
 * 
 * Usage:
 *   import { AIClient } from '../ai/ai_client';
 *   const result = await AIClient.generateJSON(systemPrompt, userPrompt);
 */
export const AIClient = {
    /**
     * Get the active provider name
     */
    get providerName(): string {
        return provider.name;
    },

    /**
     * Get the active model name
     */
    get modelName(): string {
        return provider.model;
    },

    /**
     * Generate JSON response from the LLM provider.
     * 
     * @param systemPrompt - System-level instructions for the LLM
     * @param userPrompt - User's request/query
     * @returns Promise resolving to parsed JSON object
     */
    async generateJSON(systemPrompt: string, userPrompt: string): Promise<Record<string, unknown>> {
        console.log(`   [AIClient] Calling ${provider.name} for JSON generation...`);
        return provider.generate_json(systemPrompt, userPrompt);
    },

    /**
     * Generate raw text response from the LLM provider.
     * 
     * @param systemPrompt - System-level instructions for the LLM
     * @param userPrompt - User's request/query
     * @returns Promise resolving to text response
     */
    async generateText(systemPrompt: string, userPrompt: string): Promise<string> {
        console.log(`   [AIClient] Calling ${provider.name} for text generation...`);
        return provider.generate_text(systemPrompt, userPrompt);
    }
};

// Also export the provider type for advanced usage
export type { AIProvider } from './types';
