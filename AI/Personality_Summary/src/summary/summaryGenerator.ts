// ============================================================================
// Summary Generator - LLM-Powered Personality Summary
// ============================================================================
// Generates a third-person personality summary for Unora discovery cards.
// 
// DESIGN PRINCIPLES:
// - Third person only ("They tend to...", "This person often...")
// - One paragraph, 50-90 words
// - Calm, neutral, non-judgmental, human tone
// - No labels, numbers, or psychological terms
// - No absolutes (always, never, definitely)
// - No confidence/data/analysis mentioned
// - Sounds like a thoughtful friend's observation
// ============================================================================

import { AIClient } from '../ai/ai_client';
import {
    PersonalityDimension,
    DimensionState
} from '../types';

// ============================================================================
// Summary Configuration
// ============================================================================

const SUMMARY_CONFIG = {
    // Generation parameters
    temperature: 0.7,
    maxTokens: 150,

    // Validation - word count based (50-90 words)
    minWords: 50,
    maxWords: 90,

    // Forbidden words and phrases (trigger regeneration if found)
    forbiddenWords: [
        // Personality labels
        'introvert', 'extrovert', 'ambivert',
        'type a', 'type b', 'personality type',
        'calm', 'anxious', 'avoidant', 'secure', 'mature', 'stable',
        // Clinical/psychological terms
        'diagnosis', 'disorder', 'therapy', 'clinical',
        'neurotic', 'psychotic', 'attachment style',
        'test results', 'assessment', 'evaluation',
        // Data references
        'score', 'rating', 'percentage', '%', 'data',
        'confidence', 'analysis', 'measured', 'indicates',
        'model', 'ai', 'algorithm',
        // Absolutes
        'always', 'never', 'definitely', 'certainly', 'absolutely',
        // Judgments
        'good', 'bad', 'better', 'worse', 'perfect', 'ideal',
        // Second person (must be third person)
        'you are', 'you tend', 'you often', 'you prefer', 'your ',
        // Comparisons
        'more than others', 'less than others', 'compared to',
        // Emojis and formatting
        '•', '1.', '2.'
    ]
};

// ============================================================================
// Dimension Interpretation
// ============================================================================

// Human-readable interpretations of dimension scores (third-person)
const DIMENSION_INTERPRETATIONS: Record<PersonalityDimension, { low: string; high: string }> = {
    emotional_regulation: {
        low: 'experiences emotions deeply and may need time to process them',
        high: 'tends to stay grounded and processes emotions smoothly'
    },
    communication_style: {
        low: 'often prefers indirect communication and reading between the lines',
        high: 'tends to communicate directly and say what they mean'
    },
    emotional_availability: {
        low: 'takes time to open up and values their privacy',
        high: 'is often emotionally open and accessible to others'
    },
    consistency_style: {
        low: 'enjoys spontaneity and adapts easily to change',
        high: 'values routine and predictability in their day'
    },
    conflict_posture: {
        low: 'is inclined to seek harmony and avoid confrontation',
        high: 'tends to address disagreements directly when they arise'
    },
    energy_orientation: {
        low: 'recharges through solitude and smaller gatherings',
        high: 'often gains energy from social interaction'
    },
    decision_pace: {
        low: 'takes time to deliberate before making decisions',
        high: 'tends to decide quickly and trust their instincts'
    }
};

// Structural shapes for paragraph variation
type ParagraphShape = 'INTERNAL_FIRST' | 'BEHAVIOR_FIRST' | 'CONTEXTUAL_FIRST';

/**
 * Select paragraph shape based on dominant dimension
 */
function selectParagraphShape(primaryDimension: PersonalityDimension): ParagraphShape {
    // Internal-first: dimensions about processing
    if (['emotional_regulation', 'decision_pace'].includes(primaryDimension)) {
        return 'INTERNAL_FIRST';
    }
    // Contextual-first: dimensions about response to situations
    if (['conflict_posture', 'consistency_style'].includes(primaryDimension)) {
        return 'CONTEXTUAL_FIRST';
    }
    // Behavior-first: dimensions about outward expression
    return 'BEHAVIOR_FIRST';
}

/**
 * Get confidence level category
 */
function getConfidenceLevel(confidence: number): 'high' | 'medium' | 'low' {
    if (confidence >= 0.6) return 'high';
    if (confidence >= 0.35) return 'medium';
    return 'low';
}

// ============================================================================
// Prompt Building
// ============================================================================

/**
 * Build the prompt for the LLM based on personality state
 * Implements uniqueness-focused generation with structural variation
 */
