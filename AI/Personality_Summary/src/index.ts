// ============================================================================
// Express Application Entry Point - Personality Intelligence Service
// ============================================================================
// Main entry point for the personality service microservice.
// Configures Express, middleware, and routes.
// ============================================================================

import express, { Express, Request, Response, NextFunction } from 'express';
import profileRoutes from './routes/profileRoutes';
import questionRoutes from './routes/questionRoutes';
import answerRoutes from './routes/answerRoutes';
import summaryRoutes from './routes/summaryRoutes';

// ============================================================================
// Load Environment Variables
// ============================================================================

// dotenv is optional - allows overriding via .env file
try {
    require('dotenv').config();
} catch {
    // dotenv not installed, use environment variables directly
}

// ============================================================================
// App Configuration
// ============================================================================

const app: Express = express();
const PORT = process.env.PORT || 3000;

// ============================================================================
// Middleware
// ============================================================================

// Parse JSON request bodies
app.use(express.json());

// Request logging (development)
app.use((req: Request, res: Response, next: NextFunction) => {
    const timestamp = new Date().toISOString();
    console.log(`[${timestamp}] ${req.method} ${req.path}`);
    next();
});

// CORS headers (for development)
app.use((req: Request, res: Response, next: NextFunction) => {
    res.header('Access-Control-Allow-Origin', '*');
    res.header('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE, OPTIONS');
    res.header('Access-Control-Allow-Headers', 'Content-Type, Authorization');

    if (req.method === 'OPTIONS') {
        return res.sendStatus(200);
    }

    next();
});

// ============================================================================
// Routes
// ============================================================================

// Health check endpoint
app.get('/health', (req: Request, res: Response) => {
    res.json({
        status: 'healthy',
        service: 'personality-intelligence-service',
        version: '1.0.0',
        timestamp: new Date().toISOString()
    });
});

// API documentation endpoint
app.get('/', (req: Request, res: Response) => {
    res.json({
        service: 'Personality Intelligence Service',
        version: '1.0.0',
        description: 'Infers probabilistic personality dimensions through situational questions',
        endpoints: {
            'POST /profile/init': 'Create a new user personality profile',
            'GET /profile/:userId': 'Get user profile and personality state',
            'GET /question/next': 'Get the next adaptive question',
            'POST /answer': 'Submit an answer to a question',
            'GET /summary': 'Generate a personality summary'
        },
        documentation: {
            modes: ['ONBOARDING (6-8 questions)', 'STREAK_CHECKIN (1 question)'],
            dimensions: [
                'emotional_regulation',
                'communication_style',
                'emotional_availability',
                'consistency_style',
                'conflict_posture',
                'energy_orientation',
                'decision_pace'
            ]
        }
    });
});

// Mount route handlers
app.use('/profile', profileRoutes);
app.use('/question', questionRoutes);
app.use('/answer', answerRoutes);
app.use('/summary', summaryRoutes);

// ============================================================================
// Error Handling
// ============================================================================

// 404 handler
app.use((req: Request, res: Response) => {
    res.status(404).json({
        error: 'Endpoint not found',
        code: 'NOT_FOUND',
        path: req.path
    });
});

// Global error handler
app.use((err: Error, req: Request, res: Response, next: NextFunction) => {
    console.error('Unhandled error:', err);
    res.status(500).json({
        error: 'Internal server error',
        code: 'INTERNAL_ERROR',
        message: process.env.NODE_ENV === 'development' ? err.message : undefined
    });
});

// ============================================================================
// Server Startup
// ============================================================================

app.listen(PORT, () => {
    console.log('');
    console.log('╔════════════════════════════════════════════════════════════╗');
    console.log('║     Personality Intelligence Service                       ║');
    console.log('╠════════════════════════════════════════════════════════════╣');
    console.log(`║  🚀 Server running on http://localhost:${PORT}               ║`);
    console.log('║                                                            ║');
    console.log('║  Endpoints:                                                ║');
    console.log('║    POST /profile/init   - Create user profile              ║');
    console.log('║    GET  /question/next  - Get adaptive question            ║');
    console.log('║    POST /answer         - Submit answer                    ║');
    console.log('║    GET  /summary        - Generate personality summary     ║');
    console.log('╚════════════════════════════════════════════════════════════╝');
    console.log('');
});

export default app;
