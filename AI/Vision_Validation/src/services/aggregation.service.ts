import { PhotoValidationResult, AggregateResult } from '../schemas/validation.schema';

/**
 * Aggregation Service
 * 
 * Aggregates individual photo validation results and applies business rules:
 * - Minimum 3 usable photos required
 * - At least 1 usable face photo required
 */

/**
 * Aggregation requirements
 */
export const AGGREGATION_REQUIREMENTS = {
    minUsablePhotos: 3,
    minFacePhotos: 1
} as const;

/**
 * Aggregate photo validation results
 * 
 * @param results - Array of per-photo validation results
 * @returns Aggregate result with counts and status
 */
export function aggregateResults(results: PhotoValidationResult[]): AggregateResult {
    const usable_photos = results.filter(r => r.usable).length;
    const face_photos = results.filter(r => r.usable && r.has_face).length;

    const meetsRequirements =
        usable_photos >= AGGREGATION_REQUIREMENTS.minUsablePhotos &&
        face_photos >= AGGREGATION_REQUIREMENTS.minFacePhotos;

    return {
        usable_photos,
        face_photos,
        status: meetsRequirements ? 'pass' : 'needs_more_photos'
    };
}

/**
 * Calculate what's still needed to pass
 */
export function calculateDeficit(aggregate: AggregateResult): {
    usable_needed: number;
    face_needed: number;
    message: string;
} {
    const usable_needed = Math.max(0, AGGREGATION_REQUIREMENTS.minUsablePhotos - aggregate.usable_photos);
    const face_needed = Math.max(0, AGGREGATION_REQUIREMENTS.minFacePhotos - aggregate.face_photos);

    let message = '';

    if (usable_needed > 0 && face_needed > 0) {
        message = `Please upload ${usable_needed} more usable photo(s), including at least ${face_needed} with a clear face.`;
    } else if (usable_needed > 0) {
        message = `Please upload ${usable_needed} more usable photo(s).`;
    } else if (face_needed > 0) {
        message = `Please upload at least ${face_needed} photo(s) with a clear face.`;
    }

    return {
        usable_needed,
        face_needed,
        message
    };
}

/**
 * Generate summary of aggregation results
 */
export function getAggregationSummary(
    results: PhotoValidationResult[],
    aggregate: AggregateResult
): string {
    const total = results.length;
    const { usable_photos, face_photos, status } = aggregate;

    if (status === 'pass') {
        return `✅ Validation passed: ${usable_photos}/${total} photos usable, ${face_photos} with face.`;
    }

    const deficit = calculateDeficit(aggregate);
    return `⚠️ ${deficit.message} (Current: ${usable_photos}/${total} usable, ${face_photos} with face)`;
}

/**
 * Group results by status for detailed reporting
 */
export function groupResultsByStatus(results: PhotoValidationResult[]): {
    usable: PhotoValidationResult[];
    unusable: PhotoValidationResult[];
    withFace: PhotoValidationResult[];
    withoutFace: PhotoValidationResult[];
} {
    return {
        usable: results.filter(r => r.usable),
        unusable: results.filter(r => !r.usable),
        withFace: results.filter(r => r.has_face),
        withoutFace: results.filter(r => !r.has_face)
    };
}