function buildPrompt(states: DimensionState[]): string {
    // Sort by confidence to identify primary and secondary dimensions
    const sortedStates = [...states]
        .filter(s => s.confidence >= 0.2)
        .sort((a, b) => b.confidence - a.confidence);

    // Identify primary (center of gravity) and secondary dimensions
    const primaryState = sortedStates[0];
    const secondaryState = sortedStates[1];
    const tertiaryState = sortedStates[2];

    // Select paragraph shape based on primary dimension
    const shape = primaryState ? selectParagraphShape(primaryState.dimension) : 'BEHAVIOR_FIRST';

    // Build dimension context with score direction
    const getDimensionContext = (state: DimensionState) => {
        const interp = DIMENSION_INTERPRETATIONS[state.dimension];
        const position = state.score < 0.4 ? 'low' : state.score > 0.6 ? 'high' : 'balanced';
        const confLevel = getConfidenceLevel(state.confidence);
        return {
            dimension: state.dimension.replace(/_/g, ' '),
            tendency: position === 'low' ? interp.low : position === 'high' ? interp.high : 'balanced',
            position,
            confidenceLevel: confLevel
        };
    };

    const primary = primaryState ? getDimensionContext(primaryState) : null;
    const secondary = secondaryState ? getDimensionContext(secondaryState) : null;
    const tertiary = tertiaryState ? getDimensionContext(tertiaryState) : null;

    const shapeInstructions = {
        'INTERNAL_FIRST': 'Start with how they process situations internally, then move outward to behavior.',
        'BEHAVIOR_FIRST': 'Start with how they act or show up, then briefly hint at internal drivers.',
        'CONTEXTUAL_FIRST': 'Start with how they tend to respond in specific situations (tension, uncertainty, familiarity).'
    };

    return `You are generating a third-person personality summary for a real person on Unora, an intentional, commitment-first connection platform.

This summary is visible to other users inside a discovery card detail view. It is NOT shown to the person being described.

Your primary responsibility is UNIQUENESS.
Each summary must be structurally and semantically unique to the input state.
If two summaries could be swapped between users without feeling wrong, you have failed.

────────────────────────────────
PERSONALITY STATE
────────────────────────────────

PRIMARY DIMENSION (center of gravity):
${primary ? `- ${primary.dimension}: ${primary.tendency} [confidence: ${primary.confidenceLevel}]` : '- insufficient data'}

SECONDARY DIMENSION (qualifying shade):
${secondary ? `- ${secondary.dimension}: ${secondary.tendency} [confidence: ${secondary.confidenceLevel}]` : '- none'}

${tertiary ? `TERTIARY DIMENSION (subtle texture):\n- ${tertiary.dimension}: ${tertiary.tendency} [confidence: ${tertiary.confidenceLevel}]` : ''}

────────────────────────────────
PARAGRAPH SHAPE: ${shape}
────────────────────────────────
${shapeInstructions[shape]}

────────────────────────────────
UNIQUENESS REQUIREMENTS (MANDATORY)
────────────────────────────────

1. CENTER OF GRAVITY: The primary dimension must dominate the tone and opening
2. SECONDARY SHADE: Use secondary dimension to qualify or soften the primary
3. NON-GENERIC STRUCTURE: Do NOT follow "They tend to X. They also Y. They value Z."
4. CONFIDENCE-SENSITIVE LANGUAGE:
   - high confidence → clearer, steadier phrasing
   - medium confidence → balanced, observational phrasing  
   - low confidence → tentative or open-ended phrasing
5. SPECIFICITY CHECK: Before finalizing, ask "Could this describe a random thoughtful adult?" If yes → rewrite.

────────────────────────────────
STRUCTURAL CONSTRAINTS
────────────────────────────────

- ONE paragraph only
- 50-90 words
- Third person ONLY (they / this person)
- No repeated sentence starters
- No mirrored sentence lengths
- Avoid rhythmic symmetry

────────────────────────────────
LANGUAGE CONSTRAINTS (STRICT)
────────────────────────────────

YOU MUST:
- Use observational, grounded language
- Use probabilistic phrasing: "tends to", "often prefers", "is inclined to", "shows a pattern of", "usually approaches"
- Sound like an impression formed over time

YOU MUST NOT:
- Use second person ("you")
- Use labels (anxious, calm, mature, avoidant, secure, stable, etc.)
- Use clinical or psychological terminology
- Use numbers, scores, metrics, or rankings
- Use absolutes ("always", "never", "definitely")
- Compare them to other people
- Mention data, confidence, models, or AI
- Use emojis, bullet points, headings, or lists

────────────────────────────────
FINAL SELF-AUDIT (SILENT)
────────────────────────────────

Before outputting, verify:
- The paragraph would feel wrong if assigned to a different personality state
- The dominant dimension is clearly felt, not just mentioned
- The paragraph does NOT feel "balanced" or "neutral"
- The tone feels human, calm, and earned

────────────────────────────────
OUTPUT FORMAT
────────────────────────────────

Return ONLY the paragraph text.
No titles. No explanations. No metadata.

Write the summary now:`;
}

// ============================================================================
// Generation and Validation
// ============================================================================

/**
 * Generate a personality summary for a user
 * 
 * @param states - The user's personality state across all dimensions
 * @returns The generated summary string
 */
