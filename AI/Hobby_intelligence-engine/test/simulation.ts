/**
 * Hobby Intelligence Engine - Test Simulation
 * 
 * Runs a comprehensive test suite with 12 hobbies:
 * - Common valid hobbies (gym, painting, reading, cooking, gardening)
 * - Uncommon valid hobbies (urban sketching, lockpicking, beekeeping)
 * - Invalid hobbies (stalking, hacking people, catfishing, being funny)
 * 
 * For each hobby:
 * 1. Run validation
 * 2. If valid → run classification
 * 3. Generate a question
 * 4. Simulate user selecting first option
 * 5. Generate Hobby Echo
 * 
 * Prints full JSON outputs step-by-step.
 */

import { config } from '../src/config/env.js';
import { createAIClient } from '../src/ai/client.js';
import { createHobbyValidator } from '../src/modules/hobbyValidator.js';
import { createHobbyClassifier } from '../src/modules/hobbyClassifier.js';
import { createQuestionGenerator } from '../src/modules/questionGenerator.js';
import { createHobbyEchoGenerator } from '../src/modules/hobbyEchoGenerator.js';

// Test hobbies grouped by category
const TEST_HOBBIES = {
    common: [
        'gym',
        'painting',
        'reading',
        'cooking',
        'gardening',
    ],
    uncommon: [
        'urban sketching',
        'lockpicking (sport)',
        'beekeeping',
    ],
    invalid: [
        'stalking',
        'hacking people',
        'catfishing',
        'being funny', // personality trait, not hobby
    ],
};

// Colors for terminal output
const colors = {
    reset: '\x1b[0m',
    bright: '\x1b[1m',
    green: '\x1b[32m',
    red: '\x1b[31m',
    yellow: '\x1b[33m',
    cyan: '\x1b[36m',
    dim: '\x1b[2m',
};

function log(message: string, color: string = colors.reset) {
    console.log(`${color}${message}${colors.reset}`);
}

function printJson(label: string, data: unknown) {
    console.log(`${colors.dim}${label}:${colors.reset}`);
    console.log(JSON.stringify(data, null, 2));
    console.log();
}

function printSeparator(hobby: string) {
    console.log('\n' + '═'.repeat(60));
    log(`  TESTING: "${hobby}"`, colors.bright + colors.cyan);
    console.log('═'.repeat(60) + '\n');
}

async function runSimulation() {
    console.log(`
╔══════════════════════════════════════════════════════════════════╗
║         🧪 HOBBY INTELLIGENCE ENGINE - TEST SIMULATION           ║
╠══════════════════════════════════════════════════════════════════╣
║  Testing ${Object.values(TEST_HOBBIES).flat().length} hobbies across validation, classification,            ║
║  question generation, and echo generation.                       ║
╚══════════════════════════════════════════════════════════════════╝
  `);

    // Initialize AI client and modules
    const aiClient = createAIClient('openai');
    const validator = createHobbyValidator(aiClient);
    const classifier = createHobbyClassifier(aiClient);
    const questionGen = createQuestionGenerator(aiClient);
    const echoGen = createHobbyEchoGenerator(aiClient);

    // Stats tracking
    const stats = {
        total: 0,
        valid: 0,
        invalid: 0,
        errors: 0,
    };

    // Process all hobbies
    const allHobbies = [
        ...TEST_HOBBIES.common.map(h => ({ hobby: h, expected: 'valid' })),
        ...TEST_HOBBIES.uncommon.map(h => ({ hobby: h, expected: 'valid' })),
        ...TEST_HOBBIES.invalid.map(h => ({ hobby: h, expected: 'invalid' })),
    ];

    for (const { hobby, expected } of allHobbies) {
        stats.total++;
        printSeparator(hobby);

        try {
            // Step 1: Validation
            log('📋 Step 1: Validation', colors.yellow);
            const validation = await validator.validate({ hobby });
            printJson('Validation Result', validation);

            if (!validation.legit) {
                log('❌ Hobby rejected - stopping pipeline', colors.red);
                stats.invalid++;

                if (expected === 'valid') {
                    log('⚠️  UNEXPECTED: Expected this hobby to be valid!', colors.red);
                } else {
                    log('✓ Expected rejection confirmed', colors.green);
                }
                continue;
            }

            stats.valid++;

            if (expected === 'invalid') {
                log('⚠️  UNEXPECTED: Expected this hobby to be rejected!', colors.yellow);
            } else {
                log('✓ Hobby validated', colors.green);
            }

            // Step 2: Classification
            log('\n📊 Step 2: Classification', colors.yellow);
            const classification = await classifier.classify({ hobby });
            printJson('Classification Result', classification);

            // Step 3: Question Generation
            log('❓ Step 3: Question Generation', colors.yellow);
            const question = await questionGen.generate({
                hobby,
                category: classification.category,
                previous_questions: [],
            });
            printJson('Question Result', question);

            // Step 4: Simulate user selection (pick first option)
            const selectedOption = question.options[0];
            log(`👆 Step 4: Simulating user selection: "${selectedOption}"`, colors.yellow);

            // Step 5: Echo Generation
            log('\n🔊 Step 5: Echo Generation', colors.yellow);
            const echo = await echoGen.generate({
                hobby,
                selected_option: selectedOption,
            });
            printJson('Echo Result', echo);

            log('✓ Full pipeline completed successfully', colors.green);

        } catch (error) {
            stats.errors++;
            log('💥 Error during processing:', colors.red);
            console.error(error);
        }
    }

    // Print summary
    console.log('\n' + '═'.repeat(60));
    log('  📊 SIMULATION SUMMARY', colors.bright + colors.cyan);
    console.log('═'.repeat(60));
    console.log(`
  Total hobbies tested: ${stats.total}
  Valid (full pipeline): ${stats.valid}
  Invalid (rejected):    ${stats.invalid}
  Errors:                ${stats.errors}
  `);

    if (stats.errors > 0) {
        log('⚠️  Some tests had errors - check output above', colors.yellow);
        process.exit(1);
    } else {
        log('✓ All tests completed successfully!', colors.green);
    }
}

// Run the simulation
runSimulation().catch(console.error);
