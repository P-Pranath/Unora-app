import dotenv from 'dotenv';

// Load environment variables
dotenv.config();

export interface Config {
    port: number;
    storagePath: string;
    openaiApiKey: string | undefined;
    useMockAI: boolean;

    // CV Validation Thresholds
    cv: {
        blurThreshold: number;
        brightnessMin: number;
        brightnessMax: number;
        minWidth: number;
        minHeight: number;
    };
}

export const config: Config = {
    port: parseInt(process.env.PORT || '3000', 10),
    storagePath: process.env.STORAGE_PATH || './storage/photos',
    openaiApiKey: process.env.OPENAI_API_KEY,
    useMockAI: process.env.USE_MOCK_AI === 'true',

    cv: {
        blurThreshold: parseInt(process.env.BLUR_THRESHOLD || '100', 10),
        brightnessMin: parseInt(process.env.BRIGHTNESS_MIN || '30', 10),
        brightnessMax: parseInt(process.env.BRIGHTNESS_MAX || '225', 10),
        minWidth: parseInt(process.env.MIN_WIDTH || '640', 10),
        minHeight: parseInt(process.env.MIN_HEIGHT || '480', 10),
    }
};

// Validate configuration
export function validateConfig(): void {
    if (!config.useMockAI && !config.openaiApiKey) {
        console.warn('⚠️  Warning: OPENAI_API_KEY not set. Using mock AI client.');
        config.useMockAI = true;
    }
}
