import OpenAI from 'openai';
import fs from 'fs';
import path from 'path';
import { IAIClient, FALLBACK_AI_RESULT } from './ai.client.interface';
import { AIValidationResult, AIValidationResultSchema } from '../schemas/validation.schema';
import { config } from '../config/env';
import { extractJsonFromString, safeJsonParse } from '../utils/helpers';

/**
 * OpenAI Vision Client
 * 
 * Uses GPT-4 Vision (gpt-4o) to validate photos for:
 * - Face presence and visibility
 * - Photo type (real vs screenshot/meme/AI-generated)
 * - Image quality assessment
 * 
 * IMPORTANT: Results are ephemeral and must NOT be stored permanently.
 */
export class OpenAIClient implements IAIClient {
    private client: OpenAI;
    private systemPrompt: string;

    constructor() {
        if (!config.openaiApiKey) {
            throw new Error('OPENAI_API_KEY is required for OpenAI client');
        }

        this.client = new OpenAI({
            apiKey: config.openaiApiKey
        });

        // Load system prompt
        this.systemPrompt = this.loadSystemPrompt();
    }

    /**
     * Load system prompt from file
     */
    private loadSystemPrompt(): string {
        const promptPath = path.resolve(__dirname, '../../prompts/validation_prompt.md');

        try {
            if (fs.existsSync(promptPath)) {
                return fs.readFileSync(promptPath, 'utf-8');
            }
        } catch (error) {
            console.warn('Could not load prompt file, using embedded prompt');
        }

        // Fallback embedded prompt
        return `You are an AI vision validation service.
Your sole responsibility is to evaluate whether an uploaded photo is usable for onboarding verification.

You are NOT performing:
- Identity verification
- Face recognition  
- Attractiveness scoring
- Age, gender, or emotion inference

You are STRICTLY checking photo usability.

OUTPUT JSON ONLY:
{
  "usable": true | false,
  "issues": [string],
  "confidence": "high" | "medium" | "low"
}

EVALUATION CRITERIA:
1. FACE PRESENCE - At least one real human face visible
2. FACE VISIBILITY - Face not fully hidden by mask/helmet/object
3. IMAGE QUALITY - Not blurry, too dark, too bright, or pixelated
4. PHOTO TYPE - Real photo, not screenshot/meme/drawing/AI-generated

ISSUE TAGS (use only these):
- "no_face"
- "blurred"
- "too_dark"
- "too_bright"
- "obstructed_face"
- "non_photographic"
- "low_resolution"

CONFIDENCE:
- "high": clear decision
- "medium": borderline but acceptable
- "low": ambiguous, lean toward reject

Be conservative: when unsure, mark unusable.
Return ONLY valid JSON.`;
    }

    /**
     * Check if client is properly configured
     */
    isReady(): boolean {
        return !!config.openaiApiKey;
    }

    /**
     * Get client name
     */
    getName(): string {
        return 'OpenAI GPT-4 Vision';
    }

    /**
     * Validate a photo using GPT-4 Vision
     */
    async validatePhoto(imageBase64: string): Promise<AIValidationResult> {
        try {
            // Ensure proper data URL format
            let imageUrl = imageBase64;
            if (!imageBase64.startsWith('data:')) {
                imageUrl = `data:image/jpeg;base64,${imageBase64}`;
            }

            const response = await this.client.chat.completions.create({
                model: 'gpt-4o',
                messages: [
                    {
                        role: 'system',
                        content: this.systemPrompt
                    },
                    {
                        role: 'user',
                        content: [
                            {
                                type: 'image_url',
                                image_url: {
                                    url: imageUrl,
                                    detail: 'low' // Use low detail for faster processing
                                }
                            },
                            {
                                type: 'text',
                                text: 'Analyze this photo and return the validation result as JSON.'
                            }
                        ]
                    }
                ],
                max_tokens: 200,
                temperature: 0.1 // Low temperature for consistent results
            });

            const content = response.choices[0]?.message?.content;

            if (!content) {
                console.error('OpenAI returned empty response');
                return FALLBACK_AI_RESULT;
            }

            // Extract and parse JSON from response
            const jsonString = extractJsonFromString(content);
            const parsed = safeJsonParse(jsonString, null);

            if (!parsed) {
                console.error('Failed to parse OpenAI response:', content);
                return FALLBACK_AI_RESULT;
            }

            // Validate against schema
            const validated = AIValidationResultSchema.safeParse(parsed);

            if (!validated.success) {
                console.error('OpenAI response failed schema validation:', validated.error);
                return FALLBACK_AI_RESULT;
            }

            return validated.data;

        } catch (error) {
            console.error('OpenAI validation error:', error);
            return FALLBACK_AI_RESULT;
        }
    }
}

/**
 * Create OpenAI client instance
 */
export function createOpenAIClient(): IAIClient {
    return new OpenAIClient();
}
