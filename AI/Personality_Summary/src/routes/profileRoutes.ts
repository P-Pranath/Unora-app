// ============================================================================
// Profile Routes - User Personality Profile Management
// ============================================================================
// Handles user profile creation and initialization of personality state.
// 
// ENDPOINTS:
// POST /profile/init - Create a new user profile with default personality state
// ============================================================================

import { Router, Request, Response } from 'express';
import prisma from '../db/prisma';
import {
    PERSONALITY_DIMENSIONS,
    InitProfileRequest,
    InitProfileResponse,
    PersonalityState,
    ErrorResponse
} from '../types';

const router = Router();

// ============================================================================
// POST /profile/init
// ============================================================================
// Creates a new user personality profile with default dimension states.
// All dimensions start with score=0.5 (neutral) and confidence=0.1 (low).
//
// Request Body:
// {
//   "userId": "string" // Unique identifier for the user
// }
//
// Response (201 Created):
// {
//   "userId": "...",
//   "personalityState": {
//     "emotional_regulation": { "score": 0.5, "confidence": 0.1 },
//     ...
//   },
//   "message": "Profile created successfully"
// }
// ============================================================================

router.post('/init', async (req: Request, res: Response) => {
    try {
        const { userId } = req.body as InitProfileRequest;

        // Validate input
        if (!userId || typeof userId !== 'string') {
            const error: ErrorResponse = {
                error: 'userId is required and must be a string',
                code: 'INVALID_INPUT'
            };
            return res.status(400).json(error);
        }

        // Check if user already exists
        const existingUser = await prisma.userProfile.findUnique({
            where: { id: userId }
        });

        if (existingUser) {
            const error: ErrorResponse = {
                error: 'User profile already exists',
                code: 'USER_EXISTS'
            };
            return res.status(409).json(error);
        }

        // Create user profile with personality states in a transaction
        const user = await prisma.$transaction(async (tx) => {
            // Create the user profile
            const profile = await tx.userProfile.create({
                data: {
                    id: userId,
                    questionsAnswered: 0,
                    lastDimensionAsked: null
                }
            });

            // Create personality state for each dimension
            for (const dimension of PERSONALITY_DIMENSIONS) {
                await tx.personalityState.create({
                    data: {
                        userId: profile.id,
                        dimension,
                        score: 0.5,      // Neutral starting point
                        confidence: 0.1  // Low confidence initially
                    }
                });
            }

            return profile;
        });

        // Fetch the created personality states
        const states = await prisma.personalityState.findMany({
            where: { userId: user.id }
        });

        // Transform to response format
        const personalityState: PersonalityState = {} as PersonalityState;
        for (const state of states) {
            personalityState[state.dimension as keyof PersonalityState] = {
                score: state.score,
                confidence: state.confidence
            };
        }

        const response: InitProfileResponse = {
            userId: user.id,
            personalityState,
            message: 'Profile created successfully'
        };

        return res.status(201).json(response);

    } catch (error) {
        console.error('Error creating profile:', error);
        const errorResponse: ErrorResponse = {
            error: 'Failed to create profile',
            code: 'INTERNAL_ERROR'
        };
        return res.status(500).json(errorResponse);
    }
});

// ============================================================================
// GET /profile/:userId
// ============================================================================
// Gets the current personality state for a user.
// ============================================================================

router.get('/:userId', async (req: Request, res: Response) => {
    try {
        const { userId } = req.params;

        // Get user profile with personality state
        const user = await prisma.userProfile.findUnique({
            where: { id: userId },
            include: {
                personalityState: true
            }
        });

        if (!user) {
            const error: ErrorResponse = {
                error: 'User not found',
                code: 'USER_NOT_FOUND'
            };
            return res.status(404).json(error);
        }

        // Transform to response format
        const personalityState: PersonalityState = {} as PersonalityState;
        for (const state of user.personalityState) {
            personalityState[state.dimension as keyof PersonalityState] = {
                score: state.score,
                confidence: state.confidence
            };
        }

        return res.json({
            userId: user.id,
            personalityState,
            questionsAnswered: user.questionsAnswered,
            createdAt: user.createdAt,
            updatedAt: user.updatedAt
        });

    } catch (error) {
        console.error('Error getting profile:', error);
        const errorResponse: ErrorResponse = {
            error: 'Failed to get profile',
            code: 'INTERNAL_ERROR'
        };
        return res.status(500).json(errorResponse);
    }
});

export default router;
