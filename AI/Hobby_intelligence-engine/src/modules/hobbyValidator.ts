/**
 * Hobby Validator Module
 * 
 * Validates whether a user-provided hobby is legitimate and safe.
 * Uses AI classification (not generation) to make the determination.
 * 
 * Legitimacy criteria:
 * - Must be skill-based, practice-based, routine-based, or interest-based
 * - Must be non-sexual, non-illegal, non-harmful
 * - Must not be surveillance-based
 * - Must not be an identity, belief, or personality trait
 * 
 * Returns structured JSON with legit boolean, reason, and risk level.
 */

import { AIClient } from '../ai/client.js';
import { HobbyValidatorInput, HobbyValidatorOutput, hobbyValidatorOutputSchema } from '../types/schemas.js';

const SYSTEM_PROMPT = `You are a hobby classification system. Your ONLY task is to determine if a user-provided hobby is legitimate and safe.

A hobby is LEGITIMATE if it meets ALL of these criteria:
1. It is skill-based, practice-based, routine-based, or interest-based
2. It involves a learnable activity that can be practiced over time
3. It is not sexual, illegal, harmful, or dangerous to others
4. It is not surveillance-based (stalking, spying, monitoring others)
5. It is not an identity, belief, personality trait, or political stance

CLASSIFICATION RULES:
- "gym", "weightlifting", "running" = LEGIT (physical activity)
- "painting", "drawing", "music" = LEGIT (creative activity)  
- "reading", "learning languages" = LEGIT (intellectual activity)
- "lockpicking" = LEGIT if framed as sport/hobby (locksport is legitimate)
- "hacking" = LEGIT only if clearly ethical/CTF/learning context
- "stalking", "following people" = NOT LEGIT (surveillance/harmful)
- "hacking people", "catfishing" = NOT LEGIT (harmful to others)
- "being kind", "being funny" = NOT LEGIT (personality trait, not hobby)
- "Christianity", "atheism" = NOT LEGIT (belief, not hobby)

RISK LEVELS:
- low: Clearly safe activity (painting, reading, cooking)
- medium: Could be misused but generally safe (lockpicking sport, knife collecting)
- high: Borderline or potentially harmful (rejected hobbies)

You must respond with ONLY a JSON object in this exact format:
{
  "legit": boolean,
  "reason": string or null (explanation if NOT legit, null if legit),
  "risk_level": "low" | "medium" | "high"
}`;

/**
 * Creates a HobbyValidator instance
 * 
 * @param aiClient - The AI client to use for validation
 * @returns A validate function
 */
export function createHobbyValidator(aiClient: AIClient) {
    /**
     * Validates a hobby for legitimacy and safety
     * 
     * @param input - The hobby to validate
     * @returns Validation result with legit status, reason, and risk level
     */
    async function validate(input: HobbyValidatorInput): Promise<HobbyValidatorOutput> {
        const userPrompt = `Classify this hobby: "${input.hobby}"`;

        const response = await aiClient.complete<HobbyValidatorOutput>({
            systemPrompt: SYSTEM_PROMPT,
            userPrompt,
            temperature: 0.1, // Very low for consistent classification
        });

        // Validate response against schema
        const parsed = hobbyValidatorOutputSchema.safeParse(response.data);

        if (!parsed.success) {
            // Fallback to safe rejection if AI response is malformed
            return {
                legit: false,
                reason: 'Unable to validate hobby',
                risk_level: 'high',
            };
        }

        return parsed.data;
    }

    return { validate };
}

export type HobbyValidator = ReturnType<typeof createHobbyValidator>;
