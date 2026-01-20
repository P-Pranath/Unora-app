import { Request, Response, NextFunction } from 'express';
import { z } from 'zod';
import { FILE_CONSTRAINTS, PhotoUploadQuerySchema } from '../schemas/photo.schema';
import {
    PhotoValidationResult,
    PhotoUploadResponse,
    getIssueMessage
} from '../schemas/validation.schema';
import { validateCVChecks, getCVValidationSummary } from '../services/cv.service';
import { validateWithAI, hasFaceFromAIResult } from '../services/ai.service';
import {
    createPhotoMetadata,
    savePhotoMetadata,
    saveValidationResult,
    getPhotosByUserId,
    getUserPhotosWithValidation
} from '../services/storage.service';
import { aggregateResults, getAggregationSummary } from '../services/aggregation.service';
import { AppError } from '../middleware/error.middleware';
import { cleanupUploadedFiles } from '../middleware/upload.middleware';
import { generateId } from '../utils/helpers';

/**
 * Photo Controller
 * 
 * Handles photo upload and validation requests.
 */

/**
 * POST /onboarding/photos
 * 
 * Upload and validate 3-6 photos for onboarding.
 * 
 * Flow:
 * 1. Accept multipart uploads
 * 2. Run CV validation (blur, brightness, resolution)
 * 3. Run AI validation (face presence, photo type) - only if CV passes
 * 4. Aggregate results and return status
 */
export async function uploadPhotos(
    req: Request,
    res: Response,
    next: NextFunction
): Promise<void> {
    try {
        // Parse query params
        const query = PhotoUploadQuerySchema.safeParse(req.query);
        const userId = query.success && query.data.user_id
            ? query.data.user_id
            : generateId(); // Mock user_id for testing

        // Validate file count
        const files = req.files as Express.Multer.File[];

        if (!files || files.length === 0) {
            throw new AppError(
                'No photos uploaded. Please upload 3-6 photos.',
                400,
                'no_files'
            );
        }

        if (files.length < FILE_CONSTRAINTS.minPhotos) {
            cleanupUploadedFiles(files);
            throw new AppError(
                `Please upload at least ${FILE_CONSTRAINTS.minPhotos} photos. You uploaded ${files.length}.`,
                400,
                'insufficient_photos'
            );
        }

        console.log(`📤 Processing ${files.length} photos for user ${userId}`);

        // Process each photo
        const results: PhotoValidationResult[] = [];

        for (const file of files) {
            console.log(`\n📸 Processing: ${file.originalname}`);

            // Create metadata
            const metadata = createPhotoMetadata(file, userId);
            savePhotoMetadata(metadata);

            // Step 1: CV Validation (deterministic checks)
            const cvResult = await validateCVChecks(file.path);
            console.log(`   CV: ${cvResult.passed ? '✓' : '✗'} ${getCVValidationSummary(cvResult)}`);

            let photoResult: PhotoValidationResult;

            if (!cvResult.passed) {
                // Immediate rejection - skip AI validation
                photoResult = {
                    photo_id: metadata.photo_id,
                    usable: false,
                    has_face: false, // Can't determine without AI
                    issues: cvResult.issues
                };
            } else {
                // Step 2: AI Validation (only if CV passes)
                const aiResult = await validateWithAI(file.path, file.originalname);

                photoResult = {
                    photo_id: metadata.photo_id,
                    usable: aiResult.usable,
                    has_face: hasFaceFromAIResult(aiResult),
                    issues: aiResult.issues
                    // Note: confidence is NOT stored (ephemeral)
                };
            }

            // Save validation result (only flags, no biometric data)
            saveValidationResult(photoResult);
            results.push(photoResult);

            console.log(`   Result: usable=${photoResult.usable}, has_face=${photoResult.has_face}, issues=[${photoResult.issues.join(', ')}]`);
        }

        // Step 3: Aggregate results
        const aggregate = aggregateResults(results);
        console.log(`\n${getAggregationSummary(results, aggregate)}`);

        // Build response
        const response: PhotoUploadResponse = {
            user_id: userId,
            photos: results,
            aggregate
        };

        res.status(200).json(response);

    } catch (error) {
        // Cleanup files on error
        if (req.files && Array.isArray(req.files)) {
            cleanupUploadedFiles(req.files as Express.Multer.File[]);
        }
        next(error);
    }
}

/**
 * GET /onboarding/photos/:user_id
 * 
 * Get photo metadata for a user (no image bytes).
 * Used for debugging and testing.
 */
export async function getPhotos(
    req: Request,
    res: Response,
    next: NextFunction
): Promise<void> {
    try {
        const { user_id } = req.params;

        if (!user_id) {
            throw new AppError('User ID is required', 400, 'missing_user_id');
        }

        // Validate UUID format
        const uuidSchema = z.string().uuid();
        const parseResult = uuidSchema.safeParse(user_id);

        if (!parseResult.success) {
            throw new AppError('Invalid user ID format', 400, 'invalid_user_id');
        }

        const photos = getUserPhotosWithValidation(user_id);

        if (photos.length === 0) {
            res.status(200).json({
                user_id,
                photos: [],
                total_count: 0,
                message: 'No photos found for this user'
            });
            return;
        }

        res.status(200).json({
            user_id,
            photos,
            total_count: photos.length
        });

    } catch (error) {
        next(error);
    }
}

/**
 * GET /health
 * 
 * Health check endpoint
 */
export async function healthCheck(
    req: Request,
    res: Response
): Promise<void> {
    res.status(200).json({
        status: 'healthy',
        service: 'vision-validation',
        timestamp: new Date().toISOString()
    });
}
