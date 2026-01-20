# Vision Validation Service - Developer Integration Guide

## Service Overview

| Property | Value |
|----------|-------|
| Base URL | `http://localhost:3000` (local) / `https://vision.unora.app` (prod) |
| Protocol | HTTP REST |
| Content-Type | `multipart/form-data` (upload) / `application/json` (response) |
| Authentication | None required (internal service) |

---

## Quick Integration

```typescript
// Minimal integration example
async function validatePhotos(files: File[]): Promise<boolean> {
  const form = new FormData();
  files.forEach(f => form.append('photos', f));

  const res = await fetch(`${VISION_SERVICE_URL}/onboarding/photos`, {
    method: 'POST',
    body: form
  });

  const data = await res.json();
  return data.aggregate.status === 'pass';
}
```

---

## API Reference

### POST `/onboarding/photos`

Validate 3-6 photos for onboarding usability.

**Request:**
```
POST /onboarding/photos?user_id=<uuid>
Content-Type: multipart/form-data

Field: photos (File[], 3-6 required)
```

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `photos` | File[] | Yes | 3-6 image files |
| `user_id` | UUID | No | Optional user identifier |

**Response (200):**
```json
{
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "photos": [
    {
      "photo_id": "a1b2c3d4-...",
      "usable": true,
      "has_face": true,
      "issues": []
    },
    {
      "photo_id": "e5f6g7h8-...",
      "usable": false,
      "has_face": false,
      "issues": ["blurred", "no_face"]
    }
  ],
  "aggregate": {
    "usable_photos": 1,
    "face_photos": 1,
    "status": "needs_more_photos"
  }
}
```

**Error Responses:**

| Status | Error Code | Cause |
|--------|------------|-------|
| 400 | `no_files` | No photos uploaded |
| 400 | `insufficient_photos` | Less than 3 photos |
| 400 | `too_many_files` | More than 6 photos |
| 400 | `file_too_large` | File exceeds 10MB |
| 400 | `invalid_file_type` | Unsupported format |

### Internal Error Codes (Logs)

The service logs structured errors for debugging. Check server logs for these codes:

| Error Code | Description | Common Cause |
|------------|-------------|--------------|
| `UPLOAD_FAILED` | File upload processing failed | Malformed multipart request |
| `CV_FAILED` | Computer vision analysis failed | Corrupted image file |
| `CV_BLUR_CHECK_FAILED` | Blur detection error | Image format issue |
| `CV_BRIGHTNESS_CHECK_FAILED` | Brightness analysis error | Unsupported color space |
| `CV_RESOLUTION_CHECK_FAILED` | Resolution check error | Missing metadata |
| `AI_FAILED` | AI validation failed | API timeout/error |
| `AI_CLIENT_NOT_READY` | AI client not initialized | Missing API key |
| `AI_RESPONSE_INVALID` | AI response parsing failed | Unexpected response format |
| `STORAGE_FAILED` | Storage operation failed | Disk/permission error |
| `CONFIG_INVALID` | Configuration error | Invalid env vars |

### Log Format

```
[timestamp] [LEVEL] Message {context}
  Error: ErrorName: message
  stack trace (dev only)
```

**Example:**
```
[2024-01-13T10:30:00.000Z] [ERROR] AI_VALIDATION_ERROR: AI analysis failed {"errorCode":"AI_FAILED","client":"OpenAI"}
  Error: Error: Request timeout
  at OpenAIClient.validatePhoto (/src/clients/openai.client.ts:95:15)
```

---

### GET `/onboarding/photos/:user_id`

Retrieve validation metadata for a user.

**Response (200):**
```json
{
  "user_id": "550e8400-...",
  "photos": [
    {
      "photo_id": "a1b2c3d4-...",
      "filename": "photo1.jpg",
      "usable": true,
      "has_face": true,
      "created_at": "2024-01-13T10:30:00.000Z"
    }
  ],
  "total_count": 3
}
```

---

### GET `/onboarding/health`

Health check endpoint.

