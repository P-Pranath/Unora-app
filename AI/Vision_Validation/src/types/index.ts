/**
 * TypeScript type definitions for Vision Validation Service
 */

// Issue tags - standardized across CV and AI validation
export type IssueTag =
    | 'no_face'
    | 'blurred'
    | 'too_dark'
    | 'too_bright'
    | 'obstructed_face'
    | 'non_photographic'
    | 'low_resolution';

// CV-specific issues (subset of IssueTag)
export type CVIssue = 'blurred' | 'too_dark' | 'too_bright' | 'low_resolution';

// Confidence levels
export type Confidence = 'high' | 'medium' | 'low';

// Aggregate status
export type AggregateStatus = 'pass' | 'needs_more_photos';

// Photo visibility (always private for this service)
export type Visibility = 'private';

/**
 * Photo metadata stored in the system
 */
export interface PhotoMetadata {
    photo_id: string;
    user_id: string;
    filename: string;
    original_filename: string;
    mimetype: string;
    size: number;
    visibility: Visibility;
    created_at: string;
    storage_path: string;
}

/**
 * CV Validation result (deterministic checks)
 */
export interface CVValidationResult {
    passed: boolean;
    issues: CVIssue[];
    metrics: {
        blur_score: number;
        brightness: number;
        width: number;
        height: number;
    };
}

/**
 * AI Validation result (ephemeral - NOT stored)
 */
export interface AIValidationResult {
    usable: boolean;
    issues: IssueTag[];
    confidence: Confidence;
}

/**
 * Per-photo validation result (stored)
 * Note: Only final flags stored, no biometric data
 */
export interface PhotoValidationResult {
    photo_id: string;
    usable: boolean;
    has_face: boolean;
    issues: IssueTag[];
    // DO NOT store: confidence, bounding boxes, embeddings
}

/**
 * Aggregate result for all photos
 */
export interface AggregateResult {
    usable_photos: number;
    face_photos: number;
    status: AggregateStatus;
}

/**
 * Complete response for photo upload endpoint
 */
export interface PhotoUploadResponse {
    user_id: string;
    photos: PhotoValidationResult[];
    aggregate: AggregateResult;
}

/**
 * Photo metadata response (for GET endpoint)
 */
export interface PhotoMetadataResponse {
    user_id: string;
    photos: Array<{
        photo_id: string;
        filename: string;
        usable: boolean;
        has_face: boolean;
        created_at: string;
    }>;
    total_count: number;
}

/**
 * Error response structure
 */
export interface ErrorResponse {
    error: string;
    message: string;
    details?: string[];
}

/**
 * Processing context for a single photo
 */
export interface PhotoProcessingContext {
    file: Express.Multer.File;
    metadata: PhotoMetadata;
    cv_result?: CVValidationResult;
    ai_result?: AIValidationResult;
    final_result?: PhotoValidationResult;
}