export async function generateSummary(states: DimensionState[]): Promise<string> {
    // Check if we have enough confidence to generate a meaningful summary
    const avgConfidence = states.reduce((sum, s) => sum + s.confidence, 0) / states.length;

    if (avgConfidence < 0.15) {
        return generateLowConfidenceSummary();
    }

    // Build the prompt
    const prompt = buildPrompt(states);

    try {
        const systemPrompt = 'You are a thoughtful observer writing calm, third-person personality descriptions. You sound like a wise friend who has known someone for a while.';

        // Call the LLM via AIClient
        const summary = await AIClient.generateText(systemPrompt, prompt);

        // Validate the summary
        const validationResult = validateSummary(summary);

        if (!validationResult.valid) {
            console.warn('Summary validation failed:', validationResult.reason);
            // Try one more time with stricter prompt
            return await regenerateSummary(states);
        }

        return summary;

    } catch (error) {
        console.error('Error generating summary:', error);
        return generateFallbackSummary(states);
    }
}

/**
 * Regenerate summary with stricter constraints
 */
async function regenerateSummary(states: DimensionState[]): Promise<string> {
    const prompt = buildPrompt(states) + '\n\nCRITICAL REMINDER: UNIQUENESS IS MANDATORY. Use THIRD PERSON ONLY. 50-90 words. The primary dimension must dominate. NO labels, NO absolutes, NO "you".';

    try {
        const systemPrompt = 'You write exclusively in third person. You never use "you" or second person. You are careful and precise.';

        // Call the LLM via AIClient
        const summary = await AIClient.generateText(systemPrompt, prompt);

        // If still invalid, use fallback
        if (!validateSummary(summary).valid) {
            return generateFallbackSummary(states);
        }

        return summary;

    } catch (error) {
        return generateFallbackSummary(states);
    }
}

/**
 * Count words in a string
 */
function countWords(text: string): number {
    return text.trim().split(/\s+/).filter(word => word.length > 0).length;
}

/**
 * Validate that the summary meets our requirements
 */
function validateSummary(summary: string): { valid: boolean; reason?: string } {
    // Check word count (40-80 words)
    const wordCount = countWords(summary);

    if (wordCount < SUMMARY_CONFIG.minWords) {
        return { valid: false, reason: `Summary too short: ${wordCount} words (min: ${SUMMARY_CONFIG.minWords})` };
    }

    if (wordCount > SUMMARY_CONFIG.maxWords) {
        return { valid: false, reason: `Summary too long: ${wordCount} words (max: ${SUMMARY_CONFIG.maxWords})` };
    }

    // Check for forbidden words
    const lowerSummary = summary.toLowerCase();
    for (const word of SUMMARY_CONFIG.forbiddenWords) {
        if (lowerSummary.includes(word.toLowerCase())) {
            return { valid: false, reason: `Contains forbidden word/phrase: "${word}"` };
        }
    }

    // Check it's not multiple paragraphs
    const paragraphs = summary.split(/\n\n+/).filter(p => p.trim().length > 0);
    if (paragraphs.length > 1) {
        return { valid: false, reason: 'Contains multiple paragraphs' };
    }

    // Check for second person usage (additional check)
    if (lowerSummary.match(/\byou\b/)) {
        return { valid: false, reason: 'Contains second person ("you")' };
    }

    return { valid: true };
}

/**
 * Generate a summary when confidence is too low
 */
function generateLowConfidenceSummary(): string {
    return "This person is still getting to know themselves through the platform. " +
        "As they answer more questions, a clearer picture of their natural tendencies will emerge. " +
        "Early signs suggest a thoughtful approach to connection.";
}

/**
 * Generate a fallback summary when LLM fails (third-person, 40-80 words)
 */
function generateFallbackSummary(states: DimensionState[]): string {
    // Find the most confident dimensions
    const sortedStates = [...states].sort((a, b) => b.confidence - a.confidence);
    const topStates = sortedStates.slice(0, 3);

    const traits: string[] = [];

    for (const state of topStates) {
        if (state.confidence < 0.3) continue;

        const interp = DIMENSION_INTERPRETATIONS[state.dimension];
        if (state.score < 0.4) {
            traits.push(interp.low);
        } else if (state.score > 0.6) {
            traits.push(interp.high);
        }
    }

    if (traits.length === 0) {
        return generateLowConfidenceSummary();
    }

    // Helper to clean trait text for sentence construction
    const cleanTrait = (trait: string): string => {
        return trait
            .replace(/^(tends to |often |is inclined to |is often )/, '')
            .replace(/^is /, 'are '); // Fix "they is" -> "they are"
    };

    // Construct a simple fallback in third person
    let summary = `This person ${traits[0]}. `;
    if (traits[1]) {
        summary += `They also ${cleanTrait(traits[1])}. `;
    }
    if (traits[2]) {
        summary += `Generally, they ${cleanTrait(traits[2])}. `;
    }
    summary += "These tendencies reflect early observations and may evolve over time.";

    return summary;
}

// ============================================================================
// Exports
// ============================================================================

export { validateSummary, generateLowConfidenceSummary, generateFallbackSummary };
