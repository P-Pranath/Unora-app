// ============================================================================
// Answer Routes - Answer Submission and Score Updates
// ============================================================================
// Handles submitting answers to questions and updating personality scores.
// 
// ENDPOINTS:
// POST /answer - Submit an answer to a question
// ============================================================================

import { Router, Request, Response } from 'express';
import prisma from '../db/prisma';
import { processAnswer } from '../scoring/scorer';
import { selectNextQuestion, updateLastDimension } from '../engine/questionSelector';
import { getQuestionById } from '../questions/questionBank';
import {
    SubmitAnswerRequest,
    SubmitAnswerResponse,
    QuestionMode,
    ErrorResponse
} from '../types';

const router = Router();

// ============================================================================
// POST /answer
// ============================================================================
// Submits an answer to a question and updates the user's personality state.
// Optionally returns the next question.
//
// Request Body:
// {
//   "userId": "user-123",
//   "questionId": "Q_ER_01",
//   "selectedOption": 0,  // Index of selected option (0-based)
//   "skip": false         // Optional: true to skip this question
// }
//
// Response:
// {
//   "updated": true,
//   "skipped": false,
//   "questionsAnswered": 3,
//   "nextQuestion": { ... }  // Optional: next question if available
// }
// ============================================================================

router.post('/', async (req: Request, res: Response) => {
    try {
        const {
            userId,
            questionId,
            selectedOption,
            skip = false
        } = req.body as SubmitAnswerRequest;

        // Validate required fields
        if (!userId || typeof userId !== 'string') {
            const error: ErrorResponse = {
                error: 'userId is required',
                code: 'INVALID_INPUT'
            };
            return res.status(400).json(error);
        }

        if (!questionId || typeof questionId !== 'string') {
            const error: ErrorResponse = {
                error: 'questionId is required',
                code: 'INVALID_INPUT'
            };
            return res.status(400).json(error);
        }

        // Check if user exists
        const user = await prisma.userProfile.findUnique({
            where: { id: userId }
        });

        if (!user) {
            const error: ErrorResponse = {
                error: 'User not found. Please initialize profile first.',
                code: 'USER_NOT_FOUND'
            };
            return res.status(404).json(error);
        }

        // Get the question
        const question = getQuestionById(questionId);

        if (!question) {
            const error: ErrorResponse = {
                error: `Question not found: ${questionId}`,
                code: 'QUESTION_NOT_FOUND'
            };
            return res.status(404).json(error);
        }

        // Check if already answered
        const existingAnswer = await prisma.answerLog.findFirst({
            where: {
                userId,
                questionId
            }
        });

        if (existingAnswer) {
            const error: ErrorResponse = {
                error: 'Question already answered',
                code: 'ALREADY_ANSWERED'
            };
            return res.status(409).json(error);
        }

        // Validate selectedOption if not skipping
        if (!skip) {
            if (selectedOption === undefined || selectedOption === null) {
                const error: ErrorResponse = {
                    error: 'selectedOption is required when not skipping',
                    code: 'INVALID_INPUT'
                };
                return res.status(400).json(error);
            }

            if (selectedOption < 0 || selectedOption >= question.options.length) {
                const error: ErrorResponse = {
                    error: `selectedOption must be between 0 and ${question.options.length - 1}`,
                    code: 'INVALID_OPTION'
                };
                return res.status(400).json(error);
            }
        }

        // Process the answer
        const selectedOpt = skip ? null : question.options[selectedOption!];
        const updates = await processAnswer(userId, questionId, selectedOpt, skip);

        // Update last dimension asked
        if (!skip && question.dimensionTargets.length > 0) {
            await updateLastDimension(userId, question.dimensionTargets[0]);
        }

        // Get updated question count
        const updatedUser = await prisma.userProfile.findUnique({
            where: { id: userId },
            select: { questionsAnswered: true }
        });

        // Try to get the next question
        const nextResult = await selectNextQuestion(userId, 'ONBOARDING');

        const response: SubmitAnswerResponse = {
            updated: updates.length > 0,
            skipped: skip,
            questionsAnswered: updatedUser?.questionsAnswered || 0
        };

        // Include next question if available
        if (!nextResult.complete && nextResult.question) {
            response.nextQuestion = {
                id: nextResult.question.id,
                scenario: nextResult.question.scenario,
                options: nextResult.question.options.map(opt => ({ label: opt.label }))
            };
        }

        return res.json(response);

    } catch (error) {
        console.error('Error submitting answer:', error);
        const errorResponse: ErrorResponse = {
            error: 'Failed to submit answer',
            code: 'INTERNAL_ERROR'
        };
        return res.status(500).json(errorResponse);
    }
});

export default router;
