/**
 * Hobby Intelligence Engine - Entry Point
 * 
 * A production-grade microservice for hobby validation, classification,
 * and daily check-in generation.
 * 
 * Endpoints:
 * - POST /validate  - Validate a hobby
 * - POST /classify  - Classify a hobby
 * - POST /question  - Generate check-in question
 * - POST /echo      - Generate privacy-safe update
 * - POST /pipeline  - Run full validation→classification→question flow
 * - GET  /health    - Health check
 * - GET  /          - Test UI
 */

import Fastify from 'fastify';
import cors from '@fastify/cors';
import fastifyStatic from '@fastify/static';
import path from 'path';
import { fileURLToPath } from 'url';
import { config } from './config/env.js';
import { hobbyRoutes } from './routes/hobby.routes.js';

// Get __dirname equivalent in ESM
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// Create Fastify instance with logging
const fastify = Fastify({
    logger: true,
});

// Start server
async function start() {
    try {
        // Register CORS for cross-origin requests
        await fastify.register(cors, {
            origin: true, // Allow all origins in development
        });

        // Register static file serving for test UI
        await fastify.register(fastifyStatic, {
            root: path.join(__dirname, '..', 'public'),
            prefix: '/',
        });

        // Register routes
        await fastify.register(hobbyRoutes);

        await fastify.listen({
            port: config.PORT,
            host: config.HOST,
        });

        console.log(`
╔══════════════════════════════════════════════════════════╗
║         🎯 Hobby Intelligence Engine                     ║
╠══════════════════════════════════════════════════════════╣
║  Server running at http://${config.HOST}:${config.PORT}                  ║
║                                                          ║
║  Test UI:    http://localhost:${config.PORT}/                         ║
║                                                          ║
║  Endpoints:                                              ║
║    POST /validate   - Validate a hobby                   ║
║    POST /classify   - Classify a hobby                   ║
║    POST /question   - Generate check-in question         ║
║    POST /echo       - Generate privacy-safe update       ║
║    POST /pipeline   - Run full pipeline                  ║
║    GET  /health     - Health check                       ║
╚══════════════════════════════════════════════════════════╝
    `);
    } catch (err) {
        fastify.log.error(err);
        process.exit(1);
    }
}

start();
