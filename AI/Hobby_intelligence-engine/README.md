# Hobby Intelligence Engine

A production-grade AI-powered microservice for validating hobbies, classifying them into a fixed ontology, generating daily check-in questions, and producing privacy-safe progress updates.

## Architecture

```
hobby-intelligence-engine/
├── src/
│   ├── index.ts                    # Server entry point
│   ├── config/env.ts               # Environment configuration
│   ├── ai/
│   │   ├── types.ts                # AI client interfaces
│   │   ├── client.ts               # AI client factory
│   │   └── providers/openai.ts     # OpenAI implementation
│   ├── modules/
│   │   ├── hobbyValidator.ts       # Validates hobby legitimacy
│   │   ├── hobbyClassifier.ts      # Classifies into categories
│   │   ├── questionGenerator.ts    # Generates check-in questions
│   │   └── hobbyEchoGenerator.ts   # Generates privacy-safe updates
│   ├── routes/hobby.routes.ts      # API endpoints
│   └── types/schemas.ts            # Zod validation schemas
└── test/simulation.ts              # Test runner
```

## Setup

```bash
# Install dependencies
npm install

# Copy environment file and add your OpenAI API key
cp .env.example .env
# Edit .env and set OPENAI_API_KEY

# Start development server
npm run dev

# Run tests
npm test
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/validate` | Validate a hobby for legitimacy |
| POST | `/classify` | Classify into category/effort/cadence |
| POST | `/question` | Generate check-in question |
| POST | `/echo` | Generate privacy-safe update |
| POST | `/pipeline` | Run full validation→classification→question |
| GET | `/health` | Health check |

## Example Requests

### Validate a Hobby
```bash
curl -X POST http://localhost:3000/validate \
  -H "Content-Type: application/json" \
  -d '{"hobby": "painting"}'
```
```json
{"legit": true, "reason": null, "risk_level": "low"}
```

### Classify a Hobby
```bash
curl -X POST http://localhost:3000/classify \
  -H "Content-Type: application/json" \
  -d '{"hobby": "painting"}'
```
```json
{"category": "Creative", "effort_type": "solo", "cadence": "flexible"}
```

### Generate a Question
```bash
curl -X POST http://localhost:3000/question \
  -H "Content-Type: application/json" \
  -d '{"hobby": "painting", "category": "Creative", "previous_questions": []}'
```
```json
{"question": "What did you work on?", "options": ["Landscape", "Portrait", "Abstract", "Other"]}
```

### Generate an Echo
```bash
curl -X POST http://localhost:3000/echo \
  -H "Content-Type: application/json" \
  -d '{"hobby": "painting", "selected_option": "Landscape"}'
```
```json
{"echo": "Your partner checked in with painting and focused on landscapes."}
```

### Run Full Pipeline
```bash
curl -X POST http://localhost:3000/pipeline \
  -H "Content-Type: application/json" \
  -d '{"hobby": "urban sketching", "previous_questions": []}'
```
```json
{
  "validation": {"legit": true, "reason": null, "risk_level": "low"},
  "classification": {"category": "Creative", "effort_type": "solo", "cadence": "flexible"},
  "question": {"question": "What was your subject?", "options": ["Architecture", "Nature", "People", "Other"]}
}
```

### Invalid Hobby (Rejected)
```bash
curl -X POST http://localhost:3000/pipeline \
  -H "Content-Type: application/json" \
  -d '{"hobby": "stalking"}'
```
```json
{
  "validation": {"legit": false, "reason": "Surveillance-based activity that invades others' privacy", "risk_level": "high"}
}
```

## Design Decisions

1. **AI Abstraction**: All AI calls go through `AIClient` interface in `src/ai/`. Swap providers by implementing new provider in `providers/` folder.

2. **Schema Validation**: Zod schemas in `src/types/schemas.ts` validate all inputs/outputs at runtime.

3. **Deterministic Outputs**: Low temperature (0.1-0.3) and JSON mode ensure consistent, repeatable responses.

4. **Safety First**: HobbyValidator acts as a gate - invalid hobbies are rejected before any further processing.

5. **Privacy by Design**: HobbyEcho uses strict third-person format with no inferred details.

## Swapping AI Providers

To add a new provider (e.g., local LLM):

1. Create `src/ai/providers/local.ts` implementing `AIClient` interface
2. Add case to `src/ai/client.ts` factory function
3. Update environment variables as needed

## License

MIT
