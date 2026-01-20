import { IAIClient } from './ai.client.interface';
import { AIValidationResult, IssueTag } from '../schemas/validation.schema';

/**
 * Mock AI Client
 * 
 * Deterministic mock implementation for testing.
 * Returns results based on filename patterns:
 * 
 * - *_pass.* → usable, high confidence
 * - *_no_face.* → no_face issue
 * - *_blur.* → blurred issue
 * - *_dark.* → too_dark issue
 * - *_bright.* → too_bright issue
 * - *_obstructed.* → obstructed_face issue
 * - *_screenshot.* → non_photographic issue
 * - *_lowres.* → low_resolution issue
 * - Default → random with 70% pass rate
 */
export class MockAIClient implements IAIClient {
    private filenameHint: string = '';
    private forceResult: AIValidationResult | null = null;

    constructor() {
        console.log('🧪 Using Mock AI Client for testing');
    }

    isReady(): boolean {
        return true;
    }

    getName(): string {
        return 'Mock AI Client (Testing)';
    }

    /**
     * Set a filename hint for pattern-based responses
     * This should be called before validatePhoto
     */
    setFilenameHint(filename: string): void {
        this.filenameHint = filename.toLowerCase();
    }

    /**
     * Force a specific result for the next validation
     */
    setForceResult(result: AIValidationResult | null): void {
        this.forceResult = result;
    }

    /**
     * Validate photo with deterministic mock behavior
     */
    async validatePhoto(imageBase64: string): Promise<AIValidationResult> {
        // Simulate processing delay
        await new Promise(resolve => setTimeout(resolve, 100));

        // Check for forced result
        if (this.forceResult) {
            const result = this.forceResult;
            this.forceResult = null; // Reset after use
            return result;
        }

        // Pattern-based responses
        const filename = this.filenameHint;

        // Clear hint after use
        this.filenameHint = '';

        // Check for specific patterns
        if (filename.includes('_pass') || filename.includes('valid')) {
            return {
                usable: true,
                issues: [],
                confidence: 'high'
            };
        }

        if (filename.includes('no_face') || filename.includes('noface')) {
            return {
                usable: false,
                issues: ['no_face'],
                confidence: 'high'
            };
        }

        if (filename.includes('blur')) {
            return {
                usable: false,
                issues: ['blurred'],
                confidence: 'high'
            };
        }

        if (filename.includes('dark')) {
            return {
                usable: false,
                issues: ['too_dark'],
                confidence: 'high'
            };
        }

        if (filename.includes('bright')) {
            return {
                usable: false,
                issues: ['too_bright'],
                confidence: 'high'
            };
        }

        if (filename.includes('obstructed') || filename.includes('masked')) {
            return {
                usable: false,
                issues: ['obstructed_face'],
                confidence: 'high'
            };
        }

        if (filename.includes('screenshot') || filename.includes('meme')) {
            return {
                usable: false,
                issues: ['non_photographic'],
                confidence: 'high'
            };
        }

        if (filename.includes('lowres') || filename.includes('small')) {
            return {
                usable: false,
                issues: ['low_resolution'],
                confidence: 'high'
            };
        }

        if (filename.includes('multi_issue')) {
            return {
                usable: false,
                issues: ['blurred', 'too_dark'],
                confidence: 'medium'
            };
        }

        if (filename.includes('borderline')) {
            return {
                usable: true,
                issues: [],
                confidence: 'medium'
            };
        }

        // Default: random with 70% pass rate
        const random = Math.random();

        if (random < 0.7) {
            return {
                usable: true,
                issues: [],
                confidence: 'high'
            };
        } else {
            // Random issue for failures
            const possibleIssues: IssueTag[] = ['no_face', 'blurred', 'obstructed_face'];
            const randomIssue = possibleIssues[Math.floor(Math.random() * possibleIssues.length)];

            return {
                usable: false,
                issues: [randomIssue],
                confidence: 'medium'
            };
        }
    }
}

/**
 * Create mock client instance
 */
export function createMockAIClient(): IAIClient {
    return new MockAIClient();
}

/**
 * Extended mock client for more control in tests
 */
export class TestMockAIClient extends MockAIClient {
    private callCount: number = 0;
    private callLog: Array<{ timestamp: number; imageSize: number }> = [];

    async validatePhoto(imageBase64: string): Promise<AIValidationResult> {
        this.callCount++;
        this.callLog.push({
            timestamp: Date.now(),
            imageSize: imageBase64.length
        });

        return super.validatePhoto(imageBase64);
    }

    getCallCount(): number {
        return this.callCount;
    }

    getCallLog(): Array<{ timestamp: number; imageSize: number }> {
        return [...this.callLog];
    }

    reset(): void {
        this.callCount = 0;
        this.callLog = [];
        this.setForceResult(null);
        this.setFilenameHint('');
    }
}
