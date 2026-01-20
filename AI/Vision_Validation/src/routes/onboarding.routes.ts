import { Router } from 'express';
import { uploadPhotos as uploadPhotosMiddleware } from '../middleware/upload.middleware';
import { uploadPhotos, getPhotos, healthCheck } from '../controllers/photo.controller';

/**
 * Onboarding Routes
 * 
 * API endpoints for photo validation during onboarding.
 */
const router = Router();

/**
 * POST /onboarding/photos
 * 
 * Upload 3-6 photos for validation.
 * 
 * Request:
 * - Content-Type: multipart/form-data
 * - Field: photos (array of files)
 * - Query: user_id (optional, UUID)
 * 
 * Response:
 * {
 *   user_id: string,
 *   photos: Array<{
 *     photo_id: string,
 *     usable: boolean,
 *     has_face: boolean,
 *     issues: string[]
 *   }>,
 *   aggregate: {
 *     usable_photos: number,
 *     face_photos: number,
 *     status: 'pass' | 'needs_more_photos'
 *   }
 * }
 */
router.post('/photos', uploadPhotosMiddleware, uploadPhotos);

/**
 * GET /onboarding/photos/:user_id
 * 
 * Get photo metadata for a user (debugging/testing).
 * Returns metadata only, no image bytes.
 * 
 * Response:
 * {
 *   user_id: string,
 *   photos: Array<{
 *     photo_id: string,
 *     filename: string,
 *     usable: boolean,
 *     has_face: boolean,
 *     created_at: string
 *   }>,
 *   total_count: number
 * }
 */
router.get('/photos/:user_id', getPhotos);

/**
 * GET /onboarding/health
 * 
 * Health check endpoint
 */
router.get('/health', healthCheck);

export default router;
