import { z } from 'zod';

/**
 * Issue tags - standardized across the system
 */
export const ISSUE_TAGS = [
    'no_face',
    'blurred',
    'too_dark',
    'too_bright',
    'obstructed_face',
    'non_photographic',
    'low_resolution'
] as const;

export type IssueTag = typeof ISSUE_TAGS[number];

/**
 * CV-specific issue tags
 */
export const CV_ISSUE_TAGS = [
    'blurred',
    'too_dark',
    'too_bright',
    'low_resolution'
] as const;

export type CVIssueTag = typeof CV_ISSUE_TAGS[number];

/**
 * Confidence levels
 */
export const CONFIDENCE_LEVELS = ['high', 'medium', 'low'] as const;
export type Confidence = typeof CONFIDENCE_LEVELS[number];

/**
 * CV Validation result schema
 */
export const CVValidationResultSchema = z.object({
    passed: z.boolean(),
    issues: z.array(z.enum(CV_ISSUE_TAGS)),
    metrics: z.object({
        blur_score: z.number(),
        brightness: z.number(),
        width: z.number().int().positive(),
        height: z.number().int().positive()
    })
});

export type CVValidationResult = z.infer<typeof CVValidationResultSchema>;

/**
 * AI Validation result schema (ephemeral - NOT to be stored)
 */
export const AIValidationResultSchema = z.object({
    usable: z.boolean(),
    issues: z.array(z.enum(ISSUE_TAGS)),
    confidence: z.enum(CONFIDENCE_LEVELS)
});

export type AIValidationResult = z.infer<typeof AIValidationResultSchema>;

/**
 * Per-photo validation result schema (stored)
 * CRITICAL: Only final usability flags - no biometric data
 */
export const PhotoValidationResultSchema = z.object({
    photo_id: z.string().uuid(),
    usable: z.boolean(),
    has_face: z.boolean(),
    issues: z.array(z.enum(ISSUE_TAGS))
    // DO NOT ADD: confidence, bounding_boxes, embeddings, face_count
});

export type PhotoValidationResult = z.infer<typeof PhotoValidationResultSchema>;

/**
 * Aggregate result schema
 */
export const AggregateResultSchema = z.object({
    usable_photos: z.number().int().min(0),
    face_photos: z.number().int().min(0),
    status: z.enum(['pass', 'needs_more_photos'])
});

export type AggregateResult = z.infer<typeof AggregateResultSchema>;

/**
 * Complete upload response schema
 */
export const PhotoUploadResponseSchema = z.object({
    user_id: z.string().uuid(),
    photos: z.array(PhotoValidationResultSchema),
    aggregate: AggregateResultSchema
});

export type PhotoUploadResponse = z.infer<typeof PhotoUploadResponseSchema>;

/**
 * User-friendly error messages for issues
 */
export const ISSUE_MESSAGES: Record<IssueTag, string> = {
    'no_face': 'We couldn\'t detect a face in your photo. Please upload a clear photo of yourself.',
    'blurred': 'Your photo appears blurry. Please upload a clearer image.',
    'too_dark': 'Your photo is too dark. Please use better lighting.',
    'too_bright': 'Your photo is overexposed. Please reduce the brightness.',
    'obstructed_face': 'Your face appears to be covered. Please remove any obstructions.',
    'non_photographic': 'Please upload an actual photo, not a screenshot or drawing.',
    'low_resolution': 'Your photo resolution is too low. Please upload a higher quality image.'
};

/**
 * Get user-friendly message for issues
 */
export function getIssueMessage(issues: IssueTag[]): string {
    if (issues.length === 0) return '';
    return issues.map(issue => ISSUE_MESSAGES[issue]).join(' ');
}
