/**
 * Hobby Echo Generator Module
 * 
 * Generates privacy-safe progress updates ("Hobby Echoes").
 * These are shareable summaries of check-in activity.
 * 
 * Rules (Non-negotiable):
 * - Third-person only ("Your partner...")
 * - No exaggeration
 * - No inferred details beyond what was selected
 * - No emotional or performance claims
 * - No identity, time, or location information
 * - Privacy-safe for sharing with connections
 */

import { AIClient } from '../ai/client.js';
import {
    HobbyEchoInput,
    HobbyEchoOutput,
    hobbyEchoOutputSchema
} from '../types/schemas.js';

const SYSTEM_PROMPT = `You are a privacy-safe progress update generator. Create a brief, neutral third-person statement about a hobby check-in.

STRICT RULES:
1. Use ONLY third-person ("Your partner...")
2. NO exaggeration (amazing, incredible, fantastic)
3. NO inferred details beyond what was provided
4. NO emotional claims (felt great, was happy, enjoyed)
5. NO performance assessment (did well, improved, mastered)
6. NO time references (today, yesterday, this morning)
7. NO location references  
8. NO identity claims
9. Keep it SHORT - one sentence only
10. State ONLY what was checked in, nothing more

GOOD EXAMPLES:
- Hobby: "painting", Option: "Landscapes" → "Your partner checked in with painting and focused on landscapes."
- Hobby: "gym", Option: "Strength" → "Your partner checked in with their gym routine and worked on strength."
- Hobby: "reading", Option: "Fiction" → "Your partner checked in with reading and chose fiction."

BAD EXAMPLES (DO NOT GENERATE):
- "Your partner had an amazing painting session!" (exaggeration)
- "Your partner is becoming a great artist!" (identity claim)
- "Your partner felt inspired while painting!" (emotional claim)
- "Your partner painted for 2 hours!" (inferred time)

Respond with ONLY a JSON object:
{
  "echo": "Your privacy-safe message here"
}`;

/**
 * Creates a HobbyEchoGenerator instance
 * 
 * @param aiClient - The AI client to use for generation
 * @returns A generate function
 */
export function createHobbyEchoGenerator(aiClient: AIClient) {
    /**
     * Generates a privacy-safe progress update
     * 
     * @param input - Hobby and selected option from check-in
     * @returns Echo message
     */
    async function generate(input: HobbyEchoInput): Promise<HobbyEchoOutput> {
        const { hobby, selected_option } = input;

        const userPrompt = `Generate an echo for:
Hobby: "${hobby}"
Selected option: "${selected_option}"`;

        const response = await aiClient.complete<HobbyEchoOutput>({
            systemPrompt: SYSTEM_PROMPT,
            userPrompt,
            temperature: 0.2, // Low for consistency
        });

        // Validate response against schema
        const parsed = hobbyEchoOutputSchema.safeParse(response.data);

        if (!parsed.success) {
            // Fallback to safe generic echo
            return {
                echo: `Your partner checked in with ${hobby}.`,
            };
        }

        return parsed.data;
    }

    return { generate };
}

export type HobbyEchoGenerator = ReturnType<typeof createHobbyEchoGenerator>;
