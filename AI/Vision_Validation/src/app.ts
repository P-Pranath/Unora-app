import express, { Application } from 'express';
import cors from 'cors';
import path from 'path';
import onboardingRoutes from './routes/onboarding.routes';
import { errorHandler, notFoundHandler } from './middleware/error.middleware';

/**
 * Create and configure Express application
 */
export function createApp(): Application {
    const app = express();

    // Middleware
    app.use(cors());
    app.use(express.json());
    app.use(express.urlencoded({ extended: true }));

    // Request logging (development)
    if (process.env.NODE_ENV !== 'production') {
        app.use((req, res, next) => {
            console.log(`[${new Date().toISOString()}] ${req.method} ${req.path}`);
            next();
        });
    }

    // Routes
    app.use('/onboarding', onboardingRoutes);

    // Serve static files (Test UI)
    app.use(express.static(path.join(__dirname, '../public')));

    // Root endpoint - serve test UI
    app.get('/', (req, res) => {
        res.sendFile(path.join(__dirname, '../public/index.html'));
    });

    // API info endpoint
    app.get('/api', (req, res) => {
        res.json({
            service: 'Vision Validation Service',
            version: '1.0.0',
            description: 'AI-powered photo quality and face presence validation for Unora onboarding',
            endpoints: {
                'POST /onboarding/photos': 'Upload and validate photos',
                'GET /onboarding/photos/:user_id': 'Get photo metadata',
                'GET /onboarding/health': 'Health check'
            },
            testUI: 'http://localhost:3000'
        });
    });

    // Error handling
    app.use(notFoundHandler);
    app.use(errorHandler);

    return app;
}

export default createApp;
