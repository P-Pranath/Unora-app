/**
 * Zod Schemas for Hobby Intelligence Engine
 * 
 * Runtime type validation for all module inputs/outputs.
 * These schemas are the source of truth for data structures.
 * Used for API validation and internal type safety.
 */

import { z } from 'zod';

// ============================================
// HOBBY VALIDATOR SCHEMAS
// ============================================

export const hobbyValidatorInputSchema = z.object({
    hobby: z.string().min(1, 'Hobby is required').max(200, 'Hobby too long'),
});

export const riskLevelSchema = z.enum(['low', 'medium', 'high']);

export const hobbyValidatorOutputSchema = z.object({
    legit: z.boolean(),
    reason: z.string().nullable(),
    risk_level: riskLevelSchema,
});

export type HobbyValidatorInput = z.infer<typeof hobbyValidatorInputSchema>;
export type HobbyValidatorOutput = z.infer<typeof hobbyValidatorOutputSchema>;
export type RiskLevel = z.infer<typeof riskLevelSchema>;

// ============================================
// HOBBY CLASSIFIER SCHEMAS
// ============================================

export const hobbyCategorySchema = z.enum([
    'Physical',
    'Creative',
    'Skill / Practice',
    'Learning',
    'Mindfulness',
    'Social',
    'Collection',
    'Maintenance',
]);

export const effortTypeSchema = z.enum(['solo', 'social', 'mixed']);

export const cadenceSchema = z.enum(['daily', 'irregular', 'flexible']);

export const hobbyClassifierInputSchema = z.object({
    hobby: z.string().min(1),
});

export const hobbyClassifierOutputSchema = z.object({
    category: hobbyCategorySchema,
    effort_type: effortTypeSchema,
    cadence: cadenceSchema,
});

export type HobbyCategory = z.infer<typeof hobbyCategorySchema>;
export type EffortType = z.infer<typeof effortTypeSchema>;
export type Cadence = z.infer<typeof cadenceSchema>;
export type HobbyClassifierInput = z.infer<typeof hobbyClassifierInputSchema>;
export type HobbyClassifierOutput = z.infer<typeof hobbyClassifierOutputSchema>;

// ============================================
// QUESTION GENERATOR SCHEMAS
// ============================================

export const questionGeneratorInputSchema = z.object({
    hobby: z.string().min(1),
    category: hobbyCategorySchema,
    previous_questions: z.array(z.string()).default([]),
});

export const questionGeneratorOutputSchema = z.object({
    question: z.string().min(1),
    options: z.array(z.string()).min(2).max(4),
});

export type QuestionGeneratorInput = z.infer<typeof questionGeneratorInputSchema>;
export type QuestionGeneratorOutput = z.infer<typeof questionGeneratorOutputSchema>;

// ============================================
// HOBBY ECHO GENERATOR SCHEMAS
// ============================================

export const hobbyEchoInputSchema = z.object({
    hobby: z.string().min(1),
    selected_option: z.string().min(1),
});

export const hobbyEchoOutputSchema = z.object({
    echo: z.string().min(1),
});

export type HobbyEchoInput = z.infer<typeof hobbyEchoInputSchema>;
export type HobbyEchoOutput = z.infer<typeof hobbyEchoOutputSchema>;

// ============================================
// PIPELINE SCHEMAS (Full Flow)
// ============================================

export const pipelineInputSchema = z.object({
    hobby: z.string().min(1).max(200),
    previous_questions: z.array(z.string()).default([]),
});

export const pipelineOutputSchema = z.object({
    validation: hobbyValidatorOutputSchema,
    classification: hobbyClassifierOutputSchema.optional(),
    question: questionGeneratorOutputSchema.optional(),
});

export type PipelineInput = z.infer<typeof pipelineInputSchema>;
export type PipelineOutput = z.infer<typeof pipelineOutputSchema>;
