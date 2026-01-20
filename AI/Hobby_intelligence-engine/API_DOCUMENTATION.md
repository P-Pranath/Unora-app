# Hobby Intelligence Engine — API Documentation

> **For Unora App Developers**  
> Version: 1.0  
> Base URL: `http://localhost:3000` (development)

---

## Overview

The Hobby Intelligence Engine is a **stateless AI microservice** that powers Unora's hobby-based check-in system. Each API call is independent—**the service has no memory of previous requests**.

The service supports **three primary flows**, each mapped to specific endpoints:

| Flow | Purpose | Endpoint |
|------|---------|----------|
| **Flow A** | Hobby Legitimacy Validation | `POST /validate` |
| **Flow B** | Daily Check-In Question Generation | `POST /question` |
| **Flow C** | Hobby Echo Generation (Post-Answer) | `POST /echo` |

Additionally, a **Pipeline endpoint** combines Flows A + B into a single request.

---

## Quick Start

### Prerequisites
- Service must be running at the configured host/port
- All requests use `Content-Type: application/json`

### Health Check
```bash
GET /health
```
```json
{ "status": "ok", "service": "hobby-intelligence-engine" }
```

---

## Flow A — Hobby Legitimacy Validation

**Purpose**: Validate if a user-submitted hobby is sane, legal, and non-harmful before allowing it into the system.

### Endpoint
```
POST /validate
```

### Request Body
```json
{
  "hobby": "string (required, 1-200 characters)"
}
```

### Response — Success (200)
```json
{
  "legit": true | false,
  "reason": "string | null",
  "risk_level": "low" | "medium" | "high"
}
```

### Field Definitions

| Field | Type | Description |
|-------|------|-------------|
| `legit` | boolean | `true` if hobby is acceptable, `false` if rejected |
| `reason` | string \| null | Explanation if rejected, `null` if accepted |
| `risk_level` | enum | Safety classification: `low`, `medium`, or `high` |

### Examples

**Valid Hobby:**
```bash
curl -X POST http://localhost:3000/validate \
  -H "Content-Type: application/json" \
  -d '{"hobby": "painting"}'
```
```json
{
  "legit": true,
  "reason": null,
  "risk_level": "low"
}
```

**Invalid Hobby:**
```bash
curl -X POST http://localhost:3000/validate \
  -H "Content-Type: application/json" \
  -d '{"hobby": "stalking"}'
```
```json
{
  "legit": false,
  "reason": "Surveillance-based activity that invades others' privacy",
  "risk_level": "high"
}
```

### Integration Notes

> [!IMPORTANT]
> - This is NOT about consistency or commitment—only legitimacy and safety
> - Always validate hobbies before storing them in user profiles
> - If `legit` is `false`, do NOT proceed to Flow B or C

---

## Flow B — Daily Check-In Question Generation

**Purpose**: Generate a daily check-in question with one-tap answer options for a validated hobby.

### Endpoint
```
POST /question
```

### Request Body
```json
{
  "hobby": "string (required)",
  "category": "string (required, see valid values)",
  "previous_questions": ["string"] (optional, default: [])
}
```

### Valid Category Values
```
"Physical" | "Creative" | "Skill / Practice" | "Learning" | 
"Mindfulness" | "Social" | "Collection" | "Maintenance"
```

> [!TIP]
> Use the `/classify` endpoint to get the correct category for a hobby if you don't have it cached.

### Response — Success (200)
```json
{
  "question": "string",
  "options": ["string", "string", "string", "string"]
}
```

### Field Definitions

| Field | Type | Description |
|-------|------|-------------|
| `question` | string | The check-in question to display to the user |
| `options` | string[] | 2-4 tap-friendly answer options |

### Examples

**Basic Request:**
```bash
curl -X POST http://localhost:3000/question \
  -H "Content-Type: application/json" \
  -d '{
    "hobby": "painting",
    "category": "Creative"
  }'
```
```json
{
  "question": "What did you work on?",
  "options": ["Landscape", "Portrait", "Abstract", "Other"]
}
```

**With Previous Questions (to avoid repetition):**
```bash
curl -X POST http://localhost:3000/question \
  -H "Content-Type: application/json" \
  -d '{
    "hobby": "painting",
    "category": "Creative",
    "previous_questions": ["What did you work on?", "How long did you paint?"]
  }'
```

### Integration Notes

> [!WARNING]
> - Do NOT call this endpoint before validating the hobby (Flow A)
> - Questions are neutral and low-effort—no performance framing
> - Pass `previous_questions` to avoid repetitive daily questions

---

## Flow C — Hobby Echo Generation (Post-Answer)

**Purpose**: Generate a privacy-safe, third-person statement about the user's hobby check-in, suitable for display to their connection.

### Endpoint
```
POST /echo
```

### Request Body
```json
{
  "hobby": "string (required)",
  "selected_option": "string (required)"
}
```

### Response — Success (200)
```json
{
  "echo": "string (under 20 words)"
}
```

### Examples

