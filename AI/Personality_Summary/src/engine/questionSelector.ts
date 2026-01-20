// ============================================================================
// Question Selector - Hardened Adaptive Question Selection Engine
// ============================================================================
// Implements production-ready question selection with:
// - Priority scoring: (1 - confidence) * (1 / (1 + answered_count))
// - Multi-tier fallback logic for graceful degradation
// - Weighted random selection with seeded randomness (deterministic per user/day)
// - Structured logging for debuggability
// ============================================================================

import prisma from '../db/prisma';
import {
    PersonalityDimension,
    PERSONALITY_DIMENSIONS,
    Question,
    QuestionMode,
    MODE_CONFIG,
    DimensionState
} from '../types';
import { getAllQuestions } from '../questions/questionBank';

// ============================================================================
// Types
// ============================================================================

interface SelectionContext {
    userId: string;
    mode: QuestionMode;
    answeredQuestionIds: string[];
    personalityState: DimensionState[];
    lastDimensionAsked: string | null;
    questionsAnswered: number;
    /**
     * Count of answered questions per dimension.
     * Used for priority scoring to deprioritize over-explored dimensions.
     */
    answeredCountByDimension: Record<string, number>;
}

interface SelectionResult {
    question: Question | null;
    complete: boolean;
    reason?: string;
}

/**
 * Priority score computed for each dimension.
 * Higher priority = should be explored more urgently.
 */
interface DimensionPriority {
    dimension: PersonalityDimension;
    priority: number;
    confidence: number;
    answeredCount: number;
}

/**
 * Structured log for selection decisions.
 * Enables debugging and auditing of algorithm behavior.
 */
interface SelectionLog {
    userId: string;
    mode: QuestionMode;
    dimensionPriorities: { dimension: string; priority: number }[];
    topDimensions: string[];
    candidatePoolSize: number;
    fallbackLevel: 0 | 1 | 2 | 3; // 0=primary, 1=expanded, 2=any-except-last, 3=none
    selectedQuestionId: string | null;
    reason: string;
}

// ============================================================================
// Constants
// ============================================================================

/**
 * Number of top-priority dimensions to select from.
 * Using 2 provides variety while maintaining focus on under-understood areas.
 */
const TOP_N_DIMENSIONS = 2;

/**
 * Expanded pool size for first-level fallback.
 * Doubles the initial pool size for broader selection.
 */
const EXPANDED_N_DIMENSIONS = 4;

// ============================================================================
// Seeded Random Number Generator
// ============================================================================

/**
 * Simple seeded PRNG (Linear Congruential Generator).
 * Produces deterministic sequences given the same seed.
 * 
 * WHY SEEDED RANDOMNESS:
 * - Provides controlled randomness that is reproducible for debugging
 * - Same user on same day gets consistent (but varied) behavior
 * - Eliminates purely chaotic question ordering
 */
function createSeededRandom(seed: number): () => number {
    let state = seed;
    return () => {
        // LCG parameters (Numerical Recipes)
        state = (state * 1664525 + 1013904223) >>> 0;
        return state / 0xFFFFFFFF;
    };
}

/**
 * Creates a daily seed for a user.
 * Combines userId hash with current date for deterministic daily behavior.
 */
function createDailySeed(userId: string): number {
    const today = new Date();
    const dateString = `${today.getFullYear()}-${today.getMonth()}-${today.getDate()}`;
    const combined = userId + dateString;

    // Simple string hash (djb2)
    let hash = 5381;
    for (let i = 0; i < combined.length; i++) {
        hash = ((hash << 5) + hash) + combined.charCodeAt(i);
        hash = hash >>> 0; // Convert to unsigned 32-bit
    }
    return hash;
}

// ============================================================================
// Core Selection Logic
// ============================================================================

/**
 * Selects the next question for a user using a hardened, priority-based algorithm.
 * 
 * ALGORITHM OVERVIEW:
 * 1. Check max questions per mode → return null if reached
 * 2. Fetch all unanswered questions for the user
 * 3. Compute priority score for each dimension: (1 - confidence) * (1 / (1 + answered_count))
 * 4. Sort dimensions by priority (descending)
 * 5. Exclude last-asked dimension from top position if possible
 * 6. Select top N dimensions (N=2)
 * 7. Filter unanswered questions targeting at least one of these dimensions
 * 8. Apply multi-tier fallback if no matches
 * 9. Perform weighted random selection
 * 10. Return selected question
 * 
 * @param userId - User ID to select question for
 * @param mode - ONBOARDING or STREAK_CHECKIN
 * @returns Selection result with question or completion status
 */
