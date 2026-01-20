# Vision Validation API Documentation

## Endpoint

### POST `/api/v1/validate-photo`

Validates whether an uploaded photo is usable for onboarding verification.

---

## Request

### Headers

| Header | Value | Required |
|--------|-------|----------|
| `Content-Type` | `multipart/form-data` or `application/json` | Yes |
| `Authorization` | `Bearer <token>` | Yes |

### Body (multipart/form-data)

| Field | Type | Description |
|-------|------|-------------|
| `image` | File | The photo file to validate |

### Body (application/json)

| Field | Type | Description |
|-------|------|-------------|
| `image` | string | Base64-encoded image or URL |

---

## Response

### Success Response (200 OK)

```json
{
  "usable": true,
  "issues": [],
  "confidence": "high"
}
```

### Validation Failed Response (200 OK)

```json
{
  "usable": false,
  "issues": ["no_face", "blurred"],
  "confidence": "high"
}
```

### Error Response (400 Bad Request)

```json
{
  "error": "invalid_image",
  "message": "The provided image could not be processed"
}
```

### Error Response (401 Unauthorized)

```json
{
  "error": "unauthorized",
  "message": "Invalid or missing authentication token"
}
```

---

## Response Fields

### `usable` (boolean)

Indicates whether the photo meets minimum requirements for onboarding.

| Value | Meaning |
|-------|---------|
| `true` | Photo is acceptable for use |
| `false` | Photo should be rejected |

### `issues` (string[])

Array of issue tags identifying problems with the photo. Empty array if photo is usable.

**Valid Issue Tags:**
- `no_face`
- `blurred`
- `too_dark`
- `too_bright`
- `obstructed_face`
- `non_photographic`
- `low_resolution`

### `confidence` (string)

Confidence level of the validation decision.

| Value | Meaning |
|-------|---------|
| `high` | Clear, unambiguous decision |
| `medium` | Borderline but acceptable |
| `low` | Ambiguous, conservative decision |

---

## Example Usage

### cURL (File Upload)

```bash
curl -X POST https://api.unora.app/v1/validate-photo \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -F "image=@/path/to/photo.jpg"
```

### cURL (Base64)

```bash
curl -X POST https://api.unora.app/v1/validate-photo \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"image": "data:image/jpeg;base64,/9j/4AAQ..."}'
```

### JavaScript (Fetch)

```javascript
const formData = new FormData();
formData.append('image', fileInput.files[0]);

const response = await fetch('/api/v1/validate-photo', {
  method: 'POST',
  headers: {
    'Authorization': `Bearer ${token}`
  },
  body: formData
});

const result = await response.json();

if (result.usable) {
  console.log('Photo accepted');
} else {
  console.log('Photo rejected:', result.issues);
}
```

---

## Rate Limits

| Tier | Requests/Minute |
|------|-----------------|
| Free | 10 |
| Plus | 30 |
| Pro | 100 |

---

## Supported Formats

| Format | Max Size |
|--------|----------|
| JPEG | 10 MB |
| PNG | 10 MB |
| WebP | 10 MB |
| HEIC | 10 MB |
