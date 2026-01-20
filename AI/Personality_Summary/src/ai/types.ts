// ============================================================================
// AI Provider Types & Interface Contract
// ============================================================================
// Both PuterProvider and OpenAIProvider must implement this interface.
// No provider-specific logic may leak outside of the provider implementations.
// ============================================================================

/**
 * Response from AI provider after generating JSON content
 */
export interface AIResponse {
    /** Raw text response from the LLM */
    rawText: string;
    /** Parsed JSON object */
    data: Record<string, unknown>;
    /** Provider that generated this response */
    provider: string;
}

/**
 * Configuration for provider initialization
 */
export interface ProviderConfig {
    /** Model name to use */
    model: string;
    /** Temperature for generation (0.0 - 1.0) */
    temperature?: number;
    /** Maximum tokens in response */
    maxTokens?: number;
}

/**
 * Interface contract that ALL AI providers must implement.
 * This ensures consistent behavior regardless of which provider is active.
 */
export interface AIProvider {
    /** Human-readable name of the provider */
    readonly name: string;
    
    /** Model being used */
    readonly model: string;

    /**
     * Generate a JSON response from the LLM.
     * 
     * @param systemPrompt - System-level instructions for the LLM
     * @param userPrompt - User's request/query
     * @returns Promise resolving to parsed JSON object
     * @throws Error if generation fails or response is not valid JSON
     */
    generate_json(systemPrompt: string, userPrompt: string): Promise<Record<string, unknown>>;

    /**
     * Generate raw text response from the LLM.
     * 
     * @param systemPrompt - System-level instructions for the LLM
     * @param userPrompt - User's request/query
     * @returns Promise resolving to text response
     */
    generate_text(systemPrompt: string, userPrompt: string): Promise<string>;
}

/**
 * Error thrown when AI generation fails
 */
export class AIProviderError extends Error {
    constructor(
        public readonly provider: string,
        public readonly originalError: Error | unknown,
        message: string
    ) {
        super(`[${provider}] ${message}`);
        this.name = 'AIProviderError';
    }
}
