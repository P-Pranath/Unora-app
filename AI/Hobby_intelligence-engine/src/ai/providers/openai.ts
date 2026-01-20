/**
 * OpenAI Provider Implementation
 * 
 * Implements the AIClient interface using OpenAI's API.
 * Forces JSON output mode for structured, deterministic responses.
 * 
 * Design decisions:
 * - Uses gpt-4o-mini for cost efficiency while maintaining quality
 * - Low temperature (0.3) for more deterministic outputs
 * - JSON mode enforced via response_format
 * - Automatic retry logic for transient failures
 */

import OpenAI from 'openai';
import { AIClient, AIRequest, AIResponse, AIError } from '../types.js';
import { config } from '../../config/env.js';

const MAX_RETRIES = 3;
const RETRY_DELAY_MS = 1000;

/**
 * Creates an OpenAI-backed AI client
 */
export function createOpenAIClient(): AIClient {
    const client = new OpenAI({
        apiKey: config.OPENAI_API_KEY,
    });

    return {
        async complete<T = unknown>(request: AIRequest): Promise<AIResponse<T>> {
            const { systemPrompt, userPrompt, temperature, maxTokens } = request;

            let lastError: unknown;

            for (let attempt = 1; attempt <= MAX_RETRIES; attempt++) {
                try {
                    const response = await client.chat.completions.create({
                        model: config.AI_MODEL,
                        messages: [
                            { role: 'system', content: systemPrompt },
                            { role: 'user', content: userPrompt },
                        ],
                        temperature: temperature ?? config.AI_TEMPERATURE,
                        max_tokens: maxTokens ?? config.AI_MAX_TOKENS,
                        // Force JSON output for structured responses
                        response_format: { type: 'json_object' },
                    });

                    const content = response.choices[0]?.message?.content;

                    if (!content) {
                        throw new AIError(
                            'Empty response from OpenAI',
                            'EMPTY_RESPONSE',
                            'openai'
                        );
                    }

                    // Parse JSON response
                    let data: T;
                    try {
                        data = JSON.parse(content) as T;
                    } catch (parseError) {
                        throw new AIError(
                            'Failed to parse JSON response from OpenAI',
                            'JSON_PARSE_ERROR',
                            'openai',
                            parseError
                        );
                    }

                    return {
                        data,
                        usage: {
                            promptTokens: response.usage?.prompt_tokens ?? 0,
                            completionTokens: response.usage?.completion_tokens ?? 0,
                            totalTokens: response.usage?.total_tokens ?? 0,
                        },
                    };
                } catch (error) {
                    lastError = error;

                    // Don't retry on client errors (4xx)
                    if (error instanceof OpenAI.APIError && error.status && error.status < 500) {
                        throw new AIError(
                            `OpenAI API error: ${error.message}`,
                            error.code ?? 'API_ERROR',
                            'openai',
                            error
                        );
                    }

                    // Retry on server errors (5xx) and network issues
                    if (attempt < MAX_RETRIES) {
                        await sleep(RETRY_DELAY_MS * attempt);
                        continue;
                    }
                }
            }

            // All retries exhausted
            throw new AIError(
                `OpenAI request failed after ${MAX_RETRIES} attempts`,
                'MAX_RETRIES_EXCEEDED',
                'openai',
                lastError
            );
        },
    };
}

function sleep(ms: number): Promise<void> {
    return new Promise((resolve) => setTimeout(resolve, ms));
}
