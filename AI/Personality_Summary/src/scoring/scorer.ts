// ============================================================================
// Scorer - Score and Confidence Update Logic
// ============================================================================
// Implements the scoring algorithm for personality dimension updates.
// 
// SCORING FORMULA:
// newScore = (oldScore * confidence + delta * weight) / (confidence + weight)
// 
// CONFIDENCE FORMULA:
// - Normal: newConfidence = min(confidence + 0.08, 1.0)
// - High confidence (>0.8): newConfidence = min(confidence + 0.02, 1.0) (diminishing returns)
// 
// DESIGN PRINCIPLES:
// - Early answers have more impact (low confidence = high weight)
// - Later answers have less impact (high confidence = stability)
// - Skipping has no penalty (no score/confidence change)
// ============================================================================

import prisma from '../db/prisma';
import {
    PersonalityDimension,
    DimensionImpact,
    ScoreUpdate,
    QuestionOption
} from '../types';

// ============================================================================
// Configuration
// ============================================================================

const SCORING_CONFIG = {
    // Base weight for new answers
    BASE_WEIGHT: 1.0,

    // Reduced weight when confidence is already high
    HIGH_CONFIDENCE_WEIGHT: 0.5,

    // Threshold for high confidence (diminishing returns)
    HIGH_CONFIDENCE_THRESHOLD: 0.8,

    // Normal confidence increment
    NORMAL_CONFIDENCE_INCREMENT: 0.08,

    // Reduced confidence increment for high confidence
    HIGH_CONFIDENCE_INCREMENT: 0.02,

    // Maximum confidence value
    MAX_CONFIDENCE: 1.0,

    // Minimum and maximum score bounds
    MIN_SCORE: 0.0,
    MAX_SCORE: 1.0
};

// ============================================================================
// Core Scoring Functions
// ============================================================================

/**
 * Process an answer and update the user's personality state
 * 
 * @param userId - The user's ID
 * @param questionId - The question that was answered
 * @param selectedOption - The option that was selected
 * @param skip - Whether the user skipped this question
 * @returns Array of score updates that were applied
 */
export async function processAnswer(
    userId: string,
    questionId: string,
    selectedOption: QuestionOption | null,
    skip: boolean = false
): Promise<ScoreUpdate[]> {
    // If skipped, no score updates - just log the skip
    if (skip || !selectedOption) {
        await logAnswer(userId, questionId, null, true);
        await incrementQuestionsAnswered(userId);
        return [];
    }

    // Get the dimension impacts from the selected option
    const impacts = selectedOption.impacts;

    // Apply each dimension impact
    const updates: ScoreUpdate[] = [];

    for (const [dimension, delta] of Object.entries(impacts)) {
        const update = await applyDimensionUpdate(
            userId,
            dimension as PersonalityDimension,
            delta
        );
        updates.push(update);
    }

    // Log the answer
    await logAnswer(userId, questionId, 0, false); // We don't track option index here
    await incrementQuestionsAnswered(userId);

    return updates;
}

/**
 * Apply a score update to a specific dimension
 * 
 * ALGORITHM:
 * 1. Get current score and confidence
 * 2. Calculate weight based on current confidence
 * 3. Apply weighted average formula for new score
 * 4. Increment confidence with diminishing returns
 */
