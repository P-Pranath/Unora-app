// ============================================================================
// Question Routes - Adaptive Question Serving
// ============================================================================
// Handles retrieving the next question for a user based on their
// personality state and answering history.
// 
// ENDPOINTS:
// GET /question/next - Get the next adaptive question
// ============================================================================

import { Router, Request, Response } from 'express';
import { selectNextQuestion, updateLastDimension } from '../engine/questionSelector';
import {
    QuestionMode,
    NextQuestionResponse,
    QuestionCompleteResponse,
    ErrorResponse
} from '../types';

const router = Router();

// ============================================================================
// GET /question/next
// ============================================================================
// Returns the next question for a user based on their personality state.
// Uses adaptive selection to prioritize low-confidence dimensions.
//
// Query Parameters:
// - userId (required): The user's ID
// - mode (optional): 'ONBOARDING' (default) or 'STREAK_CHECKIN'
//
// Response (when question available):
// {
//   "question": {
//     "id": "Q_ER_01",
//     "scenario": "A plan you were excited about...",
//     "options": [{ "label": "..." }, ...]
//   },
//   "questionsAnswered": 2,
//   "mode": "ONBOARDING"
// }
//
// Response (when complete):
// {
//   "complete": true,
//   "message": "All questions completed for this mode",
//   "questionsAnswered": 8
// }
// ============================================================================

router.get('/next', async (req: Request, res: Response) => {
    try {
        const { userId, mode = 'ONBOARDING' } = req.query;

        // Validate input
        if (!userId || typeof userId !== 'string') {
            const error: ErrorResponse = {
                error: 'userId query parameter is required',
                code: 'INVALID_INPUT'
            };
            return res.status(400).json(error);
        }

        // Validate mode
        const validModes: QuestionMode[] = ['ONBOARDING', 'STREAK_CHECKIN'];
        const questionMode = (mode as string).toUpperCase() as QuestionMode;

        if (!validModes.includes(questionMode)) {
            const error: ErrorResponse = {
                error: `Invalid mode. Must be one of: ${validModes.join(', ')}`,
                code: 'INVALID_MODE'
            };
            return res.status(400).json(error);
        }

        // Select the next question
        const result = await selectNextQuestion(userId, questionMode);

        // If no more questions, return complete status
        if (result.complete || !result.question) {
            const { getQuestionsAnswered } = await import('../engine/questionSelector');
            const questionsAnswered = await getQuestionsAnswered(userId);

            const response: QuestionCompleteResponse = {
                complete: true,
                message: result.reason || 'All questions completed for this mode',
                questionsAnswered
            };
            return res.json(response);
        }

        // Update last dimension asked (use primary dimension target)
        const primaryDimension = result.question.dimensionTargets[0];
        await updateLastDimension(userId, primaryDimension);

        // Get questions answered count
        const { getQuestionsAnswered } = await import('../engine/questionSelector');
        const questionsAnswered = await getQuestionsAnswered(userId);

        // Return the question (hide dimension impacts from client)
        const response: NextQuestionResponse = {
            question: {
                id: result.question.id,
                scenario: result.question.scenario,
                options: result.question.options.map(opt => ({ label: opt.label }))
            },
            questionsAnswered,
            mode: questionMode
        };

        return res.json(response);

    } catch (error) {
        console.error('Error getting next question:', error);

        // Check for specific error types
        if (error instanceof Error && error.message.includes('User not found')) {
            const errorResponse: ErrorResponse = {
                error: 'User not found. Please initialize profile first.',
                code: 'USER_NOT_FOUND'
            };
            return res.status(404).json(errorResponse);
        }

        const errorResponse: ErrorResponse = {
            error: 'Failed to get next question',
            code: 'INTERNAL_ERROR'
        };
        return res.status(500).json(errorResponse);
    }
});

export default router;
