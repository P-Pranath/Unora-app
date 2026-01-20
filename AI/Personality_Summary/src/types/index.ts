// ============================================================================
// TypeScript Type Definitions - Personality Intelligence Service
// ============================================================================
// Central type definitions for personality dimensions, questions, options,
// and API request/response structures.
// ============================================================================

// ============================================================================
// Personality Dimensions
// ============================================================================
// The 7 core personality dimensions tracked by the system.
// Each dimension represents a spectrum, not a binary label.

export const PERSONALITY_DIMENSIONS = [
    'emotional_regulation',    // How one manages emotional responses
    'communication_style',     // Direct vs. indirect communication preferences
    'emotional_availability',  // Openness to emotional connection
    'consistency_style',       // Preference for routine vs. spontaneity
    'conflict_posture',        // Approach to handling disagreements
    'energy_orientation',      // Social energy preferences
    'decision_pace',           // Speed of decision-making
] as const;

export type PersonalityDimension = typeof PERSONALITY_DIMENSIONS[number];

// ============================================================================
// Dimension State
// ============================================================================
// Represents the current state of a single personality dimension.

export interface DimensionState {
    dimension: PersonalityDimension;
    score: number;      // 0.0 - 1.0 (0.5 = neutral)
    confidence: number; // 0.0 - 1.0 (how certain we are)
}

// ============================================================================
// Full Personality State
// ============================================================================
// Complete personality profile with all 7 dimensions.

export type PersonalityState = Record<PersonalityDimension, {
    score: number;
    confidence: number;
}>;

// ============================================================================
// Question Types
// ============================================================================

// Impact on a dimension from selecting an option
export interface DimensionImpact {
    [dimension: string]: number; // e.g., { "emotional_regulation": 0.2 }
}

// A single answer option within a question
export interface QuestionOption {
    label: string;
    impacts: DimensionImpact;
}

// A complete question with scenario and options
export interface Question {
    id: string;
    scenario: string;
    dimensionTargets: PersonalityDimension[];
    options: QuestionOption[];
}

// Question as stored in database (with JSON strings)
export interface QuestionDB {
    id: string;
    scenarioText: string;
    dimensionTargets: string;  // JSON array
    options: string;           // JSON array
}

// ============================================================================
// Question Selection Modes
// ============================================================================

export type QuestionMode = 'ONBOARDING' | 'STREAK_CHECKIN';

// Mode-specific configuration
export const MODE_CONFIG: Record<QuestionMode, { maxQuestions: number }> = {
    ONBOARDING: { maxQuestions: 8 },
    STREAK_CHECKIN: { maxQuestions: 1 },
};

// ============================================================================
// API Request Types
// ============================================================================

// POST /profile/init
export interface InitProfileRequest {
    userId: string;
}

// POST /answer
export interface SubmitAnswerRequest {
    userId: string;
    questionId: string;
    selectedOption?: number; // Index of selected option
    skip?: boolean;          // True if user skipped
}

// ============================================================================
// API Response Types
// ============================================================================

// POST /profile/init response
export interface InitProfileResponse {
    userId: string;
    personalityState: PersonalityState;
    message: string;
}

// GET /question/next response (when question available)
export interface NextQuestionResponse {
    question: {
        id: string;
        scenario: string;
        options: { label: string }[]; // Impacts hidden from client
    };
    questionsAnswered: number;
    mode: QuestionMode;
}

// GET /question/next response (when no more questions)
export interface QuestionCompleteResponse {
    complete: true;
    message: string;
    questionsAnswered: number;
}

// POST /answer response
export interface SubmitAnswerResponse {
    updated: boolean;
    skipped: boolean;
    questionsAnswered: number;
    nextQuestion?: NextQuestionResponse['question'];
}

// GET /summary response
export interface SummaryResponse {
    summary: string;
    generatedAt: string;
    confidence: {
        overall: number;
        dimensions: Record<PersonalityDimension, number>;
    };
}

// Error response
export interface ErrorResponse {
    error: string;
    code?: string;
}

// ============================================================================
// Scoring Types
// ============================================================================

export interface ScoreUpdate {
    dimension: PersonalityDimension;
    oldScore: number;
    newScore: number;
    oldConfidence: number;
    newConfidence: number;
}
