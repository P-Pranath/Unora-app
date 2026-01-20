// ============================================================================
// Summary Routes - Personality Summary Generation
// ============================================================================
// Handles generating human-readable personality summaries using LLM.
// 
// ENDPOINTS:
// GET /summary - Generate a fresh personality summary
// ============================================================================

import { Router, Request, Response } from 'express';
import prisma from '../db/prisma';
import { generateSummary } from '../summary/summaryGenerator';
import {
    PersonalityDimension,
    DimensionState,
    SummaryResponse,
    ErrorResponse
} from '../types';

const router = Router();

// ============================================================================
// GET /summary
// ============================================================================
// Generates a fresh personality summary based on the user's current state.
// Summaries are NOT cached - they are regenerated on each request.
//
// Query Parameters:
// - userId (required): The user's ID
//
// Response:
// {
//   "summary": "Tends to stay grounded during tense moments...",
//   "generatedAt": "2024-01-15T10:30:00Z",
//   "confidence": {
//     "overall": 0.65,
//     "dimensions": {
//       "emotional_regulation": 0.7,
//       ...
//     }
//   }
// }
// ============================================================================

router.get('/', async (req: Request, res: Response) => {
    try {
        const { userId } = req.query;

        // Validate input
        if (!userId || typeof userId !== 'string') {
            const error: ErrorResponse = {
                error: 'userId query parameter is required',
                code: 'INVALID_INPUT'
            };
            return res.status(400).json(error);
        }

        // Get user's personality state
        const user = await prisma.userProfile.findUnique({
            where: { id: userId },
            include: {
                personalityState: true
            }
        });

        if (!user) {
            const error: ErrorResponse = {
                error: 'User not found. Please initialize profile first.',
                code: 'USER_NOT_FOUND'
            };
            return res.status(404).json(error);
        }

        // Transform database state to DimensionState array
        const states: DimensionState[] = user.personalityState.map(ps => ({
            dimension: ps.dimension as PersonalityDimension,
            score: ps.score,
            confidence: ps.confidence
        }));

        // Check if user has answered any questions
        if (user.questionsAnswered === 0) {
            const response: SummaryResponse = {
                summary: "We haven't learned enough about you yet. Answer a few questions to receive your personality summary.",
                generatedAt: new Date().toISOString(),
                confidence: {
                    overall: 0.1,
                    dimensions: states.reduce((acc, s) => {
                        acc[s.dimension] = s.confidence;
                        return acc;
                    }, {} as Record<PersonalityDimension, number>)
                }
            };
            return res.json(response);
        }

        // Generate the summary using LLM
        const summary = await generateSummary(states);

        // Calculate overall confidence
        const overallConfidence = states.reduce((sum, s) => sum + s.confidence, 0) / states.length;

        // Build dimension confidences map
        const dimensionConfidences: Record<PersonalityDimension, number> = {} as Record<PersonalityDimension, number>;
        for (const state of states) {
            dimensionConfidences[state.dimension] = state.confidence;
        }

        const response: SummaryResponse = {
            summary,
            generatedAt: new Date().toISOString(),
            confidence: {
                overall: Math.round(overallConfidence * 100) / 100,
                dimensions: dimensionConfidences
            }
        };

        return res.json(response);

    } catch (error) {
        console.error('Error generating summary:', error);
        const errorResponse: ErrorResponse = {
            error: 'Failed to generate summary',
            code: 'INTERNAL_ERROR'
        };
        return res.status(500).json(errorResponse);
    }
});

export default router;