export async function selectNextQuestion(
    userId: string,
    mode: QuestionMode = 'ONBOARDING'
): Promise<SelectionResult> {
    // Build context with all data needed for selection
    const context = await buildSelectionContext(userId, mode);

    // -------------------------------------------------------------------------
    // STEP 1: Check if user has reached max questions for this mode
    // -------------------------------------------------------------------------
    const maxQuestions = MODE_CONFIG[mode].maxQuestions;
    if (context.questionsAnswered >= maxQuestions) {
        logSelection({
            userId,
            mode,
            dimensionPriorities: [],
            topDimensions: [],
            candidatePoolSize: 0,
            fallbackLevel: 0,
            selectedQuestionId: null,
            reason: `Max questions reached for ${mode} mode (${maxQuestions})`
        });
        return {
            question: null,
            complete: true,
            reason: `Reached maximum questions for ${mode} mode (${maxQuestions})`
        };
    }

    // -------------------------------------------------------------------------
    // STEP 2: Get all unanswered questions
    // -------------------------------------------------------------------------
    const allQuestions = getAllQuestions();
    const unansweredQuestions = allQuestions.filter(
        q => !context.answeredQuestionIds.includes(q.id)
    );

    if (unansweredQuestions.length === 0) {
        logSelection({
            userId,
            mode,
            dimensionPriorities: [],
            topDimensions: [],
            candidatePoolSize: 0,
            fallbackLevel: 0,
            selectedQuestionId: null,
            reason: 'All questions have been answered'
        });
        return {
            question: null,
            complete: true,
            reason: 'All questions have been answered'
        };
    }

    // -------------------------------------------------------------------------
    // STEP 3-9: Select best question using priority algorithm
    // -------------------------------------------------------------------------
    const selectedQuestion = selectNextQuestionForUser(unansweredQuestions, context);

    if (!selectedQuestion.question) {
        return {
            question: null,
            complete: true,
            reason: selectedQuestion.reason || 'No suitable question found'
        };
    }

    return {
        question: selectedQuestion.question,
        complete: false
    };
}

/**
 * Core question selection algorithm implementing all priority and fallback logic.
 * 
 * WHY THIS APPROACH:
 * - Priority scoring balances confidence gaps AND exploration diversity
 * - Controlled randomness prevents deterministic repetition while remaining debuggable
 * - Multi-tier fallback ensures graceful degradation instead of abrupt failures
 */
function selectNextQuestionForUser(
    unansweredQuestions: Question[],
    context: SelectionContext
): { question: Question | null; reason: string } {
    // -------------------------------------------------------------------------
    // STEP 3: Compute priority scores for each dimension
    // -------------------------------------------------------------------------
    // Priority formula: (1 - confidence) * (1 / (1 + answered_count))
    // 
    // WHY THIS FORMULA:
    // - (1 - confidence): Higher priority for dimensions we're uncertain about
    // - (1 / (1 + answered_count)): Diminishing priority for over-explored dimensions
    // - Combined: Balances exploring uncertain areas while preventing over-concentration
    const dimensionPriorities = computeDimensionPriorities(context);

    // -------------------------------------------------------------------------
    // STEP 4: Sort dimensions by priority (descending)
    // -------------------------------------------------------------------------
    const sortedDimensions = [...dimensionPriorities].sort(
        (a, b) => b.priority - a.priority
    );

    // -------------------------------------------------------------------------
    // STEP 5: Exclude last-asked dimension from top position if possible
    // -------------------------------------------------------------------------
    // WHY: Prevents asking about same dimension twice in a row, improving UX
    const adjustedDimensions = adjustForLastAskedDimension(
        sortedDimensions,
        context.lastDimensionAsked
    );

    // Log priority scores for debugging
    const priorityLog = adjustedDimensions.map(d => ({
        dimension: d.dimension,
        priority: Math.round(d.priority * 1000) / 1000
    }));

    // -------------------------------------------------------------------------
    // STEP 6-8: Multi-tier selection with fallback
    // -------------------------------------------------------------------------
    let candidateQuestions: Question[] = [];
    let fallbackLevel: 0 | 1 | 2 | 3 = 0;
    let topDimensions: string[] = [];

    // TIER 0: Top N dimensions (N=2)
    topDimensions = adjustedDimensions.slice(0, TOP_N_DIMENSIONS).map(d => d.dimension);
    candidateQuestions = filterQuestionsByDimensions(unansweredQuestions, topDimensions);

    // TIER 1: Expanded to top 4 dimensions
    if (candidateQuestions.length === 0) {
        fallbackLevel = 1;
        topDimensions = adjustedDimensions.slice(0, EXPANDED_N_DIMENSIONS).map(d => d.dimension);
        candidateQuestions = filterQuestionsByDimensions(unansweredQuestions, topDimensions);
    }

    // TIER 2: All dimensions except last-asked
    if (candidateQuestions.length === 0) {
        fallbackLevel = 2;
        topDimensions = adjustedDimensions
            .filter(d => d.dimension !== context.lastDimensionAsked)
            .map(d => d.dimension);
        candidateQuestions = filterQuestionsByDimensions(unansweredQuestions, topDimensions);
    }

    // TIER 3: No questions available
    if (candidateQuestions.length === 0) {
        fallbackLevel = 3;
        logSelection({
            userId: context.userId,
            mode: context.mode,
            dimensionPriorities: priorityLog,
            topDimensions,
            candidatePoolSize: 0,
            fallbackLevel,
            selectedQuestionId: null,
            reason: 'No questions match any fallback tier'
        });
        return {
            question: null,
            reason: 'No suitable questions available after all fallbacks'
        };
    }

    // -------------------------------------------------------------------------
    // STEP 9: Weighted random selection
    // -------------------------------------------------------------------------
    // WHY WEIGHTED SELECTION:
    // - Questions targeting higher-priority dimensions should be more likely
    // - Avoids purely random selection that ignores priority scores
    // - Seeded randomness ensures reproducibility for debugging
    const selectedQuestion = weightedRandomSelect(
        candidateQuestions,
        adjustedDimensions,
        context.userId
    );

    // Log the selection decision
    logSelection({
        userId: context.userId,
        mode: context.mode,
        dimensionPriorities: priorityLog,
        topDimensions,
        candidatePoolSize: candidateQuestions.length,
        fallbackLevel,
        selectedQuestionId: selectedQuestion.id,
        reason: `Selected via weighted random (fallback level ${fallbackLevel})`
    });

    return {
        question: selectedQuestion,
        reason: `Selected from pool of ${candidateQuestions.length} candidates`
    };
}

