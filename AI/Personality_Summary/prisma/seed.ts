// ============================================================================
// Database Seed Script - Question Bank Population
// ============================================================================
// Seeds the database with all questions from the question bank.
// Run with: npx ts-node prisma/seed.ts
// ============================================================================

import { PrismaClient } from '@prisma/client';
import { QUESTION_BANK } from '../src/questions/questionBank';

const prisma = new PrismaClient();

async function main() {
    console.log('🌱 Starting database seed...');
    console.log('');

    // Clear existing questions
    console.log('📋 Clearing existing questions...');
    await prisma.question.deleteMany({});

    // Seed questions
    console.log(`📝 Seeding ${QUESTION_BANK.length} questions...`);

    for (const question of QUESTION_BANK) {
        await prisma.question.create({
            data: {
                id: question.id,
                scenarioText: question.scenario,
                dimensionTargets: JSON.stringify(question.dimensionTargets),
                options: JSON.stringify(question.options)
            }
        });
    }

    console.log('');
    console.log('✅ Database seeded successfully!');
    console.log('');
    console.log('Question count by dimension:');

    // Count questions per dimension
    const dimensionCounts: Record<string, number> = {};
    for (const question of QUESTION_BANK) {
        for (const dim of question.dimensionTargets) {
            dimensionCounts[dim] = (dimensionCounts[dim] || 0) + 1;
        }
    }

    for (const [dim, count] of Object.entries(dimensionCounts).sort()) {
        console.log(`  - ${dim}: ${count} questions`);
    }

    console.log('');
}

main()
    .catch((e) => {
        console.error('❌ Seed failed:', e);
        process.exit(1);
    })
    .finally(async () => {
        await prisma.$disconnect();
    });