```bash
curl -X POST http://localhost:3000/echo \
  -H "Content-Type: application/json" \
  -d '{
    "hobby": "painting",
    "selected_option": "Landscape"
  }'
```
```json
{
  "echo": "Your partner checked in with painting and focused on landscapes."
}
```

```bash
curl -X POST http://localhost:3000/echo \
  -H "Content-Type: application/json" \
  -d '{
    "hobby": "running",
    "selected_option": "Morning jog"
  }'
```
```json
{
  "echo": "Your partner went for a morning run today."
}
```

### Echo Constraints

> [!IMPORTANT]
> The echo output ALWAYS follows these rules:
> - **Third-person only** — Never uses "I" or "you"
> - **Privacy-safe** — No location, time, or personally identifiable information
> - **No advice** — Purely observational, not motivational or prescriptive
> - **No identity inference** — Never implies user traits or characteristics
> - **No effort inference** — Never suggests how hard they worked
> - **Under 20 words** — Concise and scannable

---

## Combined Pipeline Endpoint

**Purpose**: Run Flow A (Validation) + automatic classification + Flow B (Question) in a single request. Useful for onboarding new hobbies.

### Endpoint
```
POST /pipeline
```

### Request Body
```json
{
  "hobby": "string (required, 1-200 characters)",
  "previous_questions": ["string"] (optional, default: [])
}
```

### Response — Success (200)

**If hobby is valid:**
```json
{
  "validation": {
    "legit": true,
    "reason": null,
    "risk_level": "low"
  },
  "classification": {
    "category": "Creative",
    "effort_type": "solo",
    "cadence": "flexible"
  },
  "question": {
    "question": "What was your subject?",
    "options": ["Architecture", "Nature", "People", "Other"]
  }
}
```

**If hobby is invalid (pipeline stops early):**
```json
{
  "validation": {
    "legit": false,
    "reason": "Illegal activity that causes harm",
    "risk_level": "high"
  }
}
```

> [!NOTE]
> When `validation.legit` is `false`, the `classification` and `question` fields will be absent.

### Classification Fields

| Field | Type | Values |
|-------|------|--------|
| `category` | enum | `Physical`, `Creative`, `Skill / Practice`, `Learning`, `Mindfulness`, `Social`, `Collection`, `Maintenance` |
| `effort_type` | enum | `solo`, `social`, `mixed` |
| `cadence` | enum | `daily`, `irregular`, `flexible` |

---

## Error Handling

### Input Validation Errors (400)
```json
{
  "error": "Invalid input",
  "details": {
    "fieldErrors": {
      "hobby": ["Hobby is required"]
    }
  }
}
```

### Server Errors (500)
```json
{
  "error": "Validation failed"
}
```

### Insufficient Input (Flow Detection Failure)

If your client code cannot determine which flow to call, you should respond with:
```json
{
  "error": "insufficient_or_invalid_input"
}
```

---

## Integration Patterns

### Pattern 1: New Hobby Onboarding
```
User submits hobby → POST /pipeline → Store hobby + classification → Display question
```

### Pattern 2: Daily Check-In
```
Load user's hobby + cached category → POST /question → Display question → 
User selects option → POST /echo → Display echo to partner
```

### Pattern 3: Hobby Validation Only
```
User types hobby → POST /validate → Show error if invalid, or allow save if valid
```

---

## TypeScript Types (Reference)

```typescript
// Flow A — Validation
interface ValidateRequest {
  hobby: string; // 1-200 chars
}

interface ValidateResponse {
  legit: boolean;
  reason: string | null;
  risk_level: 'low' | 'medium' | 'high';
}

// Flow B — Question
interface QuestionRequest {
  hobby: string;
  category: 'Physical' | 'Creative' | 'Skill / Practice' | 'Learning' | 
            'Mindfulness' | 'Social' | 'Collection' | 'Maintenance';
  previous_questions?: string[];
}

interface QuestionResponse {
  question: string;
  options: string[]; // 2-4 items
}

// Flow C — Echo
interface EchoRequest {
  hobby: string;
  selected_option: string;
}

interface EchoResponse {
  echo: string; // Under 20 words
}

// Pipeline
interface PipelineRequest {
  hobby: string; // 1-200 chars
  previous_questions?: string[];
}

interface PipelineResponse {
  validation: ValidateResponse;
  classification?: {
    category: string;
    effort_type: 'solo' | 'social' | 'mixed';
    cadence: 'daily' | 'irregular' | 'flexible';
  };
  question?: QuestionResponse;
}
```

---

## Service Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | 3000 | Server port |
| `HOST` | 0.0.0.0 | Server host |
| `AI_PROVIDER` | huggingface | AI backend (huggingface, openai, gemini) |

---

## Rate Limits & Performance

- **Stateless**: Each request is independent; no session required
- **Latency**: Expect 200-800ms per call (AI processing)
- **No caching**: Fresh AI responses on every call
- **Concurrent safe**: Service handles parallel requests

---

## Support

For issues, check:
1. Health endpoint returns `{ "status": "ok" }`
2. Request body is valid JSON with correct fields
3. AI provider credentials are configured (see `.env`)

---

*Last updated: January 2026*
