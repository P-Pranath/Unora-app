/**
 * Hugging Face Provider Implementation
 * 
 * Implements the AIClient interface using Hugging Face's Inference API.
 * Uses the OpenAI-compatible chat completions endpoint for structured responses.
 * 
 * Design decisions:
 * - Uses serverless inference for cost efficiency
 * - Low temperature (0.3) for more deterministic outputs
 * - Automatic retry logic for transient failures
 */

import { HfInference } from '@huggingface/inference';
import { AIClient, AIRequest, AIResponse, AIError } from '../types.js';
import { config } from '../../config/env.js';

const MAX_RETRIES = 3;
const RETRY_DELAY_MS = 1000;

/**
 * Creates a Hugging Face-backed AI client
 */
export function createHuggingFaceClient(): AIClient {
    const client = new HfInference(config.HUGGINGFACE_API_KEY);

    return {
        async complete<T = unknown>(request: AIRequest): Promise<AIResponse<T>> {
            const { systemPrompt, userPrompt, temperature, maxTokens } = request;

            let lastError: unknown;

            for (let attempt = 1; attempt <= MAX_RETRIES; attempt++) {
                try {
                    const response = await client.chatCompletion({
                        model: config.AI_MODEL,
                        messages: [
                            { role: 'system', content: systemPrompt },
                            { role: 'user', content: userPrompt },
                        ],
                        temperature: temperature ?? config.AI_TEMPERATURE,
                        max_tokens: maxTokens ?? config.AI_MAX_TOKENS,
                    });

                    const content = response.choices[0]?.message?.content;

                    if (!content) {
                        throw new AIError(
                            'Empty response from Hugging Face',
                            'EMPTY_RESPONSE',
                            'huggingface'
                        );
                    }

                    // Parse JSON response - extract JSON from potential markdown code blocks
                    let data: T;
                    try {
                        // Try to extract JSON from markdown code block if present
                        const jsonMatch = content.match(/```(?:json)?\s*([\s\S]*?)\s*```/);
                        const jsonString = jsonMatch ? jsonMatch[1] : content;
                        data = JSON.parse(jsonString.trim()) as T;
                    } catch (parseError) {
                        throw new AIError(
                            'Failed to parse JSON response from Hugging Face',
                            'JSON_PARSE_ERROR',
                            'huggingface',
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

                    // Don't retry on AIError (our custom errors)
                    if (error instanceof AIError) {
                        throw error;
                    }

                    // Retry on other errors (network, server issues)
                    if (attempt < MAX_RETRIES) {
                        await sleep(RETRY_DELAY_MS * attempt);
                        continue;
                    }
                }
            }

            // All retries exhausted
            console.error('Hugging Face API Error Details:', lastError);
            throw new AIError(
                `Hugging Face request failed after ${MAX_RETRIES} attempts`,
                'MAX_RETRIES_EXCEEDED',
                'huggingface',
                lastError
            );
        },
    };
}

function sleep(ms: number): Promise<void> {
    return new Promise((resolve) => setTimeout(resolve, ms));
}
