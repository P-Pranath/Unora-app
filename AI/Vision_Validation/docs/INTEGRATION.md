# Vision Validation Integration Guide

## Overview

This guide explains how to integrate the Vision Validation service into your Unora application workflow.

---

## Integration Flow

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│   User Upload   │ ──▶ │ Vision Service  │ ──▶ │  App Response   │
│    (Photo)      │     │  (Validation)   │     │ (Accept/Reject) │
└─────────────────┘     └─────────────────┘     └─────────────────┘
```

---

## Onboarding Integration

### Step 1: Photo Upload Screen

When user uploads a photo during onboarding:

```typescript
async function handlePhotoUpload(file: File): Promise<ValidationResult> {
  const formData = new FormData();
  formData.append('image', file);

  const response = await fetch('/api/v1/validate-photo', {
    method: 'POST',
    headers: {
      'Authorization': `Bearer ${authToken}`
    },
    body: formData
  });

  return response.json();
}
```

### Step 2: Handle Validation Result

```typescript
interface ValidationResult {
  usable: boolean;
  issues: string[];
  confidence: 'high' | 'medium' | 'low';
}

function handleValidationResult(result: ValidationResult) {
  if (result.usable) {
    // Proceed with onboarding
    proceedToNextStep();
  } else {
    // Show user-friendly error message
    showPhotoRetryUI(result.issues);
  }
}
```

### Step 3: User-Friendly Error Messages

Map issue tags to user-facing messages:

```typescript
const issueMessages: Record<string, string> = {
  'no_face': 'We couldn\'t detect a face in your photo. Please upload a clear photo of yourself.',
  'blurred': 'Your photo appears blurry. Please upload a clearer image.',
  'too_dark': 'Your photo is too dark. Please use better lighting.',
  'too_bright': 'Your photo is overexposed. Please reduce the brightness.',
  'obstructed_face': 'Your face appears to be covered. Please remove any obstructions.',
  'non_photographic': 'Please upload an actual photo, not a screenshot or drawing.',
  'low_resolution': 'Your photo resolution is too low. Please upload a higher quality image.'
};

function getErrorMessage(issues: string[]): string {
  return issues.map(issue => issueMessages[issue] || 'Please upload a different photo.').join(' ');
}
```

---

## UI Recommendations

### Photo Upload State Machine

```
┌──────────┐
│  IDLE    │
└────┬─────┘
     │ User selects photo
     ▼
┌──────────┐
│ UPLOADING│
└────┬─────┘
     │ Validation complete
     ▼
┌──────────────────────────────────┐
│ VALIDATED                         │
│ ┌──────────┐    ┌──────────────┐ │
│ │ ACCEPTED │ OR │   REJECTED   │ │
│ └──────────┘    └──────────────┘ │
└──────────────────────────────────┘
```

### Rejected Photo UI

When a photo is rejected:

1. **Show the rejected photo** with a subtle overlay
2. **Display clear error message** based on issue tags
3. **Provide "Try Again" CTA** to upload a new photo
4. **Do NOT block the user** - allow them to retry immediately

### Loading State

While validation is in progress:

1. Show a subtle loading indicator
2. Display message: "Checking your photo..."
3. Disable the continue button until validation completes

---

## Error Handling

### Network Errors

```typescript
async function validateWithRetry(file: File, maxRetries = 3): Promise<ValidationResult> {
  for (let attempt = 1; attempt <= maxRetries; attempt++) {
    try {
      return await handlePhotoUpload(file);
    } catch (error) {
      if (attempt === maxRetries) {
        throw new Error('Unable to validate photo. Please try again later.');
      }
      await delay(1000 * attempt); // Exponential backoff
    }
  }
}
```

### Timeout Handling

```typescript
const VALIDATION_TIMEOUT = 30000; // 30 seconds

async function validateWithTimeout(file: File): Promise<ValidationResult> {
  const controller = new AbortController();
  const timeoutId = setTimeout(() => controller.abort(), VALIDATION_TIMEOUT);

  try {
    const response = await fetch('/api/v1/validate-photo', {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${authToken}`
      },
      body: formData,
      signal: controller.signal
    });
    return response.json();
  } finally {
    clearTimeout(timeoutId);
  }
}
```

---

## Testing

### Test Cases

| Scenario | Expected Result |
|----------|-----------------|
| Clear face photo | `usable: true` |
| Blurry photo | `usable: false`, issues: `["blurred"]` |
| Dark photo | `usable: false`, issues: `["too_dark"]` |
| Masked face | `usable: false`, issues: `["obstructed_face"]` |
| Screenshot | `usable: false`, issues: `["non_photographic"]` |
| No face | `usable: false`, issues: `["no_face"]` |

### Mock Service for Development

```typescript
function mockValidatePhoto(file: File): Promise<ValidationResult> {
  return new Promise(resolve => {
    setTimeout(() => {
      // Simulate validation based on file size (for testing)
      if (file.size < 10000) {
        resolve({
          usable: false,
          issues: ['low_resolution'],
          confidence: 'high'
        });
      } else {
        resolve({
          usable: true,
          issues: [],
          confidence: 'high'
        });
      }
    }, 500);
  });
}
```

---

## Security Considerations

1. **Never log image data** - Images may contain sensitive information
2. **Use HTTPS only** - All API calls must be encrypted
3. **Validate file types** - Only accept expected image formats
4. **Size limits** - Enforce maximum file size (10MB)
5. **Rate limiting** - Protect against abuse