async function applyDimensionUpdate(
    userId: string,
    dimension: PersonalityDimension,
    delta: number
): Promise<ScoreUpdate> {
    // Get current state
    const currentState = await prisma.personalityState.findUnique({
        where: {
            userId_dimension: { userId, dimension }
        }
    });

    if (!currentState) {
        throw new Error(`PersonalityState not found for ${userId} / ${dimension}`);
    }

    const oldScore = currentState.score;
    const oldConfidence = currentState.confidence;

    // Calculate weight (reduced if high confidence)
    const weight = oldConfidence > SCORING_CONFIG.HIGH_CONFIDENCE_THRESHOLD
        ? SCORING_CONFIG.HIGH_CONFIDENCE_WEIGHT
        : SCORING_CONFIG.BASE_WEIGHT;

    // Apply weighted average formula
    // newScore = (oldScore * confidence + delta) / (confidence + weight)
    // But we need to add delta to the score, so:
    // newScore = (oldScore * confidence + (oldScore + delta) * weight) / (confidence + weight)
    // Simplified: move score toward (oldScore + delta) with weight
    const targetScore = oldScore + delta;
    const newScoreRaw = (oldScore * oldConfidence + targetScore * weight) / (oldConfidence + weight);

    // Clamp to valid range
    const newScore = Math.max(
        SCORING_CONFIG.MIN_SCORE,
        Math.min(SCORING_CONFIG.MAX_SCORE, newScoreRaw)
    );

    // Calculate new confidence (with diminishing returns)
    const confidenceIncrement = oldConfidence > SCORING_CONFIG.HIGH_CONFIDENCE_THRESHOLD
        ? SCORING_CONFIG.HIGH_CONFIDENCE_INCREMENT
        : SCORING_CONFIG.NORMAL_CONFIDENCE_INCREMENT;

    const newConfidence = Math.min(
        oldConfidence + confidenceIncrement,
        SCORING_CONFIG.MAX_CONFIDENCE
    );

    // Update database
    await prisma.personalityState.update({
        where: {
            userId_dimension: { userId, dimension }
        },
        data: {
            score: newScore,
            confidence: newConfidence
        }
    });

    return {
        dimension,
        oldScore,
        newScore,
        oldConfidence,
        newConfidence
    };
}

/**
 * Log an answer for auditing and preventing repeat questions
 */
async function logAnswer(
    userId: string,
    questionId: string,
    selectedOption: number | null,
    skipped: boolean
): Promise<void> {
    await prisma.answerLog.create({
        data: {
            userId,
            questionId,
            selectedOption,
            skipped
        }
    });
}

/**
 * Increment the questions answered counter for a user
 */
async function incrementQuestionsAnswered(userId: string): Promise<void> {
    await prisma.userProfile.update({
        where: { id: userId },
        data: {
            questionsAnswered: { increment: 1 }
        }
    });
}

// ============================================================================
// Utility Functions
// ============================================================================

/**
 * Reset a user's personality state to defaults
 * Useful for testing or when user requests a reset
 */
export async function resetPersonalityState(userId: string): Promise<void> {
    await prisma.personalityState.updateMany({
        where: { userId },
        data: {
            score: 0.5,
            confidence: 0.1
        }
    });

    await prisma.userProfile.update({
        where: { id: userId },
        data: {
            questionsAnswered: 0,
            lastDimensionAsked: null
        }
    });

    await prisma.answerLog.deleteMany({
        where: { userId }
    });
}

/**
 * Get the impact description for a delta value
 * Useful for debugging and logging
 */
export function describeImpact(delta: number): string {
    if (delta >= 0.2) return 'strong positive';
    if (delta >= 0.1) return 'moderate positive';
    if (delta > 0) return 'slight positive';
    if (delta === 0) return 'neutral';
    if (delta > -0.1) return 'slight negative';
    if (delta > -0.2) return 'moderate negative';
    return 'strong negative';
}

/**
 * Calculate the expected score after applying a delta
 * without actually modifying the database
 */
export function previewScoreUpdate(
    currentScore: number,
    currentConfidence: number,
    delta: number
): { newScore: number; newConfidence: number } {
    const weight = currentConfidence > SCORING_CONFIG.HIGH_CONFIDENCE_THRESHOLD
        ? SCORING_CONFIG.HIGH_CONFIDENCE_WEIGHT
        : SCORING_CONFIG.BASE_WEIGHT;

    const targetScore = currentScore + delta;
    const newScoreRaw = (currentScore * currentConfidence + targetScore * weight) /
        (currentConfidence + weight);

    const newScore = Math.max(
        SCORING_CONFIG.MIN_SCORE,
        Math.min(SCORING_CONFIG.MAX_SCORE, newScoreRaw)
    );

    const confidenceIncrement = currentConfidence > SCORING_CONFIG.HIGH_CONFIDENCE_THRESHOLD
        ? SCORING_CONFIG.HIGH_CONFIDENCE_INCREMENT
        : SCORING_CONFIG.NORMAL_CONFIDENCE_INCREMENT;

    const newConfidence = Math.min(
        currentConfidence + confidenceIncrement,
        SCORING_CONFIG.MAX_CONFIDENCE
    );

    return { newScore, newConfidence };
}
