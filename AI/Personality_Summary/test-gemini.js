// List available Gemini models
const fs = require('fs');
const path = require('path');

// Read API key from .env
const envPath = path.join(__dirname, '.env');
const envContent = fs.readFileSync(envPath, 'utf8');
let apiKey = '';
for (const line of envContent.split('\n')) {
    if (line.trim().startsWith('GEMINI_API_KEY=')) {
        apiKey = line.split('=')[1]?.trim();
    }
}

async function listModels() {
    console.log('Fetching available Gemini models...\n');

    const url = `https://generativelanguage.googleapis.com/v1beta/models?key=${apiKey}`;

    const response = await fetch(url);
    const data = await response.json();

    if (data.models) {
        console.log('Available models for generateContent:\n');
        for (const model of data.models) {
            if (model.supportedGenerationMethods?.includes('generateContent')) {
                console.log(`  ✓ ${model.name.replace('models/', '')}`);
            }
        }
    } else {
        console.log('Error:', JSON.stringify(data, null, 2));
    }
}

listModels();
