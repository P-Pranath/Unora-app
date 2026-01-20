/**
 * Quick test script to verify Hugging Face API connection
 * Run with: npx tsx test-hf.ts
 */

import { HfInference } from '@huggingface/inference';
import dotenv from 'dotenv';

// Load .env
dotenv.config();

const API_KEY = process.env.HUGGINGFACE_API_KEY;
const MODEL = process.env.AI_MODEL || 'mistralai/Mistral-7B-Instruct-v0.3';

console.log('='.repeat(50));
console.log('Hugging Face Connection Test');
console.log('='.repeat(50));
console.log('API Key:', API_KEY ? `${API_KEY.substring(0, 10)}...` : 'NOT SET!');
console.log('Model:', MODEL);
console.log('='.repeat(50));

if (!API_KEY) {
    console.error('❌ ERROR: HUGGINGFACE_API_KEY is not set in .env file!');
    process.exit(1);
}

async function testConnection() {
    console.log('\n🔄 Testing Hugging Face API...\n');

    try {
        const client = new HfInference(API_KEY);

        const response = await client.chatCompletion({
            model: MODEL,
            messages: [
                { role: 'user', content: 'Say "Hello, test successful!" and nothing else.' }
            ],
            max_tokens: 50,
        });

        console.log('✅ SUCCESS! Response:');
        console.log(response.choices[0]?.message?.content);
        console.log('\n✅ Hugging Face API is working correctly!');

    } catch (error: any) {
        console.error('❌ FAILED! Error details:');
        console.error('Message:', error.message);
        console.error('Full error:', error);

        if (error.message?.includes('401')) {
            console.log('\n💡 Tip: Your API key is invalid. Get a new one from https://huggingface.co/settings/tokens');
        } else if (error.message?.includes('403')) {
            console.log('\n💡 Tip: You need to accept the license for this model on Hugging Face');
        }
    }
}

testConnection();
