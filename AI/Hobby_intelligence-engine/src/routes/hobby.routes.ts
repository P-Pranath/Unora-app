/**
 * Hobby Routes
 * 
 * API endpoints for all Hobby Intelligence Engine operations.
 * Each module is exposed as a separate POST endpoint.
 * Pipeline endpoint runs the full flow in sequence.
 */

import { FastifyInstance, FastifyRequest, FastifyReply } from 'fastify';
import { createAIClient } from '../ai/client.js';
import { createHobbyValidator } from '../modules/hobbyValidator.js';
import { createHobbyClassifier } from '../modules/hobbyClassifier.js';
import { createQuestionGenerator } from '../modules/questionGenerator.js';
import { createHobbyEchoGenerator } from '../modules/hobbyEchoGenerator.js';
import {
    hobbyValidatorInputSchema,
    hobbyClassifierInputSchema,
    questionGeneratorInputSchema,
    hobbyEchoInputSchema,
    pipelineInputSchema,
    HobbyValidatorInput,
    HobbyClassifierInput,
    QuestionGeneratorInput,
    HobbyEchoInput,
    PipelineInput,
    PipelineOutput,
} from '../types/schemas.js';

// Initialize AI client and modules once (uses Hugging Face by default)
const aiClient = createAIClient();
const validator = createHobbyValidator(aiClient);
const classifier = createHobbyClassifier(aiClient);
const questionGen = createQuestionGenerator(aiClient);
const echoGen = createHobbyEchoGenerator(aiClient);

/**
 * Registers all hobby-related routes
 */
export async function hobbyRoutes(fastify: FastifyInstance) {

    /**
     * POST /validate
     * Validates if a hobby is legitimate and safe
     */
    fastify.post('/validate', async (request: FastifyRequest, reply: FastifyReply) => {
        const parsed = hobbyValidatorInputSchema.safeParse(request.body);

        if (!parsed.success) {
            return reply.status(400).send({
                error: 'Invalid input',
                details: parsed.error.flatten(),
            });
        }

        try {
            const result = await validator.validate(parsed.data);
            return reply.send(result);
        } catch (error) {
            fastify.log.error(error);
            return reply.status(500).send({ error: 'Validation failed' });
        }
    });

    /**
     * POST /classify
     * Classifies a hobby into category, effort type, and cadence
     */
    fastify.post('/classify', async (request: FastifyRequest, reply: FastifyReply) => {
        const parsed = hobbyClassifierInputSchema.safeParse(request.body);

        if (!parsed.success) {
            return reply.status(400).send({
                error: 'Invalid input',
                details: parsed.error.flatten(),
            });
        }

        try {
            const result = await classifier.classify(parsed.data);
            return reply.send(result);
        } catch (error) {
            fastify.log.error(error);
            return reply.status(500).send({ error: 'Classification failed' });
        }
    });

    /**
     * POST /question
     * Generates a daily check-in question with tap options
     */
    fastify.post('/question', async (request: FastifyRequest, reply: FastifyReply) => {
        const parsed = questionGeneratorInputSchema.safeParse(request.body);

        if (!parsed.success) {
            return reply.status(400).send({
                error: 'Invalid input',
                details: parsed.error.flatten(),
            });
        }

        try {
            const result = await questionGen.generate(parsed.data);
            return reply.send(result);
        } catch (error) {
            fastify.log.error(error);
            return reply.status(500).send({ error: 'Question generation failed' });
        }
    });

    /**
     * POST /echo
     * Generates a privacy-safe progress update
     */
    fastify.post('/echo', async (request: FastifyRequest, reply: FastifyReply) => {
        const parsed = hobbyEchoInputSchema.safeParse(request.body);

        if (!parsed.success) {
            return reply.status(400).send({
                error: 'Invalid input',
                details: parsed.error.flatten(),
            });
        }

        try {
            const result = await echoGen.generate(parsed.data);
            return reply.send(result);
        } catch (error) {
            fastify.log.error(error);
            return reply.status(500).send({ error: 'Echo generation failed' });
        }
    });

    /**
     * POST /pipeline
     * Runs the full pipeline: validate → classify → question
     * Stops early if validation fails
     */
    fastify.post('/pipeline', async (request: FastifyRequest, reply: FastifyReply) => {
        const parsed = pipelineInputSchema.safeParse(request.body);

        if (!parsed.success) {
            return reply.status(400).send({
                error: 'Invalid input',
                details: parsed.error.flatten(),
            });
        }

        try {
            const { hobby, previous_questions } = parsed.data;

            // Step 1: Validate
            const validation = await validator.validate({ hobby });

            // If not legit, stop here
            if (!validation.legit) {
                const result: PipelineOutput = { validation };
                return reply.send(result);
            }

            // Step 2: Classify
            const classification = await classifier.classify({ hobby });

            // Step 3: Generate question
            const question = await questionGen.generate({
                hobby,
                category: classification.category,
                previous_questions,
            });

            const result: PipelineOutput = {
                validation,
                classification,
                question,
            };

            return reply.send(result);
        } catch (error) {
            fastify.log.error(error);
            return reply.status(500).send({ error: 'Pipeline failed' });
        }
    });

    /**
     * GET /health
     * Health check endpoint
     */
    fastify.get('/health', async (_request: FastifyRequest, reply: FastifyReply) => {
        return reply.send({ status: 'ok', service: 'hobby-intelligence-engine' });
    });
}