// ============================================================================
// Priority Computation
// ============================================================================

/**
 * Computes priority score for each personality dimension.
 * 
 * FORMULA: priority = (1 - confidence) * (1 / (1 + answered_count))
 * 
 * Example scenarios:
 * - confidence=0.2, answered=0 → priority = 0.8 * 1.0 = 0.8 (HIGH)
 * - confidence=0.2, answered=5 → priority = 0.8 * 0.17 = 0.13 (LOWER)
 * - confidence=0.8, answered=0 → priority = 0.2 * 1.0 = 0.2 (LOW)
 */
function computeDimensionPriorities(context: SelectionContext): DimensionPriority[] {
    return context.personalityState.map(state => {
        const answeredCount = context.answeredCountByDimension[state.dimension] || 0;

        // Core priority formula
        const confidenceFactor = 1 - state.confidence;
        const explorationFactor = 1 / (1 + answeredCount);
        const priority = confidenceFactor * explorationFactor;

        return {
            dimension: state.dimension,
            priority,
            confidence: state.confidence,
            answeredCount
        };
    });
}

/**
 * Adjusts dimension ordering to prevent asking about same dimension consecutively.
 * If the top-priority dimension was just asked, swap it with the second.
 */
function adjustForLastAskedDimension(
    sortedDimensions: DimensionPriority[],
    lastAsked: string | null
): DimensionPriority[] {
    if (!lastAsked || sortedDimensions.length < 2) {
        return sortedDimensions;
    }

    // If the highest-priority dimension was just asked, swap with second
    if (sortedDimensions[0].dimension === lastAsked) {
        const adjusted = [...sortedDimensions];
        [adjusted[0], adjusted[1]] = [adjusted[1], adjusted[0]];
        return adjusted;
    }

    return sortedDimensions;
}

// ============================================================================
// Question Filtering
// ============================================================================

/**
 * Filters questions that target at least one of the specified dimensions.
 */
function filterQuestionsByDimensions(
    questions: Question[],
    targetDimensions: string[]
): Question[] {
    return questions.filter(q =>
        q.dimensionTargets.some(target => targetDimensions.includes(target))
    );
}

// ============================================================================
// Weighted Random Selection
// ============================================================================

/**
 * Performs weighted random selection from candidate questions.
 * 
 * WEIGHTING LOGIC:
 * - Each question's weight = sum of priorities of its targeted dimensions
 * - Higher-priority dimension targets → higher chance of selection
 * - Uses seeded random for reproducibility (same user + same day = consistent behavior)
 */
