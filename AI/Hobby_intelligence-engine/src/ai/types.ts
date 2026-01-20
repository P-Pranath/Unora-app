/**
 * AI Client Type Definitions
 * 
 * Provider-agnostic interfaces for AI operations.
 * All AI usage in the codebase goes through these types.
 * Designed for easy provider swapping (OpenAI → local LLM, etc.)
 */

/**
 * Request structure for AI completions
 * Uses structured input to ensure deterministic, JSON-only responses
 */
export interface AIRequest {
    /** System prompt defining the AI's role and constraints */
    systemPrompt: string;

    /** User input/query to process */
    userPrompt: string;

    /** Optional: Override default temperature (0-1, lower = more deterministic) */
    temperature?: number;

    /** Optional: Override default max tokens */
    maxTokens?: number;
}

/**
 * Response structure from AI completions
 * Always returns parsed JSON - never raw text
 */
export interface AIResponse<T = unknown> {
    /** Parsed JSON response from the AI */
    data: T;

    /** Token usage for monitoring/billing */
    usage: {
        promptTokens: number;
        completionTokens: number;
        totalTokens: number;
    };
}

/**
 * AI Client Interface
 * 
 * Single method interface for maximum flexibility.
 * All AI operations go through complete() with appropriate prompts.
 */
export interface AIClient {
    /**
     * Send a completion request and receive structured JSON response
     * @param request - The AI request with prompts and optional overrides
     * @returns Parsed JSON response with usage stats
     * @throws AIError on failure
     */
    complete<T = unknown>(request: AIRequest): Promise<AIResponse<T>>;
}

/**
 * Custom error type for AI-related failures
 */
export class AIError extends Error {
    constructor(
        message: string,
        public readonly code: string,
        public readonly provider: string,
        public readonly cause?: unknown
    ) {
        super(message);
        this.name = 'AIError';
    }
}
