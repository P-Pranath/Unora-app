import multer, { FileFilterCallback, StorageEngine } from 'multer';
import path from 'path';
import fs from 'fs';
import { Request } from 'express';
import { config } from '../config/env';
import { FILE_CONSTRAINTS, isValidMimetype } from '../schemas/photo.schema';
import { generateStorageFilename } from '../utils/helpers';

/**
 * Ensure storage directory exists
 */
function ensureStorageDirectory(): void {
    const storagePath = path.resolve(config.storagePath);
    if (!fs.existsSync(storagePath)) {
        fs.mkdirSync(storagePath, { recursive: true });
        console.log(`📁 Created storage directory: ${storagePath}`);
    }
}

// Ensure storage directory exists on module load
ensureStorageDirectory();

/**
 * Custom storage engine for multer
 */
const storage: StorageEngine = multer.diskStorage({
    destination: (req: Request, file, cb) => {
        const storagePath = path.resolve(config.storagePath);
        cb(null, storagePath);
    },
    filename: (req: Request, file, cb) => {
        const filename = generateStorageFilename(file.originalname);
        cb(null, filename);
    }
});

/**
 * File filter to validate uploads
 */
function fileFilter(
    req: Request,
    file: Express.Multer.File,
    cb: FileFilterCallback
): void {
    // Check MIME type
    if (!isValidMimetype(file.mimetype)) {
        cb(new Error(`Invalid file type: ${file.mimetype}. Allowed: jpg, jpeg, png, heic`));
        return;
    }

    cb(null, true);
}

/**
 * Multer upload configuration
 */
export const upload = multer({
    storage,
    fileFilter,
    limits: {
        fileSize: FILE_CONSTRAINTS.maxFileSize,
        files: FILE_CONSTRAINTS.maxPhotos
    }
});

/**
 * Middleware for handling photo uploads
 * Expects field name 'photos' with 3-6 files
 */
export const uploadPhotos = upload.array('photos', FILE_CONSTRAINTS.maxPhotos);

/**
 * Cleanup uploaded files on error
 */
export function cleanupUploadedFiles(files: Express.Multer.File[]): void {
    for (const file of files) {
        try {
            if (fs.existsSync(file.path)) {
                fs.unlinkSync(file.path);
            }
        } catch (err) {
            console.error(`Failed to cleanup file: ${file.path}`, err);
        }
    }
}
