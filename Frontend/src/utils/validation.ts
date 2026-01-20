import {z} from 'zod';

// Common Regex patterns
export const EMAIL_REGEX = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
export const NAME_REGEX = /^[A-Za-z\s'-]+$/;
export const NO_ANGLE_BRACKETS_REGEX = /^[^<>]*$/;
export const PHONE_REGEX = /^\+?[1-9]\d{1,14}$/;
export const OTP_REGEX = /^\d{6}$/;

/**
 * Email field schema
 */
export const emailField = () =>
  z.string().min(1, 'Email is required').regex(EMAIL_REGEX, 'Invalid email');

/**
 * Password field schema
 */
export const passwordField = (minLength = 8) =>
  z
    .string()
    .min(minLength, `Password must be at least ${minLength} characters`);

/**
 * Name field schema
 */
export const nameField = (label = 'Name') =>
  z
    .string()
    .min(1, `${label} is required`)
    .regex(
      NAME_REGEX,
      'Only letters, spaces, hyphens, and apostrophes allowed',
    );

/**
 * Phone field schema
 */
export const phoneField = () =>
  z
    .string()
    .min(1, 'Phone is required')
    .regex(PHONE_REGEX, 'Invalid phone number');

/**
 * OTP field schema
 */
export const otpField = () =>
  z
    .string()
    .length(6, 'OTP must be 6 digits')
    .regex(OTP_REGEX, 'OTP must be numeric');

/**
 * Safe text field schema (no XSS characters)
 */
export const safeTextField = (label = 'Field') =>
  z
    .string()
    .min(1, `${label} is required`)
    .regex(NO_ANGLE_BRACKETS_REGEX, 'Invalid characters detected');

/**
 * Required string schema
 */
export const requiredStringSchema = z
  .string()
  .min(1, 'This field is required')
  .trim();

/**
 * URL schema (optional - allows empty string)
 */
export const urlSchema = z
  .string()
  .url('Invalid URL')
  .or(z.literal('').transform(() => undefined));

/**
 * Number schema (coerced from string)
 */
export const numberSchema = z.coerce
  .number()
  .min(1, 'Value must be greater than 0');

// Common form schemas
export const loginSchema = z.object({
  email: emailField(),
  password: passwordField(),
});

export const signupSchema = z
  .object({
    email: emailField(),
    password: passwordField(),
    confirmPassword: z.string().min(1, 'Please confirm your password'),
  })
  .refine(data => data.password === data.confirmPassword, {
    path: ['confirmPassword'],
    message: 'Passwords do not match',
  });

export const profileSchema = z.object({
  firstName: nameField('First name'),
  lastName: nameField('Last name'),
  bio: z.string().optional(),
});

// Type inference
export type LoginFormValues = z.infer<typeof loginSchema>;
export type SignupFormValues = z.infer<typeof signupSchema>;
export type ProfileFormValues = z.infer<typeof profileSchema>;
