import sharp from 'sharp';
import { config } from '../config/env';
import { CVValidationResult, CVIssueTag } from '../schemas/validation.schema';
import { logger } from '../utils/logger';

/**
 * CV Validation Service
 * 
 * Performs deterministic computer vision checks on photos:
 * - Blur detection (Laplacian variance approximation)
 * - Brightness/darkness check
 * - Resolution validation
 * 
 * These checks run BEFORE AI validation to filter obviously bad images.
 */

interface ImageMetrics {
    blur_score: number;
    brightness: number;
    width: number;
    height: number;
}

/**
 * Calculate blur score using edge detection variance
 * Lower score = more blurry
 * 
 * Uses a simplified approach with Sharp:
 * 1. Convert to grayscale
 * 2. Apply Laplacian-like edge detection
 * 3. Calculate variance of edge response
 */
async function calculateBlurScore(imagePath: string): Promise<number> {
    try {
        // Load and convert to grayscale
        const image = sharp(imagePath);

        // Apply Laplacian kernel for edge detection
        // Kernel approximates: [0, -1, 0], [-1, 4, -1], [0, -1, 0]
        const edgeDetected = await image
            .greyscale()
            .convolve({
                width: 3,
                height: 3,
                kernel: [0, -1, 0, -1, 4, -1, 0, -1, 0]
            })
            .raw()
            .toBuffer({ resolveWithObject: true });

        const { data, info } = edgeDetected;

        // Calculate variance of edge response
        const pixels = Array.from(data as Buffer);
        const mean = pixels.reduce((a, b) => a + b, 0) / pixels.length;
        const variance = pixels.reduce((sum, val) => sum + Math.pow(val - mean, 2), 0) / pixels.length;

        return variance;
    } catch (error) {
        logger.cvValidationError(error as Error, { check: 'blur_score' });
        // Return low score on error (treat as potentially blurry)
        return 0;
    }
}

/**
 * Calculate average brightness of an image
 * Returns value 0-255
 */
async function calculateBrightness(imagePath: string): Promise<number> {
    try {
        const stats = await sharp(imagePath)
            .greyscale()
            .stats();

        // Return mean brightness from grayscale channel
        return stats.channels[0].mean;
    } catch (error) {
        logger.cvValidationError(error as Error, { check: 'brightness' });
        return 128; // Default to mid-range on error
    }
}

/**
 * Get image dimensions
 */
async function getImageDimensions(imagePath: string): Promise<{ width: number; height: number }> {
    try {
        const metadata = await sharp(imagePath).metadata();
        return {
            width: metadata.width || 0,
            height: metadata.height || 0
        };
    } catch (error) {
        logger.cvValidationError(error as Error, { check: 'dimensions' });
        return { width: 0, height: 0 };
    }
}

/**
 * Collect all image metrics
 */
async function collectMetrics(imagePath: string): Promise<ImageMetrics> {
    const [blur_score, brightness, dimensions] = await Promise.all([
        calculateBlurScore(imagePath),
        calculateBrightness(imagePath),
        getImageDimensions(imagePath)
    ]);

    return {
        blur_score,
        brightness,
        width: dimensions.width,
        height: dimensions.height
    };
}

/**
 * Validate image against CV thresholds
 */
function evaluateMetrics(metrics: ImageMetrics): { passed: boolean; issues: CVIssueTag[] } {
    const issues: CVIssueTag[] = [];
    const { cv } = config;

    // Check blur
    if (metrics.blur_score < cv.blurThreshold) {
        issues.push('blurred');
    }

    // Check brightness
    if (metrics.brightness < cv.brightnessMin) {
        issues.push('too_dark');
    } else if (metrics.brightness > cv.brightnessMax) {
        issues.push('too_bright');
    }

    // Check resolution
    if (metrics.width < cv.minWidth || metrics.height < cv.minHeight) {
        issues.push('low_resolution');
    }

    return {
        passed: issues.length === 0,
        issues
    };
}

/**
 * Main CV validation function
 * 
 * @param imagePath - Path to the image file
 * @returns CVValidationResult with pass/fail status, issues, and metrics
 */
export async function validateCVChecks(imagePath: string): Promise<CVValidationResult> {
    const metrics = await collectMetrics(imagePath);
    const evaluation = evaluateMetrics(metrics);

    return {
        passed: evaluation.passed,
        issues: evaluation.issues,
        metrics
    };
}

/**
 * Get human-readable CV validation summary
 */
export function getCVValidationSummary(result: CVValidationResult): string {
    if (result.passed) {
        return 'Image passed all quality checks';
    }

    const issueDescriptions: Record<CVIssueTag, string> = {
        'blurred': `Blur score ${result.metrics.blur_score.toFixed(0)} below threshold`,
        'too_dark': `Brightness ${result.metrics.brightness.toFixed(0)} too low`,
        'too_bright': `Brightness ${result.metrics.brightness.toFixed(0)} too high`,
        'low_resolution': `Resolution ${result.metrics.width}x${result.metrics.height} below minimum`
    };

    return result.issues.map(issue => issueDescriptions[issue]).join('; ');
}
