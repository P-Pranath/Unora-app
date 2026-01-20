# Vision Validation Service

AI-powered photo quality and face presence validation for Unora onboarding.

## Overview

This microservice validates whether user-uploaded photos meet the minimum quality and content requirements for profile onboarding.

### What It Does ✅
- Validates photo usability for onboarding
- Detects presence of human faces
- Assesses image quality (blur, lighting, resolution)
- Identifies non-photographic content (screenshots, memes, AI-generated)

### What It Does NOT Do ❌
- Identity verification
- Face recognition
- Attractiveness scoring
- Age, gender, or emotion inference
- Biometric analysis

---

## Quick Start

```bash
# 1. Install dependencies
npm install

# 2. Copy environment configuration
cp .env.example .env

# 3. Run in test mode (uses mock AI)
npm run dev

# The server starts on http://localhost:3000
```

---

## API Endpoints

### POST `/onboarding/photos`

Upload 3-6 photos for validation.

**Request:**
```bash
curl -X POST http://localhost:3000/onboarding/photos \
  -F "photos=@photo1.jpg" \
  -F "photos=@photo2.jpg" \
  -F "photos=@photo3.jpg"
```

**Response:**
```json
{
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "photos": [
    {
      "photo_id": "...",
      "usable": true,
      "has_face": true,
      "issues": []
    }
  ],
  "aggregate": {
    "usable_photos": 3,
    "face_photos": 3,
    "status": "pass"
  }
}
```

### GET `/onboarding/photos/:user_id`

Get photo metadata for a user (no image bytes).

**Request:**
```bash
curl http://localhost:3000/onboarding/photos/550e8400-e29b-41d4-a716-446655440000
```

**Response:**
```json
{
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "photos": [
    {
      "photo_id": "...",
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

## Validation Pipeline

```
┌─────────────────────────────────────────────────────────────────┐
│                     Photo Upload (3-6 images)                    │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                  CV Validation (Deterministic)                   │
│  ┌───────────┐  ┌───────────┐  ┌───────────────┐                │
│  │   Blur    │  │ Brightness│  │  Resolution   │                │
│  │ Detection │  │   Check   │  │    Check      │                │
│  └───────────┘  └───────────┘  └───────────────┘                │
└─────────────────────────────────────────────────────────────────┘
                              │
              ┌───────────────┴───────────────┐
              │                               │
         CV PASSED                       CV FAILED
              │                               │
              ▼                               ▼
┌─────────────────────────┐     ┌─────────────────────────┐
│    AI Vision Analysis   │     │   Immediate Rejection   │
│  ┌───────────────────┐  │     │  (Skip AI validation)   │
│  │   Face Presence   │  │     └─────────────────────────┘
│  │   Face Visibility │  │
│  │   Photo Type      │  │
│  └───────────────────┘  │
└─────────────────────────┘
              │
              ▼
┌─────────────────────────────────────────────────────────────────┐
│                       Aggregation Rules                          │
│  • Minimum 3 usable photos required                              │
│  • At least 1 usable photo with face required                    │
└─────────────────────────────────────────────────────────────────┘
              │
              ▼
        ┌─────────┐
        │  PASS   │ or  │ NEEDS MORE PHOTOS │
        └─────────┘
```

---

## Issue Tags

| Tag | Description |
|-----|-------------|
| `no_face` | No human face detected |
| `blurred` | Image is too blurry |
| `too_dark` | Image is too dark |
| `too_bright` | Image is overexposed |
| `obstructed_face` | Face blocked by mask, helmet, etc. |
| `non_photographic` | Screenshot, meme, drawing, or AI-generated |
| `low_resolution` | Resolution below 640×480 |

---

## Configuration

Configure via `.env` file:

```env
# Server
PORT=3000

# Storage
STORAGE_PATH=./storage/photos

# OpenAI (required for production)
OPENAI_API_KEY=sk-...

# Test Mode
USE_MOCK_AI=true

# CV Thresholds (optional)
BLUR_THRESHOLD=100
BRIGHTNESS_MIN=30
BRIGHTNESS_MAX=225
MIN_WIDTH=640
MIN_HEIGHT=480
```

---

## Test Mode

The service includes a mock AI client for testing:

```bash
# Start in test mode
USE_MOCK_AI=true npm run dev
```

The mock client uses filename patterns to determine responses:

| Filename Pattern | Result |
|-----------------|--------|
| `*_pass.*` | Usable |
| `*no_face*` | Issue: no_face |
| `*blur*` | Issue: blurred |
| `*dark*` | Issue: too_dark |

See `test-images/README.md` for full pattern list.

---

## Project Structure

```
src/
├── index.ts              # Entry point
├── app.ts                # Express app setup
├── config/
│   └── env.ts            # Environment configuration
├── schemas/
│   ├── photo.schema.ts   # Upload/photo schemas
│   └── validation.schema.ts
├── routes/
│   └── onboarding.routes.ts
├── controllers/
│   └── photo.controller.ts
├── services/
│   ├── storage.service.ts   # Photo storage
│   ├── cv.service.ts        # CV validation
│   ├── ai.service.ts        # AI orchestration
│   └── aggregation.service.ts
├── clients/
│   ├── ai.client.interface.ts
│   ├── openai.client.ts     # OpenAI implementation
│   └── mock.client.ts       # Mock for testing
├── middleware/
│   ├── upload.middleware.ts
│   └── error.middleware.ts
├── types/
│   └── index.ts
└── utils/
    └── helpers.ts
```

---

## Storage Rules (Critical)

**STORED:**
- ✅ Raw image files
- ✅ Photo metadata (id, user_id, filename, created_at)
- ✅ Final usability flags (`usable`, `has_face`)

**NOT STORED:**
- ❌ Face embeddings
- ❌ Bounding boxes
- ❌ AI confidence scores
- ❌ Any biometric data

---

## Scripts

```bash
# Development (with hot reload)
npm run dev

# Build for production
npm run build

# Run production build
npm start

# Test instructions
npm test
```

---

## Documentation

- [Service Specification](./docs/SERVICE_SPEC.md)
- [API Documentation](./docs/API.md)
- [Integration Guide](./docs/INTEGRATION.md)

---

## Error Handling

All errors return user-safe messages:

```json
{
  "error": "insufficient_photos",
  "message": "Please upload at least 3 photos. You uploaded 2."
}
```

Common error codes:
- `no_files` - No photos uploaded
- `insufficient_photos` - Less than 3 photos
- `too_many_files` - More than 6 photos
- `file_too_large` - File exceeds 10MB
- `invalid_file_type` - Unsupported format

---

## License

MIT © Unora