function weightedRandomSelect(
    candidates: Question[],
    dimensionPriorities: DimensionPriority[],
    userId: string
): Question {
    // Create priority lookup map
    const priorityMap = new Map<string, number>();
    for (const dp of dimensionPriorities) {
        priorityMap.set(dp.dimension, dp.priority);
    }

    // Compute weight for each candidate question
    const weightedCandidates = candidates.map(question => {
        let weight = 0;
        for (const target of question.dimensionTargets) {
            weight += priorityMap.get(target) || 0;
        }
        // Ensure minimum weight to give all candidates some chance
        weight = Math.max(weight, 0.01);
        return { question, weight };
    });

    // Calculate cumulative weights
    const totalWeight = weightedCandidates.reduce((sum, c) => sum + c.weight, 0);

    // Create seeded random generator for this user + day
    const random = createSeededRandom(createDailySeed(userId));
    const randomValue = random() * totalWeight;

    // Select based on weighted random value
    let cumulative = 0;
    for (const candidate of weightedCandidates) {
        cumulative += candidate.weight;
        if (randomValue <= cumulative) {
            return candidate.question;
        }
    }

    // Fallback to last candidate (should not happen if weights are correct)
    return weightedCandidates[weightedCandidates.length - 1].question;
}

// ============================================================================
// Logging
// ============================================================================

/**
 * Logs selection decisions in a structured format for debugging.
 * In production, this could be sent to a logging service.
 */
function logSelection(log: SelectionLog): void {
    const logEntry = {
        timestamp: new Date().toISOString(),
        type: 'QUESTION_SELECTION',
        ...log
    };

    // Use console.info for structured logging (can be intercepted by log aggregators)
    console.info('[QuestionSelector]', JSON.stringify(logEntry, null, 2));
}

// ============================================================================
// Context Building
// ============================================================================

/**
 * Build the complete context needed for question selection.
 * Gathers all user data including personality state and answer history.
 */
async function buildSelectionContext(
    userId: string,
    mode: QuestionMode
): Promise<SelectionContext> {
    // Get user profile with related data
    const userProfile = await prisma.userProfile.findUnique({
        where: { id: userId },
        include: {
            answerLogs: {
                select: {
                    questionId: true
                }
            },
            personalityState: true
        }
    });

    if (!userProfile) {
        throw new Error(`User not found: ${userId}`);
    }

    // Map personality state to DimensionState array
    const personalityState: DimensionState[] = userProfile.personalityState.map(ps => ({
        dimension: ps.dimension as PersonalityDimension,
        score: ps.score,
        confidence: ps.confidence
    }));

    // Get answered question IDs
    const answeredQuestionIds = userProfile.answerLogs.map(log => log.questionId);

    // Count answered questions per dimension
    const allQuestions = getAllQuestions();
    const answeredCountByDimension: Record<string, number> = {};

    // Initialize counts for all dimensions
    for (const dim of PERSONALITY_DIMENSIONS) {
        answeredCountByDimension[dim] = 0;
    }

    // Count answered questions for each dimension they target
    for (const questionId of answeredQuestionIds) {
        const question = allQuestions.find(q => q.id === questionId);
        if (question) {
            for (const dim of question.dimensionTargets) {
                answeredCountByDimension[dim] = (answeredCountByDimension[dim] || 0) + 1;
            }
        }
    }

    return {
        userId,
        mode,
        answeredQuestionIds,
        personalityState,
        lastDimensionAsked: userProfile.lastDimensionAsked,
        questionsAnswered: userProfile.questionsAnswered,
        answeredCountByDimension
    };
}

// ============================================================================
// Exported Utility Functions (Unchanged API)
// ============================================================================

/**
 * Update the last dimension asked for a user
 */
export async function updateLastDimension(
    userId: string,
    dimension: string
): Promise<void> {
    await prisma.userProfile.update({
        where: { id: userId },
        data: { lastDimensionAsked: dimension }
    });
}

/**
 * Get the current personality state for a user
 */
export async function getPersonalityState(
    userId: string
): Promise<DimensionState[]> {
    const states = await prisma.personalityState.findMany({
        where: { userId }
    });

    return states.map(s => ({
        dimension: s.dimension as PersonalityDimension,
        score: s.score,
        confidence: s.confidence
    }));
}

/**
 * Get the count of questions answered by a user
 */
export async function getQuestionsAnswered(userId: string): Promise<number> {
    const user = await prisma.userProfile.findUnique({
        where: { id: userId },
        select: { questionsAnswered: true }
    });

    return user?.questionsAnswered ?? 0;
}

/**
 * Calculate overall confidence as average of all dimension confidences
 */
export function calculateOverallConfidence(states: DimensionState[]): number {
    if (states.length === 0) return 0;
    const sum = states.reduce((acc, s) => acc + s.confidence, 0);
    return sum / states.length;
}
