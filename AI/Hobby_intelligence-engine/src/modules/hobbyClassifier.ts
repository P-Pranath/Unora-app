/**
 * Hobby Classifier Module
 * 
 * Classifies a validated hobby into a fixed ontology.
 * Determines category, effort type, and practice cadence.
 * 
 * Categories (exactly one):
 * - Physical: Sports, exercise, outdoor activities
 * - Creative: Art, music, crafts, writing
 * - Skill / Practice: Technical skills, instruments, languages
 * - Learning: Academic study, courses, research
 * - Mindfulness: Meditation, yoga, journaling
 * - Social: Group activities, games, community
 * - Collection: Collecting items, curating
 * - Maintenance: Home improvement, gardening, organization
 */

import { AIClient } from '../ai/client.js';
import {
    HobbyClassifierInput,
    HobbyClassifierOutput,
    hobbyClassifierOutputSchema
} from '../types/schemas.js';

const SYSTEM_PROMPT = `You are a hobby classification system. Classify the given hobby into EXACTLY ONE category and determine its characteristics.

CATEGORIES (choose exactly one):
- Physical: Sports, exercise, fitness, outdoor physical activities (gym, running, swimming, hiking)
- Creative: Art, music, crafts, writing, design (painting, guitar, pottery, creative writing)
- Skill / Practice: Technical skills, instruments, languages, craftsmanship (coding, lockpicking, woodworking)
- Learning: Academic study, courses, research, reading for knowledge (studying history, online courses)
- Mindfulness: Meditation, yoga, journaling, contemplative practices (meditation, breathwork)
- Social: Group activities, games, community involvement (board games, volunteering, meetups)
- Collection: Collecting items, curating, organizing collections (stamp collecting, vinyl records)
- Maintenance: Home improvement, gardening, car care, organization (gardening, home repair)

EFFORT TYPE:
- solo: Primarily done alone (reading, painting, running)
- social: Primarily done with others (team sports, board games)
- mixed: Can be done either way (gym, music)

CADENCE:
- daily: Typically practiced daily (meditation, exercise)
- irregular: Practiced on specific occasions (hiking, concerts)
- flexible: No fixed schedule, varies by person (reading, crafts)

CLASSIFICATION RULES:
- Choose the MOST PRIMARY category if hobby spans multiple
- Consider the core nature of the activity
- Be consistent - same hobby should always get same classification

Respond with ONLY a JSON object in this exact format:
{
  "category": "Physical" | "Creative" | "Skill / Practice" | "Learning" | "Mindfulness" | "Social" | "Collection" | "Maintenance",
  "effort_type": "solo" | "social" | "mixed",
  "cadence": "daily" | "irregular" | "flexible"
}`;

/**
 * Creates a HobbyClassifier instance
 * 
 * @param aiClient - The AI client to use for classification
 * @returns A classify function
 */
export function createHobbyClassifier(aiClient: AIClient) {
    /**
     * Classifies a hobby into category, effort type, and cadence
     * 
     * @param input - The hobby to classify
     * @returns Classification result
     */
    async function classify(input: HobbyClassifierInput): Promise<HobbyClassifierOutput> {
        const userPrompt = `Classify this hobby: "${input.hobby}"`;

        const response = await aiClient.complete<HobbyClassifierOutput>({
            systemPrompt: SYSTEM_PROMPT,
            userPrompt,
            temperature: 0.1, // Low for consistent classification
        });

        // Validate response against schema
        const parsed = hobbyClassifierOutputSchema.safeParse(response.data);

        if (!parsed.success) {
            // Default fallback classification
            return {
                category: 'Skill / Practice',
                effort_type: 'solo',
                cadence: 'flexible',
            };
        }

        return parsed.data;
    }

    return { classify };
}

export type HobbyClassifier = ReturnType<typeof createHobbyClassifier>;
