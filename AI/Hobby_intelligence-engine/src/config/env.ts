/**
 * Environment Configuration
 * 
 * Loads and validates all environment variables at startup.
 * Uses Zod for runtime validation - fails fast if required vars are missing.
 */

import { z } from 'zod';
import dotenv from 'dotenv';

// Load .env file
dotenv.config();

// Define environment schema with validation
const envSchema = z.object({
    // Required: Hugging Face API Key
    HUGGINGFACE_API_KEY: z.string().min(1, 'HUGGINGFACE_API_KEY is required'),

    // Server config with defaults
    PORT: z.string().default('3000').transform(Number),
    HOST: z.string().default('0.0.0.0'),

    // AI model config with sensible defaults
    AI_MODEL: z.string().default('meta-llama/Llama-3.3-70B-Instruct'),
    AI_TEMPERATURE: z.string().default('0.3').transform(Number),
    AI_MAX_TOKENS: z.string().default('1000').transform(Number),
});

// Parse and validate environment
const parsed = envSchema.safeParse(process.env);

if (!parsed.success) {
    console.error('❌ Invalid environment configuration:');
    console.error(parsed.error.format());
    process.exit(1);
}

// Export typed configuration
export const config = parsed.data;

export type Config = typeof config;