**Response (200):**
```json
{
  "status": "healthy",
  "service": "vision-validation",
  "timestamp": "2024-01-13T10:30:00.000Z"
}
```

---

## Issue Tags Reference

| Tag | Description | User Message |
|-----|-------------|--------------|
| `no_face` | No human face detected | "Please upload a photo of yourself" |
| `blurred` | Image too blurry | "Photo is blurry, please upload a clearer image" |
| `too_dark` | Underexposed | "Photo is too dark, use better lighting" |
| `too_bright` | Overexposed | "Photo is overexposed" |
| `obstructed_face` | Face covered | "Face appears covered, remove obstructions" |
| `non_photographic` | Screenshot/meme/AI | "Please upload an actual photo" |
| `low_resolution` | Below 640×480 | "Resolution too low" |

---

## Integration Patterns

### Pattern 1: Simple Pass/Fail

```typescript
const isValid = await validatePhotos(files);
if (!isValid) {
  showRetryUI();
}
```

### Pattern 2: Per-Photo Feedback

```typescript
const result = await validatePhotos(files);

result.photos.forEach((photo, i) => {
  if (!photo.usable) {
    markPhotoInvalid(i, photo.issues);
  }
});

if (result.aggregate.status !== 'pass') {
  const needed = 3 - result.aggregate.usable_photos;
  showMessage(`Upload ${needed} more valid photos`);
}
```

### Pattern 3: Retry with Filtering

```typescript
const result = await validatePhotos(files);
const validFiles = files.filter((_, i) => result.photos[i].usable);
const invalidCount = files.length - validFiles.length;

if (result.aggregate.status !== 'pass') {
  promptForMorePhotos(3 - validFiles.length);
}
```

---

## Validation Rules

| Rule | Requirement |
|------|-------------|
| Minimum photos | 3 |
| Maximum photos | 6 |
| Min usable photos | 3 |
| Min face photos | 1 |
| Max file size | 10MB |
| Supported formats | JPG, JPEG, PNG, HEIC |
| Min resolution | 640×480 |

---

## Configuration

### Environment Variables

```env
# Service URL
VISION_SERVICE_URL=http://localhost:3000

# Timeout (recommended)
VISION_SERVICE_TIMEOUT=30000
```

### TypeScript Types

```typescript
interface PhotoValidationResult {
  photo_id: string;
  usable: boolean;
  has_face: boolean;
  issues: (
    | 'no_face'
    | 'blurred'
    | 'too_dark'
    | 'too_bright'
    | 'obstructed_face'
    | 'non_photographic'
    | 'low_resolution'
  )[];
}

interface AggregateResult {
  usable_photos: number;
  face_photos: number;
  status: 'pass' | 'needs_more_photos';
}

interface ValidationResponse {
  user_id: string;
  photos: PhotoValidationResult[];
  aggregate: AggregateResult;
}
```

---

## Error Handling

```typescript
async function validateWithRetry(
  files: File[],
  maxRetries = 3
): Promise<ValidationResponse> {
  for (let i = 0; i < maxRetries; i++) {
    try {
      const res = await fetch(`${URL}/onboarding/photos`, {
        method: 'POST',
        body: createFormData(files),
        signal: AbortSignal.timeout(30000)
      });

      if (!res.ok) {
        const err = await res.json();
        throw new Error(err.message);
      }

      return res.json();
    } catch (e) {
      if (i === maxRetries - 1) throw e;
      await sleep(1000 * (i + 1));
    }
  }
}
```

---

## Testing

### Mock Mode

Set `USE_MOCK_AI=true` for deterministic testing.

Filename patterns for mock responses:

| Pattern | Result |
|---------|--------|
| `*_pass.*` | usable=true |
| `*no_face*` | issue: no_face |
| `*blur*` | issue: blurred |
| `*dark*` | issue: too_dark |

### Health Check

```bash
curl http://localhost:3000/onboarding/health
```

### Upload Test

```bash
curl -X POST http://localhost:3000/onboarding/photos \
  -F "photos=@photo1.jpg" \
  -F "photos=@photo2.jpg" \
  -F "photos=@photo3.jpg"
```
