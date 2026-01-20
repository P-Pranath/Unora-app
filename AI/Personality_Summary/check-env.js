// Quick check for environment configuration (no dependencies)
const fs = require('fs');
const path = require('path');

console.log('\n=== Environment Check ===\n');

// Read .env file directly
const envPath = path.join(__dirname, '.env');

if (!fs.existsSync(envPath)) {
    console.log('❌ .env file not found!');
    console.log('Create a .env file with:');
    console.log('  GEMINI_API_KEY=your_key_here');
    console.log('  AI_MODEL=gemini-2.0-flash');
    process.exit(1);
}

const envContent = fs.readFileSync(envPath, 'utf8');
const lines = envContent.split('\n');

let geminiKey = null;
let openaiKey = null;
let aiModel = null;

for (const line of lines) {
    const trimmed = line.trim();
    if (trimmed.startsWith('#') || !trimmed) continue;

    if (trimmed.startsWith('GEMINI_API_KEY=')) {
        geminiKey = trimmed.split('=')[1]?.trim();
    }
    if (trimmed.startsWith('OPENAI_API_KEY=')) {
        openaiKey = trimmed.split('=')[1]?.trim();
    }
    if (trimmed.startsWith('AI_MODEL=')) {
        aiModel = trimmed.split('=')[1]?.trim();
    }
}

console.log('GEMINI_API_KEY:', geminiKey
    ? `✓ SET (${geminiKey.substring(0, 10)}...)`
    : '✗ NOT SET');

console.log('OPENAI_API_KEY:', openaiKey
    ? `SET (${openaiKey.substring(0, 10)}...)`
    : 'NOT SET');

console.log('AI_MODEL:', aiModel || 'not set (will default to gemini-2.0-flash)');

console.log('\n=== Detection Result ===\n');

if (geminiKey && !geminiKey.includes('your_') && geminiKey.length > 10) {
    console.log('✓ Will use: GEMINI');
    console.log('  Endpoint: https://generativelanguage.googleapis.com/v1beta/openai/');
    console.log('\n✓ Your .env looks correct! Restart the server with: npm run dev');
} else if (openaiKey && !openaiKey.includes('your_')) {
    console.log('Will use: OpenAI');
} else {
    console.log('⚠️  No valid API key found!');
    console.log('\nYour .env should have:');
    console.log('  GEMINI_API_KEY=AIzaSy....your_actual_key');
    console.log('  AI_MODEL=gemini-2.0-flash');
}

console.log('\n');
