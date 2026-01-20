import { createApp } from './app';
import { config, validateConfig } from './config/env';
import { getAIClientInfo } from './services/ai.service';

/**
 * Vision Validation Service Entry Point
 */

// Validate configuration
validateConfig();

// Create Express app
const app = createApp();

// Start server
const server = app.listen(config.port, () => {
    console.log('\n' + '═'.repeat(60));
    console.log('  Vision Validation Service');
    console.log('═'.repeat(60));
    console.log(`  🚀 Server running on http://localhost:${config.port}`);
    console.log(`  📁 Storage path: ${config.storagePath}`);

    const aiInfo = getAIClientInfo();
    console.log(`  🤖 AI Client: ${aiInfo.name} (${aiInfo.ready ? 'ready' : 'not ready'})`);

    if (config.useMockAI) {
        console.log('  🧪 Running in TEST MODE with mock AI client');
    }

    console.log('═'.repeat(60));
    console.log('\n  Endpoints:');
    console.log('  POST /onboarding/photos     - Upload and validate photos');
    console.log('  GET  /onboarding/photos/:id - Get photo metadata');
    console.log('  GET  /onboarding/health     - Health check');
    console.log('═'.repeat(60) + '\n');
});

// Graceful shutdown
process.on('SIGTERM', () => {
    console.log('SIGTERM received, shutting down gracefully...');
    server.close(() => {
        console.log('Server closed');
        process.exit(0);
    });
});

process.on('SIGINT', () => {
    console.log('\nSIGINT received, shutting down gracefully...');
    server.close(() => {
        console.log('Server closed');
        process.exit(0);
    });
});

export default app;
