# Personality Intelligence Service

A standalone microservice for Unora that infers probabilistic personality dimensions through situational tap-based questions, updates confidence over time, and generates one human-readable personality summary via LLM.

## Overview

This service:
- Asks situational, tap-based questions (not self-descriptive)
- Infers 7 probabilistic personality dimensions
- Updates confidence over time with diminishing returns
- Generates one calm, human-readable paragraph summary via LLM
- Runs independently as a microservice

**What this is NOT:**
- A personality test
- A Big Five implementation
- A therapy or diagnosis system
- A visible scoring or ranking engine

## Quick Start

### Prerequisites
- Node.js 18+ 
- npm or yarn
- OpenAI API key (or compatible endpoint)

### Installation

```bash
# Clone and enter the directory
cd personality-service

# Install dependencies
npm install

# Copy environment file and configure
cp .env.example .env
# Edit .env with your OpenAI API key

# Setup database (generate client, create tables, seed questions)
npm run setup

# Start the development server
npm run dev
```

The server will start at `http://localhost:3000`.

### Environment Variables

```env
# Database (SQLite local file)
DATABASE_URL="file:./dev.db"

# Server port
PORT=3000

# OpenAI API key (required for summary generation)
OPENAI_API_KEY=sk-your-api-key-here

# Optional: OpenAI-compatible endpoint
# OPENAI_BASE_URL=https://api.openai.com/v1

# LLM model to use
OPENAI_MODEL=gpt-4o-mini
```

## API Endpoints

### POST /profile/init
Create a new user personality profile.

**Request:**
```json
{
  "userId": "user-123"
}
```

**Response (201):**
```json
{
  "userId": "user-123",
  "personalityState": {
    "emotional_regulation": { "score": 0.5, "confidence": 0.1 },
    "communication_style": { "score": 0.5, "confidence": 0.1 },
    "emotional_availability": { "score": 0.5, "confidence": 0.1 },
    "consistency_style": { "score": 0.5, "confidence": 0.1 },
    "conflict_posture": { "score": 0.5, "confidence": 0.1 },
    "energy_orientation": { "score": 0.5, "confidence": 0.1 },
    "decision_pace": { "score": 0.5, "confidence": 0.1 }
  },
  "message": "Profile created successfully"
}
```

### GET /question/next
Get the next adaptive question for a user.

**Query Parameters:**
- `userId` (required): User's ID
- `mode` (optional): `ONBOARDING` (default, 8 questions) or `STREAK_CHECKIN` (1 question)

**Response:**
```json
{
  "question": {
    "id": "Q_ER_01",
    "scenario": "A plan you were excited about gets cancelled last minute.",
    "options": [
      { "label": "I'm annoyed briefly, then move on to something else" },
      { "label": "It throws me off for a while, but I adjust" },
      { "label": "I feel quite disappointed and need some time to reset" }
    ]
  },
  "questionsAnswered": 0,
  "mode": "ONBOARDING"
}
```

**When complete:**
```json
{
  "complete": true,
  "message": "Reached maximum questions for ONBOARDING mode (8)",
  "questionsAnswered": 8
}
```

### POST /answer
Submit an answer to a question.

**Request:**
```json
{
  "userId": "user-123",
  "questionId": "Q_ER_01",
  "selectedOption": 0
}
```

**To skip:**
```json
{
  "userId": "user-123",
  "questionId": "Q_ER_01",
  "skip": true
}
```

**Response:**
```json
{
  "updated": true,
  "skipped": false,
  "questionsAnswered": 1,
  "nextQuestion": {
    "id": "Q_CS_02",
    "scenario": "Someone keeps interrupting you during a conversation.",
    "options": [...]
  }
}
```

### GET /summary
Generate a fresh personality summary.

**Query Parameters:**
- `userId` (required): User's ID

**Response:**
```json
{
  "summary": "Tends to stay grounded during tense moments and prefers to process things internally before responding. Shows up consistently rather than intensely, and values thoughtful communication once comfortable. Generally adapts well to changes while maintaining a preference for some structure in daily routines.",
  "generatedAt": "2024-01-15T10:30:00.000Z",
  "confidence": {
    "overall": 0.65,
    "dimensions": {
      "emotional_regulation": 0.7,
      "communication_style": 0.6,
      "emotional_availability": 0.55,
      "consistency_style": 0.65,
      "conflict_posture": 0.6,
      "energy_orientation": 0.75,
      "decision_pace": 0.7
    }
  }
}
```

## Personality Model

### Dimensions

| Dimension | Low Score | High Score |
|-----------|-----------|------------|
| `emotional_regulation` | Experiences emotions deeply, needs processing time | Stays grounded, processes emotions smoothly |
| `communication_style` | Indirect, reads between the lines | Direct, says what they mean |
| `emotional_availability` | Takes time to open up, values privacy | Emotionally open and accessible |
| `consistency_style` | Enjoys spontaneity, adapts to change | Values routine and predictability |
| `conflict_posture` | Avoids conflict, seeks harmony | Addresses disagreements directly |
| `energy_orientation` | Recharges through solitude | Gains energy from social interaction |
| `decision_pace` | Deliberate, needs time to decide | Decides quickly, trusts instincts |

