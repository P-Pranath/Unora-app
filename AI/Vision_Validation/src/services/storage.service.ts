import fs from 'fs';
import path from 'path';
import { config } from '../config/env';
import { PhotoMetadata, PhotoMetadataSchema } from '../schemas/photo.schema';
import { PhotoValidationResult } from '../schemas/validation.schema';
import { generateId, getCurrentTimestamp, normalizeMimetype } from '../utils/helpers';

/**
 * Storage Service
 * 
 * Handles photo storage and metadata management.
 * 
 * CRITICAL STORAGE RULES:
 * ✅ STORED: Raw images, photo_id, user_id, filename, visibility, created_at
 * ✅ STORED: Final usable boolean, has_face boolean
 * ❌ NOT STORED: Face embeddings, bounding boxes, AI confidence, biometric data
 */

// In-memory storage for metadata (in production, use database)
const metadataStore: Map<string, PhotoMetadata[]> = new Map();
const validationStore: Map<string, PhotoValidationResult> = new Map();

/**
 * Ensure storage directory exists
 */
function ensureStorageDir(): string {
    const storagePath = path.resolve(config.storagePath);
    if (!fs.existsSync(storagePath)) {
        fs.mkdirSync(storagePath, { recursive: true });
    }
    return storagePath;
}

/**
 * Create metadata for an uploaded photo
 */
export function createPhotoMetadata(
    file: Express.Multer.File,
    userId: string
): PhotoMetadata {
    const metadata = {
        photo_id: generateId(),
        user_id: userId,
        filename: file.filename,
        original_filename: file.originalname,
        mimetype: normalizeMimetype(file.mimetype),
        size: file.size,
        visibility: 'private' as const,
        created_at: getCurrentTimestamp(),
        storage_path: file.path
    };

    // Validate against schema
    return PhotoMetadataSchema.parse(metadata);
}

/**
 * Save photo metadata
 */
export function savePhotoMetadata(metadata: PhotoMetadata): void {
    const userPhotos = metadataStore.get(metadata.user_id) || [];
    userPhotos.push(metadata);
    metadataStore.set(metadata.user_id, userPhotos);
}

/**
 * Save photo validation result
 * IMPORTANT: Only stores usable/has_face flags, no biometric data
 */
export function saveValidationResult(result: PhotoValidationResult): void {
    validationStore.set(result.photo_id, result);
}

/**
 * Get all photos for a user
 */
export function getPhotosByUserId(userId: string): PhotoMetadata[] {
    return metadataStore.get(userId) || [];
}

/**
 * Get validation result for a photo
 */
export function getValidationResult(photoId: string): PhotoValidationResult | undefined {
    return validationStore.get(photoId);
}

/**
 * Get photo metadata with validation status
 */
export function getPhotoWithValidation(photoId: string): {
    metadata: PhotoMetadata | undefined;
    validation: PhotoValidationResult | undefined;
} {
    // Find metadata across all users
    for (const [userId, photos] of metadataStore.entries()) {
        const photo = photos.find(p => p.photo_id === photoId);
        if (photo) {
            return {
                metadata: photo,
                validation: validationStore.get(photoId)
            };
        }
    }

    return { metadata: undefined, validation: undefined };
}

/**
 * Get all photos with validation for a user
 */
export function getUserPhotosWithValidation(userId: string): Array<{
    photo_id: string;
    filename: string;
    usable: boolean;
    has_face: boolean;
    created_at: string;
}> {
    const photos = getPhotosByUserId(userId);

    return photos.map(photo => {
        const validation = validationStore.get(photo.photo_id);
        return {
            photo_id: photo.photo_id,
            filename: photo.original_filename,
            usable: validation?.usable ?? false,
            has_face: validation?.has_face ?? false,
            created_at: photo.created_at
        };
    });
}

/**
 * Delete photo by ID
 */
export function deletePhoto(photoId: string): boolean {
    const { metadata } = getPhotoWithValidation(photoId);

    if (!metadata) {
        return false;
    }

    // Delete file
    try {
        if (fs.existsSync(metadata.storage_path)) {
            fs.unlinkSync(metadata.storage_path);
        }
    } catch (error) {
        console.error('Failed to delete photo file:', error);
    }

    // Remove from stores
    validationStore.delete(photoId);

    const userPhotos = metadataStore.get(metadata.user_id);
    if (userPhotos) {
        const filtered = userPhotos.filter(p => p.photo_id !== photoId);
        metadataStore.set(metadata.user_id, filtered);
    }

    return true;
}

/**
 * Clear all data for a user
 */
export function clearUserData(userId: string): void {
    const photos = getPhotosByUserId(userId);

    // Delete all files
    for (const photo of photos) {
        try {
            if (fs.existsSync(photo.storage_path)) {
                fs.unlinkSync(photo.storage_path);
            }
            validationStore.delete(photo.photo_id);
        } catch (error) {
            console.error('Failed to delete photo:', error);
        }
    }

    metadataStore.delete(userId);
}

/**
 * Get storage statistics
 */
export function getStorageStats(): {
    totalUsers: number;
    totalPhotos: number;
    totalValidations: number;
} {
    let totalPhotos = 0;
    for (const photos of metadataStore.values()) {
        totalPhotos += photos.length;
    }

    return {
        totalUsers: metadataStore.size,
        totalPhotos,
        totalValidations: validationStore.size
    };
}

/**
 * Check if photo file exists
 */
export function photoFileExists(photoId: string): boolean {
    const { metadata } = getPhotoWithValidation(photoId);
    if (!metadata) return false;
    return fs.existsSync(metadata.storage_path);
}
