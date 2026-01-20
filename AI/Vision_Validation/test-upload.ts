/**
 * Test script for Vision Validation Service
 * Run with: npx ts-node test-upload.ts
 */

import fs from 'fs';
import path from 'path';
import http from 'http';

const API_HOST = 'localhost';
const API_PORT = 3000;

interface ValidationResponse {
    user_id: string;
    photos: Array<{
        photo_id: string;
        usable: boolean;
        has_face: boolean;
        issues: string[];
    }>;
    aggregate: {
        usable_photos: number;
        face_photos: number;
        status: 'pass' | 'needs_more_photos';
    };
}

/**
 * Upload photos using raw HTTP multipart form
 */
async function uploadPhotos(filePaths: string[]): Promise<ValidationResponse> {
    return new Promise((resolve, reject) => {
        const boundary = '----FormBoundary' + Math.random().toString(16).slice(2);
        const parts: Buffer[] = [];

        // Build multipart body
        for (const filePath of filePaths) {
            const absolutePath = path.resolve(filePath);
            if (!fs.existsSync(absolutePath)) {
                reject(new Error(`File not found: ${absolutePath}`));
                return;
            }

            const filename = path.basename(absolutePath);
            const fileData = fs.readFileSync(absolutePath);

            parts.push(Buffer.from(
                `--${boundary}\r\n` +
                `Content-Disposition: form-data; name="photos"; filename="${filename}"\r\n` +
                `Content-Type: image/png\r\n\r\n`
            ));
            parts.push(fileData);
            parts.push(Buffer.from('\r\n'));
        }

        parts.push(Buffer.from(`--${boundary}--\r\n`));
        const body = Buffer.concat(parts);

        const options: http.RequestOptions = {
            hostname: API_HOST,
            port: API_PORT,
            path: '/onboarding/photos',
            method: 'POST',
            headers: {
                'Content-Type': `multipart/form-data; boundary=${boundary}`,
                'Content-Length': body.length
            }
        };

        const req = http.request(options, (res) => {
            let data = '';
            res.on('data', (chunk) => { data += chunk; });
            res.on('end', () => {
                if (res.statusCode !== 200) {
                    reject(new Error(`Upload failed: ${res.statusCode} - ${data}`));
                    return;
                }
                try {
                    resolve(JSON.parse(data));
                } catch (e) {
                    reject(new Error(`Failed to parse response: ${data}`));
                }
            });
        });

        req.on('error', reject);
        req.write(body);
        req.end();
    });
}

async function getPhotoMetadata(userId: string): Promise<any> {
    return new Promise((resolve, reject) => {
        http.get(`http://${API_HOST}:${API_PORT}/onboarding/photos/${userId}`, (res) => {
            let data = '';
            res.on('data', (chunk) => { data += chunk; });
            res.on('end', () => {
                try {
                    resolve(JSON.parse(data));
                } catch (e) {
                    reject(new Error(`Failed to parse response: ${data}`));
                }
            });
        }).on('error', reject);
    });
}

async function testHealthCheck(): Promise<void> {
    console.log('\n📋 Testing Health Check...');
    return new Promise((resolve, reject) => {
        http.get(`http://${API_HOST}:${API_PORT}/onboarding/health`, (res) => {
            let data = '';
            res.on('data', (chunk) => { data += chunk; });
            res.on('end', () => {
                console.log('   Response:', data);
                resolve();
            });
        }).on('error', reject);
    });
}

async function testValidUpload(): Promise<string> {
    console.log('\n📤 Testing Valid Photo Upload (3 valid photos)...');

    try {
        const result = await uploadPhotos([
            'test-images/valid_face_pass.png',
            'test-images/valid_face_pass.png',
            'test-images/valid_face_pass.png'
        ]);

        console.log('   User ID:', result.user_id);
        console.log('   Photos:');
        for (const photo of result.photos) {
            console.log(`     - ${photo.photo_id.substring(0, 8)}... usable=${photo.usable}, has_face=${photo.has_face}`);
        }
        console.log('   Aggregate:', JSON.stringify(result.aggregate));
        console.log('   Status:', result.aggregate.status === 'pass' ? '✅ PASS' : '⚠️ NEEDS MORE PHOTOS');

        return result.user_id;
    } catch (error) {
        console.error('   Error:', error);
        return '';
    }
}

async function testMixedUpload(): Promise<void> {
    console.log('\n📤 Testing Mixed Photo Upload (valid + no_face + blur)...');

    try {
        const result = await uploadPhotos([
            'test-images/valid_face_pass.png',
            'test-images/no_face_test.png',
            'test-images/blur_test.png'
        ]);

        console.log('   User ID:', result.user_id);
        console.log('   Results:');
        for (const photo of result.photos) {
            const status = photo.usable ? '✓' : '✗';
            console.log(`     ${status} usable=${photo.usable}, has_face=${photo.has_face}, issues=[${photo.issues.join(', ')}]`);
        }
        console.log('   Aggregate:', JSON.stringify(result.aggregate));
        console.log('   Status:', result.aggregate.status === 'pass' ? '✅ PASS' : '⚠️ NEEDS MORE PHOTOS');
    } catch (error) {
        console.error('   Error:', error);
    }
}

async function testInsufficientPhotos(): Promise<void> {
    console.log('\n📤 Testing Insufficient Photos (only 2)...');

    try {
        const result = await uploadPhotos([
            'test-images/valid_face_pass.png',
            'test-images/valid_face_pass.png'
        ]);
        console.log('   Expected error but got:', result);
    } catch (error: any) {
        if (error.message.includes('insufficient_photos') || error.message.includes('at least')) {
            console.log('   ✅ Correctly rejected:', error.message.substring(0, 100));
        } else {
            console.log('   ⚠️ Error:', error.message.substring(0, 100));
        }
    }
}

async function main(): Promise<void> {
    console.log('═'.repeat(60));
    console.log('  Vision Validation Service - Test Suite');
    console.log('═'.repeat(60));

    try {
        await testHealthCheck();
        const userId = await testValidUpload();
        await testMixedUpload();
        await testInsufficientPhotos();

        if (userId) {
            console.log('\n📋 Testing Get Photo Metadata...');
            const metadata = await getPhotoMetadata(userId);
            console.log('   User ID:', metadata.user_id);
            console.log('   Total photos:', metadata.total_count);
        }

        console.log('\n' + '═'.repeat(60));
        console.log('  Test Suite Complete');
        console.log('═'.repeat(60));
    } catch (error) {
        console.error('\n❌ Test failed:', error);
        process.exit(1);
    }
}

main();
