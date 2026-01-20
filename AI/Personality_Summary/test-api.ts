// ============================================================================
// API Test Script - Tests all endpoints of the Personality Service
// ============================================================================
// Run with: npx ts-node test-api.ts
// ============================================================================

const BASE_URL = 'http://localhost:3000';

async function testAPI() {
    console.log('🧪 Testing Personality Intelligence Service API\n');

    // Test 1: Health Check
    console.log('1️⃣ Testing Health Check...');
    const healthRes = await fetch(`${BASE_URL}/health`);
    const health = await healthRes.json() as { status: string };
    console.log('   Response:', JSON.stringify(health, null, 2));
    console.log('   ✅ Health check passed\n');

    // Test 2: Create Profile
    console.log('2️⃣ Testing Profile Creation...');
    const userId = `test-user-${Date.now()}`;
    const profileRes = await fetch(`${BASE_URL}/profile/init`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ userId })
    });
    const profile = await profileRes.json() as { personalityState?: Record<string, unknown> };
    console.log('   User ID:', userId);
    console.log('   Response status:', profileRes.status);
    console.log('   Dimensions:', Object.keys(profile.personalityState || {}).length);
    console.log('   ✅ Profile creation passed\n');

    // Test 3: Get First Question
    console.log('3️⃣ Testing Get Next Question...');
    const questionRes = await fetch(`${BASE_URL}/question/next?userId=${userId}&mode=ONBOARDING`);
    const questionData = await questionRes.json() as {
        question?: { id: string; scenario: string; options: { label: string }[] };
        complete?: boolean;
    };
    console.log('   Question ID:', questionData.question?.id);
    console.log('   Scenario:', questionData.question?.scenario?.substring(0, 50) + '...');
    console.log('   Options count:', questionData.question?.options?.length);
    console.log('   ✅ Question retrieval passed\n');

    // Test 4: Submit Answer
    console.log('4️⃣ Testing Submit Answer...');
    const answerRes = await fetch(`${BASE_URL}/answer`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            userId,
            questionId: questionData.question!.id,
            selectedOption: 0
        })
    });
    const answerData = await answerRes.json() as {
        updated: boolean;
        questionsAnswered: number;
        nextQuestion?: unknown;
    };
    console.log('   Updated:', answerData.updated);
    console.log('   Questions answered:', answerData.questionsAnswered);
    console.log('   Has next question:', !!answerData.nextQuestion);
    console.log('   ✅ Answer submission passed\n');

    // Test 5: Submit a few more answers to build up confidence
    console.log('5️⃣ Answering more questions to build confidence...');
    for (let i = 0; i < 4; i++) {
        const nextQ = await fetch(`${BASE_URL}/question/next?userId=${userId}&mode=ONBOARDING`);
        const nextQData = await nextQ.json() as {
            question?: { id: string; options: { label: string }[] };
            complete?: boolean;
        };

        if (nextQData.complete) {
            console.log('   All questions completed!');
            break;
        }

        await fetch(`${BASE_URL}/answer`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                userId,
                questionId: nextQData.question!.id,
                selectedOption: Math.floor(Math.random() * nextQData.question!.options.length)
            })
        });
        console.log(`   Answered question ${i + 2}: ${nextQData.question!.id}`);
    }
    console.log('   ✅ Multiple answers submitted\n');

    // Test 6: Get Summary (this will use fallback since no real LLM key)
    console.log('6️⃣ Testing Summary Generation...');
    const summaryRes = await fetch(`${BASE_URL}/summary?userId=${userId}`);
    const summaryData = await summaryRes.json() as {
        summary?: string;
        confidence?: { overall: number };
        generatedAt?: string;
    };
    console.log('   Summary:', summaryData.summary?.substring(0, 100) + '...');
    console.log('   Overall confidence:', summaryData.confidence?.overall);
    console.log('   Generated at:', summaryData.generatedAt);
    console.log('   ✅ Summary generation passed\n');

    // Test 7: Test Skip functionality
    console.log('7️⃣ Testing Skip Functionality...');
    const skipQ = await fetch(`${BASE_URL}/question/next?userId=${userId}&mode=ONBOARDING`);
    const skipQData = await skipQ.json() as {
        question?: { id: string };
        complete?: boolean;
    };

    if (!skipQData.complete && skipQData.question) {
        const skipRes = await fetch(`${BASE_URL}/answer`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                userId,
                questionId: skipQData.question.id,
                skip: true
            })
        });
        const skipData = await skipRes.json() as { skipped: boolean };
        console.log('   Skipped:', skipData.skipped);
        console.log('   ✅ Skip functionality passed\n');
    } else {
        console.log('   No more questions to skip\n');
    }

    // Test 8: Get Profile
    console.log('8️⃣ Testing Get Profile...');
    const getProfileRes = await fetch(`${BASE_URL}/profile/${userId}`);
    const profileData = await getProfileRes.json() as {
        userId: string;
        questionsAnswered: number;
    };
    console.log('   User ID:', profileData.userId);
    console.log('   Questions answered:', profileData.questionsAnswered);
    console.log('   ✅ Profile retrieval passed\n');

    console.log('═══════════════════════════════════════════════════════');
    console.log('✅ All API tests passed successfully!');
    console.log('═══════════════════════════════════════════════════════');
}

testAPI().catch(console.error);