### Scoring Algorithm

When an answer is submitted:
1. **Score Update**: Weighted average moving toward the option's impact
   ```
   newScore = (oldScore × confidence + targetScore × weight) / (confidence + weight)
   ```
2. **Confidence Update**: Incremental with diminishing returns
   - Normal: `+0.08` per answer
   - After 0.8 confidence: `+0.02` per answer (diminishing returns)

3. **Skip Handling**: No score or confidence change

## Question Design

Each question:
- Is **situational**, not self-descriptive
- Targets 1-3 dimensions
- Has multiple options with weighted impacts (-0.3 to +0.3)

Example:
```json
{
  "id": "Q_ER_01",
  "scenario": "A plan you were excited about gets cancelled last minute.",
  "dimensionTargets": ["emotional_regulation"],
  "options": [
    {
      "label": "I'm annoyed briefly, then move on to something else",
      "impacts": { "emotional_regulation": 0.25 }
    },
    {
      "label": "It throws me off for a while, but I adjust",
      "impacts": { "emotional_regulation": 0.0 }
    },
    {
      "label": "I feel quite disappointed and need some time to reset",
      "impacts": { "emotional_regulation": -0.2 }
    }
  ]
}
```

## Adaptive Question Selection

The engine prioritizes:
1. **Low-confidence dimensions** (< 0.6)
2. **Avoids repetition** (never same dimension twice in a row)
3. **Multi-dimensional questions** (slight bonus)

### Modes
- **ONBOARDING**: 8 questions for initial profile building
- **STREAK_CHECKIN**: 1 question for daily engagement

## Sample Curl Commands

```bash
# 1. Create a user profile
curl -X POST http://localhost:3000/profile/init \
  -H "Content-Type: application/json" \
  -d '{"userId": "test-user-001"}'

# 2. Get the first question
curl "http://localhost:3000/question/next?userId=test-user-001&mode=ONBOARDING"

# 3. Submit an answer (select first option)
curl -X POST http://localhost:3000/answer \
  -H "Content-Type: application/json" \
  -d '{"userId": "test-user-001", "questionId": "Q_ER_01", "selectedOption": 0}'

# 4. Skip a question
curl -X POST http://localhost:3000/answer \
  -H "Content-Type: application/json" \
  -d '{"userId": "test-user-001", "questionId": "Q_CS_02", "skip": true}'

# 5. Get personality summary
curl "http://localhost:3000/summary?userId=test-user-001"

# 6. Check service health
curl http://localhost:3000/health
```

## Database Management

### Reset Database
```bash
npm run db:reset
```
This will:
1. Delete the existing database
2. Recreate all tables
3. Reseed all questions

### View Database (Optional)
```bash
npx prisma studio
```

## Adding New Questions

1. Edit `src/questions/questionBank.ts`
2. Add new questions following the structure:
```typescript
{
  id: 'Q_XX_NN',  // Dimension prefix + number
  scenario: 'The situational scenario text...',
  dimensionTargets: ['primary_dimension', 'secondary_dimension'],
  options: [
    { label: 'Option text', impacts: { primary_dimension: 0.2 } },
    { label: 'Option text', impacts: { primary_dimension: 0.0 } },
    { label: 'Option text', impacts: { primary_dimension: -0.2 } }
  ]
}
```
3. Run seed to update database:
```bash
npm run db:seed
```

## Project Structure

```
personality-service/
├── prisma/
│   ├── schema.prisma      # Database schema
│   └── seed.ts            # Question seeding script
├── src/
│   ├── db/
│   │   └── prisma.ts      # Prisma client singleton
│   ├── engine/
│   │   └── questionSelector.ts  # Adaptive selection logic
│   ├── questions/
│   │   └── questionBank.ts      # 46 situational questions
│   ├── routes/
│   │   ├── answerRoutes.ts      # POST /answer
│   │   ├── profileRoutes.ts     # POST /profile/init
│   │   ├── questionRoutes.ts    # GET /question/next
│   │   └── summaryRoutes.ts     # GET /summary
│   ├── scoring/
│   │   └── scorer.ts            # Score & confidence updates
│   ├── summary/
│   │   └── summaryGenerator.ts  # LLM summary generation
│   ├── types/
│   │   └── index.ts             # TypeScript interfaces
│   └── index.ts                 # Express app entry point
├── .env.example
├── package.json
├── tsconfig.json
└── README.md
```

## Design Principles

- **No labels, only dimensions**: Scores are spectrums, not categories
- **Probabilistic, not absolute**: Everything has confidence levels
- **Confidence matters**: Low confidence = more weight to new answers
- **Skip without penalty**: Users can skip freely
- **Situational questions**: Never ask "Are you an introvert?"
- **One paragraph only**: Summary is human, not clinical
- **Raw data hidden**: Impacts and scores are never exposed to users

## License

MIT - Unora © 2024
