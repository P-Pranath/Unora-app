// ============================================================================
// Gemini Provider - Uses Google's Native Generative AI API
// ============================================================================

import { AIProvider, AIProviderError, ProviderConfig } from '../types';

/**
 * Default configuration
 */
const DEFAULT_CONFIG: ProviderConfig = {
    model: process.env.AI_MODEL || 'gemini-2.5-pro',
    temperature: 0.7,
    maxTokens: 200
};

/**
 * Sleep helper for retry logic
 */
function sleep(ms: number): Promise<void> {
    return new Promise(resolve => setTimeout(resolve, ms));
}

/**
 * Gemini LLM provider using native Google API
 */
export class OpenAIProvider implements AIProvider {
    public readonly name = 'Gemini';
    public readonly model: string;

    private readonly apiKey: string;
    private readonly config: ProviderConfig;
    private readonly maxRetries = 3;

    constructor(config?: Partial<ProviderConfig>) {
        this.config = { ...DEFAULT_CONFIG, ...config };
        this.model = this.config.model;
        this.apiKey = process.env.GEMINI_API_KEY || process.env.OPENAI_API_KEY || '';

        if (!this.apiKey) {
            console.warn('⚠️  No API key found. Set GEMINI_API_KEY in .env');
        }

        console.log(`   [${this.name}] Provider initialized`);
        console.log(`   [${this.name}] Model: ${this.model}`);
        console.log(`   [${this.name}] API Key: ${this.apiKey ? '✓ Set' : '✗ Missing'}`);
    }

    /**
     * Call Gemini API directly
     */
    private async callGemini(prompt: string): Promise<string> {
        const url = `https://generativelanguage.googleapis.com/v1beta/models/${this.model}:generateContent?key=${this.apiKey}`;

        const response = await fetch(url, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                contents: [{ parts: [{ text: prompt }] }],
                generationConfig: {
                    temperature: this.config.temperature,
                    maxOutputTokens: this.config.maxTokens
                }
            })
        });

        if (!response.ok) {
            const error = await response.json().catch(() => ({ error: { message: 'Unknown error' } }));
            throw new Error(`${response.status}: ${error.error?.message || 'API error'}`);
        }

        const data = await response.json();
        const text = data.candidates?.[0]?.content?.parts?.[0]?.text;

        if (!text) {
            throw new Error('Empty response from Gemini');
        }

        return text.trim();
    }

    /**
     * Generate raw text response with retry logic
     */
    async generate_text(systemPrompt: string, userPrompt: string): Promise<string> {
        const fullPrompt = `${systemPrompt}\n\n${userPrompt}`;
        let lastError: Error | null = null;

        for (let attempt = 1; attempt <= this.maxRetries; attempt++) {
            try {
                console.log(`   [${this.name}] Attempt ${attempt}/${this.maxRetries}`);
                const result = await this.callGemini(fullPrompt);
                console.log(`   [${this.name}] Success!`);
                return result;
            } catch (error: unknown) {
                lastError = error as Error;
                const message = (error as Error).message || '';

                console.log(`   [${this.name}] Error: ${message}`);

                // Rate limit - wait and retry
                if (message.includes('429') || message.includes('quota') || message.includes('rate')) {
                    const waitTime = attempt * 5000; // 5s, 10s, 15s
                    console.log(`   [${this.name}] Rate limited. Waiting ${waitTime / 1000}s...`);
                    await sleep(waitTime);
                    continue;
                }

                // Other errors - break
                break;
            }
        }

        throw new AIProviderError(this.name, lastError,
            `Failed after ${this.maxRetries} attempts: ${lastError?.message || 'Unknown error'}`
        );
    }

    /**
     * Generate JSON response
     */
    async generate_json(systemPrompt: string, userPrompt: string): Promise<Record<string, unknown>> {
        const jsonPrompt = `${systemPrompt}\n\nRespond with valid JSON only. No markdown.\n\n${userPrompt}`;

        const text = await this.generate_text(jsonPrompt, '');

        try {
            return JSON.parse(text) as Record<string, unknown>;
        } catch {
            const objectMatch = text.match(/\{[\s\S]*\}/);
            if (objectMatch) {
                return JSON.parse(objectMatch[0]) as Record<string, unknown>;
            }
            throw new AIProviderError(this.name, null, `Invalid JSON: ${text.substring(0, 100)}...`);
        }
    }
}
