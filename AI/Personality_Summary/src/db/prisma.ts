// ============================================================================
// Prisma Client Singleton
// ============================================================================
// Provides a single PrismaClient instance for database operations.
// Prevents multiple client instances in development with hot reloading.
// ============================================================================

import { PrismaClient } from '@prisma/client';

// Extend the global object to store the Prisma client
declare global {
    // eslint-disable-next-line no-var
    var __prisma: PrismaClient | undefined;
}

// Create a singleton instance of PrismaClient
// In development, store it on the global object to prevent multiple instances
const prisma = global.__prisma || new PrismaClient({
    log: process.env.NODE_ENV === 'development' ? ['query', 'error', 'warn'] : ['error'],
});

if (process.env.NODE_ENV !== 'production') {
    global.__prisma = prisma;
}

export default prisma;
