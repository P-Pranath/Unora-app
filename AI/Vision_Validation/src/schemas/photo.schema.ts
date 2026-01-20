import { z } from 'zod';

/**
 * Supported MIME types for photo uploads
 */
export const SUPPORTED_MIMETYPES = [
    'image/jpeg',
    'image/png',
    'image/heic',
    'image/heif'
] as const;

export type SupportedMimetype = typeof SUPPORTED_MIMETYPES[number];

/**
 * Photo metadata schema
 */
export const PhotoMetadataSchema = z.object({
    photo_id: z.string().uuid(),
    user_id: z.string().uuid(),
    filename: z.string().min(1),
    original_filename: z.string().min(1),
    mimetype: z.enum(SUPPORTED_MIMETYPES),
    size: z.number().int().positive(),
    visibility: z.literal('private'),
    created_at: z.string().datetime(),
    storage_path: z.string().min(1)
});

export type PhotoMetadata = z.infer<typeof PhotoMetadataSchema>;

/**
 * Photo upload request validation
 * - user_id is optional (mocked for testing)
 * - 3-6 photos required
 */
export const PhotoUploadQuerySchema = z.object({
    user_id: z.string().uuid().optional()
});

/**
 * File validation constraints
 */
export const FILE_CONSTRAINTS = {
    minPhotos: 3,
    maxPhotos: 6,
    maxFileSize: 10 * 1024 * 1024, // 10 MB
    allowedExtensions: ['.jpg', '.jpeg', '.png', '.heic', '.heif']
} as const;

/**
 * Validate file extension
 */
export function isValidFileExtension(filename: string): boolean {
    const ext = filename.toLowerCase().substring(filename.lastIndexOf('.'));
    return FILE_CONSTRAINTS.allowedExtensions.includes(ext as any);
}

/**
 * Validate MIME type
 */
export function isValidMimetype(mimetype: string): mimetype is SupportedMimetype {
    return SUPPORTED_MIMETYPES.includes(mimetype as SupportedMimetype);
}
