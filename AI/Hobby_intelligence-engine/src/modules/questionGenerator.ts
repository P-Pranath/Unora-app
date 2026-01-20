/**
 * Question Generator Module
 * 
 * Generates a single daily check-in question for a hobby.
 * Questions are designed to be answered with 2-4 tap options.
 * 
 * Rules (Non-negotiable):
 * - ONE question only
 * - 2-4 tap options
 * - No "why" questions
 * - No emotions, motivation, advice, or judgment
 * - No time, location, identity, or performance claims
 * - Neutral, factual language
 * - Must avoid repetition with previous_questions
 */

import { AIClient } from '../ai/client.js';
import {
    QuestionGeneratorInput,
    QuestionGeneratorOutput,
    HobbyCategory,
    questionGeneratorOutputSchema
} from '../types/schemas.js';

// Category-specific question patterns to guide AI
const CATEGORY_PATTERNS: Record<HobbyCategory, string[]> = {
    'Physical': [
        'What type of activity did you do?',
        'What was your focus area?',
        'What intensity level?',
    ],
    'Creative': [
        'What did you work on?',
        'What medium did you use?',
        'What stage is your project at?',
    ],
    'Skill / Practice': [
        'What skill did you practice?',
        'What area did you focus on?',
        'What did you work with?',
    ],
    'Learning': [
        'What topic did you study?',
        'What format did you use?',
        'What subject area?',
    ],
    'Mindfulness': [
        'What practice did you do?',
        'What technique did you use?',
        'What was your focus?',
    ],
    'Social': [
        'What type of activity?',
        'What format was it?',
        'What was the setting?',
    ],
    'Collection': [
        'What did you do with your collection?',
        'What category did you focus on?',
        'What action did you take?',
    ],
    'Maintenance': [
        'What area did you work on?',
        'What type of task?',
        'What was the focus?',
    ],
};

const SYSTEM_PROMPT = `You are a check-in question generator. Generate ONE simple question with 2-4 tap options.

STRICT RULES:
1. Generate exactly ONE question
2. Provide 2-4 short tap options (1-3 words each)
3. NO "why" questions
4. NO emotional language (happy, frustrated, motivated, excited)
5. NO judgment or advice
6. NO time references (today, yesterday, this week)
7. NO location references
8. NO identity claims (you are, you're becoming)
9. NO performance assessment (great job, well done)
10. NEUTRAL, factual language only
11. Options should be mutually exclusive but cover common responses

GOOD EXAMPLES:
- Question: "What did you focus on?" Options: ["Technique", "Endurance", "Strength", "Flexibility"]
- Question: "What type of session?" Options: ["Practice", "Learning", "Creating", "Reviewing"]

BAD EXAMPLES (DO NOT GENERATE):
- "Why did you choose this?" (why question)
- "How did it make you feel?" (emotional)
- "Did you have a great session?" (judgment)
- "Where did you practice?" (location)

Respond with ONLY a JSON object:
{
  "question": "Your question here",
  "options": ["Option 1", "Option 2", "Option 3"]
}`;

/**
 * Creates a QuestionGenerator instance
 * 
 * @param aiClient - The AI client to use for generation
 * @returns A generate function
 */
export function createQuestionGenerator(aiClient: AIClient) {
    /**
     * Generates a daily check-in question with tap options
     * 
     * @param input - Hobby, category, and previous questions
     * @returns Generated question with options
     */
    async function generate(input: QuestionGeneratorInput): Promise<QuestionGeneratorOutput> {
        const { hobby, category, previous_questions } = input;

        // Get category-specific patterns for guidance
        const patterns = CATEGORY_PATTERNS[category] || CATEGORY_PATTERNS['Skill / Practice'];

        let userPrompt = `Generate a check-in question for the hobby: "${hobby}" (Category: ${category})

Example patterns for this category:
${patterns.map(p => `- ${p}`).join('\n')}`;

        // Add previous questions to avoid repetition
        if (previous_questions.length > 0) {
            userPrompt += `\n\nAVOID these previously asked questions:
${previous_questions.map(q => `- ${q}`).join('\n')}`;
        }

        const response = await aiClient.complete<QuestionGeneratorOutput>({
            systemPrompt: SYSTEM_PROMPT,
            userPrompt,
            temperature: 0.4, // Slightly higher for variety
        });

        // Validate response against schema
        const parsed = questionGeneratorOutputSchema.safeParse(response.data);

        if (!parsed.success) {
            // Fallback to generic question
            return {
                question: 'What did you focus on?',
                options: ['Practice', 'Learning', 'Creating', 'Other'],
            };
        }

        return parsed.data;
    }

    return { generate };
}

export type QuestionGenerator = ReturnType<typeof createQuestionGenerator>;
