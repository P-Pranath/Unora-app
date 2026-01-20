# React Hook Form + Zod for React Native - Complete Guide

This guide adapts React Hook Form (RHF) + Zod form management patterns for **React Native mobile apps**, including mobile-specific input handling, keyboard management, validation patterns, and reusable form components.

## Core Principles

- **Schema-first validation**: Define Zod schemas with type inference (`z.infer<typeof schema>`) for compile-time safety.
- **Centralized validation logic**: Keep all validation rules in schemas, not scattered in components.
- **Consistent error messaging**: User-friendly, actionable error messages defined in schemas.
- **Mobile-optimized inputs**: Use appropriate `keyboardType`, `inputMode`, `autoCapitalize`, `autoCorrect` for each field.
- **Keyboard-aware UI**: Integrate `KeyboardAvoidingView` and `ScrollView` for forms with many inputs.
- **Reusable form components**: Build a small UI layer around RHF for consistency and accessibility.

---

## Libraries and Dependencies

### Install Required Packages

```bash
npm install react-hook-form zod @hookform/resolvers
# or
yarn add react-hook-form zod @hookform/resolvers
```

### Core APIs Used

- **react-hook-form**: `useForm`, `Controller`, `useFieldArray`, `useFormContext`, `FormProvider`
- **zod**: `z.object`, `z.string`, `z.number`, `z.array`, `z.refine`, `z.coerce`, `z.infer`
- **@hookform/resolvers/zod**: `zodResolver` to bridge Zod schemas into RHF

---

## High-Level Architecture

```
┌──────────────────────────────────────────────────────────────┐
│  Screen Component                                             │
│    ├─ useForm({ resolver: zodResolver(schema) })            │
│    ├─ Form state (values, errors, isSubmitting)             │
│    └─ onSubmit handler (sanitize, call API mutation)        │
├──────────────────────────────────────────────────────────────┤
│  Form Components (Reusable UI Layer)                         │
│    ├─ FormField: wraps Controller + field rendering         │
│    ├─ FormLabel: accessible labels                          │
│    ├─ FormError: consistent error display                   │
│    └─ FormInput: styled TextInput with RHF integration      │
├──────────────────────────────────────────────────────────────┤
│  Zod Schemas (Validation Layer)                              │
│    ├─ Field-level rules (min, max, regex, custom)           │
│    ├─ Cross-field validation (refine, superRefine)          │
│    └─ Type inference for TypeScript                         │
└──────────────────────────────────────────────────────────────┘
```

---

## 1. Reusable Form Components (`src/components/form/`)

Build a small UI layer around RHF for consistency across all forms.

### `src/components/form/FormField.tsx`

Wraps RHF's `Controller` and provides a render prop pattern.

```tsx
import React from 'react';
import {Controller, Control, FieldPath, FieldValues} from 'react-hook-form';

export type FormFieldProps<
  TFieldValues extends FieldValues = FieldValues,
  TName extends FieldPath<TFieldValues> = FieldPath<TFieldValues>,
> = {
  control: Control<TFieldValues>;
  name: TName;
  render: ({field}: {field: any}) => React.ReactElement;
};

export function FormField<
  TFieldValues extends FieldValues = FieldValues,
  TName extends FieldPath<TFieldValues> = FieldPath<TFieldValues>,
>({control, name, render}: FormFieldProps<TFieldValues, TName>) {
  return <Controller control={control} name={name} render={render} />;
}
```

### `src/components/form/FormInput.tsx`

Standard text input component integrated with RHF.

```tsx
import React from 'react';
import {StyleSheet, TextInput, TextInputProps, View} from 'react-native';
import {spacing} from '../../theme/spacing';
import {typography} from '../../theme/typography';
import {sizes} from '../../theme/sizes';

export type FormInputProps = TextInputProps & {
  error?: boolean;
};

export const FormInput = React.forwardRef<TextInput, FormInputProps>(
  ({error, style, ...props}, ref) => {
    return (
      <TextInput
        ref={ref}
        style={[styles.input, error && styles.inputError, style]}
        placeholderTextColor="#999"
        {...props}
      />
    );
  },
);

const styles = StyleSheet.create({
  input: {
    height: sizes.buttons.md.height,
    borderWidth: 1,
    borderColor: '#ddd',
    borderRadius: sizes.borderRadius.md,
    paddingHorizontal: spacing.sm,
    fontSize: typography.sizes.md,
    color: '#000',
    backgroundColor: '#fff',
  },
  inputError: {
    borderColor: '#ef4444',
  },
});
```

### `src/components/form/FormLabel.tsx`

```tsx
import React from 'react';
import {StyleSheet} from 'react-native';
import {Text} from '../typography/Text';
import {spacing} from '../../theme/spacing';

export type FormLabelProps = {
  children: React.ReactNode;
  required?: boolean;
};

export const FormLabel: React.FC<FormLabelProps> = ({children, required}) => {
  return (
    <Text variant="body" weight="medium" style={styles.label}>
      {children}
      {required && <Text style={styles.required}> *</Text>}
    </Text>
  );
};

const styles = StyleSheet.create({
  label: {
    marginBottom: spacing.xs,
  },
  required: {
    color: '#ef4444',
  },
});
```

### `src/components/form/FormError.tsx`

```tsx
import React from 'react';
import {StyleSheet} from 'react-native';
import {Text} from '../typography/Text';
import {spacing} from '../../theme/spacing';

export type FormErrorProps = {
  message?: string;
};

export const FormError: React.FC<FormErrorProps> = ({message}) => {
  if (!message) return null;

  return (
    <Text variant="caption" style={styles.error}>
      {message}
    </Text>
  );
};

const styles = StyleSheet.create({
  error: {
    color: '#ef4444',
    marginTop: spacing.xs,
  },
});
```

### `src/components/form/FormItem.tsx`

Container for a single form field (label + input + error).

```tsx
import React from 'react';
import {StyleSheet, View, ViewStyle} from 'react-native';
import {spacing} from '../../theme/spacing';

export type FormItemProps = {
  children: React.ReactNode;
  style?: ViewStyle;
};

export const FormItem: React.FC<FormItemProps> = ({children, style}) => {
  return <View style={[styles.item, style]}>{children}</View>;
};

const styles = StyleSheet.create({
  item: {
    marginBottom: spacing.md,
  },
});
```

---

## 2. Zod Schema Patterns for Mobile

### Common Field Validations

```ts
import {z} from 'zod';

// Email
export const emailSchema = z
  .string()
  .min(1, 'Email is required')
  .email('Invalid email address');

// Password (min 8 chars, optional complexity)
export const passwordSchema = z
  .string()
  .min(8, 'Password must be at least 8 characters')
  .max(64, 'Password is too long');

// Name (letters and spaces only)
export const nameSchema = z
  .string()
  .min(1, 'Name is required')
  .regex(/^[A-Za-z\s]+$/, 'Only letters and spaces allowed');

// Phone number (basic pattern, adjust for locale)
export const phoneSchema = z
  .string()
  .min(1, 'Phone number is required')
  .regex(/^\+?[0-9]{10,15}$/, 'Invalid phone number');

// OTP (6 digits)
export const otpSchema = z
  .string()
  .length(6, 'OTP must be exactly 6 digits')
  .regex(/^\d{6}$/, 'OTP must be numeric');

// Numbers (coerce from string inputs)
export const numberSchema = z.coerce
  .number({invalid_type_error: 'Must be a number'})
  .min(1, 'Value must be greater than 0');

// URL
export const urlSchema = z
  .string()
  .url('Invalid URL')
  .or(z.literal('').transform(() => undefined));

// Required string
export const requiredStringSchema = z
  .string()
  .min(1, 'This field is required')
  .trim();
```

### Cross-Field Validation (Password Confirmation)

```ts
export const signupSchema = z
  .object({
    email: emailSchema,
    password: passwordSchema,
    confirmPassword: z.string().min(1, 'Please confirm your password'),
  })
  .refine(data => data.password === data.confirmPassword, {
    path: ['confirmPassword'],
    message: 'Passwords do not match',
  });
```

### Arrays of Objects (Dynamic Fields)

```ts
export const productFormSchema = z.object({
  name: requiredStringSchema,
  price: numberSchema,
  features: z
    .array(
      z.object({
        name: requiredStringSchema,
      }),
    )
    .min(1, 'At least one feature is required'),
});
```

### Optional Fields with Defaults

```ts
export const profileSchema = z.object({
  firstName: requiredStringSchema,
  lastName: requiredStringSchema,
  bio: z.string().optional(),
  notifications: z.boolean().default(true),
});
```

---

## 3. Basic Form Setup

### Pattern: Simple Login Form

```tsx
import React from 'react';
import {
  StyleSheet,
  View,
  TouchableOpacity,
  KeyboardAvoidingView,
  Platform,
} from 'react-native';
import {useForm} from 'react-hook-form';
import {zodResolver} from '@hookform/resolvers/zod';
import {z} from 'zod';
import {Screen} from '../../components/layout/Screen';
import {Content} from '../../components/layout/Content';
import {Text} from '../../components/typography/Text';
import {FormField} from '../../components/form/FormField';
import {FormItem} from '../../components/form/FormItem';
import {FormLabel} from '../../components/form/FormLabel';
import {FormInput} from '../../components/form/FormInput';
import {FormError} from '../../components/form/FormError';
import {spacing} from '../../theme/spacing';
import {sizes} from '../../theme/sizes';

const loginSchema = z.object({
  email: z.string().min(1, 'Email is required').email('Invalid email'),
  password: z.string().min(8, 'Password must be at least 8 characters'),
});

type LoginFormValues = z.infer<typeof loginSchema>;

export const LoginScreen: React.FC = () => {
  const {
    control,
    handleSubmit,
    formState: {errors, isSubmitting},
  } = useForm<LoginFormValues>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: '',
      password: '',
    },
  });

  const onSubmit = async (data: LoginFormValues) => {
    console.log('Login data:', data);
    // Call your login mutation here
  };

  return (
    <Screen>
      <KeyboardAvoidingView
        behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
        style={{flex: 1}}>
        <Content>
          <View style={styles.header}>
            <Text variant="h1">Sign In</Text>
            <Text variant="body" style={styles.subtitle}>
              Welcome back!
            </Text>
          </View>

          <FormField
            control={control}
            name="email"
            render={({field}) => (
              <FormItem>
                <FormLabel required>Email</FormLabel>
                <FormInput
                  {...field}
                  placeholder="you@example.com"
                  keyboardType="email-address"
                  autoCapitalize="none"
                  autoCorrect={false}
                  error={!!errors.email}
                  onChangeText={field.onChange}
                />
                <FormError message={errors.email?.message} />
              </FormItem>
            )}
          />

          <FormField
            control={control}
            name="password"
            render={({field}) => (
              <FormItem>
                <FormLabel required>Password</FormLabel>
                <FormInput
                  {...field}
                  placeholder="••••••••"
                  secureTextEntry
                  autoCapitalize="none"
                  autoCorrect={false}
                  error={!!errors.password}
                  onChangeText={field.onChange}
                />
                <FormError message={errors.password?.message} />
              </FormItem>
            )}
          />

          <TouchableOpacity
            style={styles.button}
            onPress={handleSubmit(onSubmit)}
            disabled={isSubmitting}>
            <Text variant="button" weight="semibold" style={styles.buttonText}>
              {isSubmitting ? 'Signing in...' : 'Sign In'}
            </Text>
          </TouchableOpacity>
        </Content>
      </KeyboardAvoidingView>
    </Screen>
  );
};

const styles = StyleSheet.create({
  header: {
    paddingTop: spacing.lg,
    paddingBottom: spacing.xl,
  },
  subtitle: {
    marginTop: spacing.xs,
    color: '#666',
  },
  button: {
    height: sizes.buttons.lg.height,
    backgroundColor: '#111',
    borderRadius: sizes.borderRadius.md,
    justifyContent: 'center',
    alignItems: 'center',
    marginTop: spacing.md,
  },
  buttonText: {
    color: '#fff',
  },
});
```

**Key Mobile Considerations:**

- `KeyboardAvoidingView` with platform-specific `behavior`.
- Appropriate `keyboardType` for each input (email, numeric, phone-pad, etc.).
- `autoCapitalize="none"` and `autoCorrect={false}` for email/password fields.
- `secureTextEntry` for password fields.

---

## 4. Dynamic Fields with `useFieldArray`

For forms with repeating sections (e.g., adding multiple addresses, features, items).

### Example: Add Product with Features

```tsx
import React from 'react';
import {StyleSheet, View, TouchableOpacity, ScrollView} from 'react-native';
import {useForm, useFieldArray} from 'react-hook-form';
import {zodResolver} from '@hookform/resolvers/zod';
import {z} from 'zod';
import {FormField} from '../../components/form/FormField';
import {FormItem} from '../../components/form/FormItem';
import {FormLabel} from '../../components/form/FormLabel';
import {FormInput} from '../../components/form/FormInput';
import {FormError} from '../../components/form/FormError';
import {Text} from '../../components/typography/Text';
import {spacing} from '../../theme/spacing';
import {sizes} from '../../theme/sizes';

const productSchema = z.object({
  name: z.string().min(1, 'Product name is required'),
  price: z.coerce.number().min(1, 'Price must be greater than 0'),
  features: z
    .array(
      z.object({
        name: z.string().min(1, 'Feature name is required'),
      }),
    )
    .min(1, 'At least one feature is required'),
});

type ProductFormValues = z.infer<typeof productSchema>;

export const AddProductForm: React.FC = () => {
  const {
    control,
    handleSubmit,
    formState: {errors},
  } = useForm<ProductFormValues>({
    resolver: zodResolver(productSchema),
    defaultValues: {
      name: '',
      price: 0,
      features: [{name: ''}],
    },
  });

  const {fields, append, remove} = useFieldArray({
    control,
    name: 'features',
  });

  const onSubmit = (data: ProductFormValues) => {
    console.log('Product data:', data);
  };

  return (
    <ScrollView style={styles.container}>
      <FormField
        control={control}
        name="name"
        render={({field}) => (
          <FormItem>
            <FormLabel required>Product Name</FormLabel>
            <FormInput
              {...field}
              placeholder="Enter product name"
              error={!!errors.name}
              onChangeText={field.onChange}
            />
            <FormError message={errors.name?.message} />
          </FormItem>
        )}
      />

      <FormField
        control={control}
        name="price"
        render={({field}) => (
          <FormItem>
            <FormLabel required>Price</FormLabel>
            <FormInput
              {...field}
              placeholder="0.00"
              keyboardType="decimal-pad"
              error={!!errors.price}
              onChangeText={field.onChange}
            />
            <FormError message={errors.price?.message} />
          </FormItem>
        )}
      />

      <View style={styles.sectionHeader}>
        <Text variant="body" weight="semibold">
          Features
        </Text>
      </View>

      {fields.map((item, index) => (
        <View key={item.id} style={styles.featureRow}>
          <View style={{flex: 1}}>
            <FormField
              control={control}
              name={`features.${index}.name`}
              render={({field}) => (
                <FormItem>
                  <FormInput
                    {...field}
                    placeholder="Feature name"
                    error={!!errors.features?.[index]?.name}
                    onChangeText={field.onChange}
                  />
                  <FormError
                    message={errors.features?.[index]?.name?.message}
                  />
                </FormItem>
              )}
            />
          </View>
          {fields.length > 1 && (
            <TouchableOpacity
              onPress={() => remove(index)}
              style={styles.removeButton}>
              <Text variant="button" style={styles.removeButtonText}>
                Remove
              </Text>
            </TouchableOpacity>
          )}
        </View>
      ))}

      <TouchableOpacity
        style={styles.addButton}
        onPress={() => append({name: ''})}>
        <Text variant="button" weight="medium">
          + Add Feature
        </Text>
      </TouchableOpacity>

      <TouchableOpacity
        style={styles.submitButton}
        onPress={handleSubmit(onSubmit)}>
        <Text
          variant="button"
          weight="semibold"
          style={styles.submitButtonText}>
          Save Product
        </Text>
      </TouchableOpacity>
    </ScrollView>
  );
};

const styles = StyleSheet.create({
  container: {
    padding: spacing.md,
  },
  sectionHeader: {
    marginTop: spacing.lg,
    marginBottom: spacing.sm,
  },
  featureRow: {
    flexDirection: 'row',
    alignItems: 'flex-start',
    gap: spacing.sm,
  },
  removeButton: {
    marginTop: spacing.sm,
    paddingVertical: spacing.xs,
    paddingHorizontal: spacing.sm,
  },
  removeButtonText: {
    color: '#ef4444',
  },
  addButton: {
    paddingVertical: spacing.sm,
    alignItems: 'center',
    borderWidth: 1,
    borderColor: '#ddd',
    borderRadius: sizes.borderRadius.md,
    marginTop: spacing.sm,
  },
  submitButton: {
    height: sizes.buttons.lg.height,
    backgroundColor: '#111',
    borderRadius: sizes.borderRadius.md,
    justifyContent: 'center',
    alignItems: 'center',
    marginTop: spacing.xl,
  },
  submitButtonText: {
    color: '#fff',
  },
});
```

---

## 5. Mobile-Specific Input Patterns

### Numeric Input (Decimal)

```tsx
<FormField
  control={control}
  name="amount"
  render={({field}) => (
    <FormItem>
      <FormLabel>Amount</FormLabel>
      <FormInput
        {...field}
        placeholder="0.00"
        keyboardType="decimal-pad"
        onChangeText={text => {
          // Remove non-numeric chars except dot
          const cleaned = text.replace(/[^0-9.]/g, '');
          field.onChange(cleaned);
        }}
      />
    </FormItem>
  )}
/>
```

Schema: `z.coerce.number().min(0.01, 'Amount must be greater than 0')`

### Phone Number Input

```tsx
<FormField
  control={control}
  name="phone"
  render={({field}) => (
    <FormItem>
      <FormLabel>Phone</FormLabel>
      <FormInput
        {...field}
        placeholder="+1 (555) 123-4567"
        keyboardType="phone-pad"
        autoCapitalize="none"
        onChangeText={field.onChange}
      />
    </FormItem>
  )}
/>
```

Schema: `z.string().regex(/^\+?[0-9]{10,15}$/, 'Invalid phone number')`

### OTP Input (6 digits)

```tsx
<FormField
  control={control}
  name="otp"
  render={({field}) => (
    <FormItem>
      <FormLabel>OTP</FormLabel>
      <FormInput
        {...field}
        placeholder="123456"
        keyboardType="number-pad"
        maxLength={6}
        onChangeText={text => {
          const cleaned = text.replace(/\D/g, '');
          field.onChange(cleaned);
        }}
      />
    </FormItem>
  )}
/>
```

Schema: `z.string().length(6, 'OTP must be 6 digits').regex(/^\d{6}$/, 'OTP must be numeric')`

### Checkbox/Switch

```tsx
import {Switch} from 'react-native';

<FormField
  control={control}
  name="acceptTerms"
  render={({field}) => (
    <View style={styles.checkboxRow}>
      <Switch value={field.value} onValueChange={field.onChange} />
      <Text variant="body">I accept the terms and conditions</Text>
    </View>
  )}
/>;
```

Schema: `z.boolean().refine(val => val === true, 'You must accept the terms')`

### Picker/Select (React Native Picker or custom)

```tsx
import {Picker} from '@react-native-picker/picker';

<FormField
  control={control}
  name="category"
  render={({field}) => (
    <FormItem>
      <FormLabel>Category</FormLabel>
      <Picker selectedValue={field.value} onValueChange={field.onChange}>
        <Picker.Item label="Select category..." value="" />
        <Picker.Item label="Electronics" value="electronics" />
        <Picker.Item label="Clothing" value="clothing" />
        <Picker.Item label="Food" value="food" />
      </Picker>
    </FormItem>
  )}
/>;
```

Schema: `z.string().min(1, 'Category is required')`

---

## 6. Validation Helpers Module

Centralize common regexes and validators in one file.

### `src/utils/validation.ts`

```ts
import {z} from 'zod';

// Email regex (more permissive than RFC 5322, but practical)
export const EMAIL_REGEX = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

// Name validation (letters, spaces, hyphens, apostrophes)
export const NAME_REGEX = /^[A-Za-z\s'-]+$/;

// No < or > (basic XSS prevention)
export const NO_ANGLE_BRACKETS_REGEX = /^[^<>]*$/;

// Phone (E.164 format: +[country][number])
export const PHONE_REGEX = /^\+?[1-9]\d{1,14}$/;

// OTP (6 digits)
export const OTP_REGEX = /^\d{6}$/;

// Common schemas
export const emailField = () =>
  z.string().min(1, 'Email is required').regex(EMAIL_REGEX, 'Invalid email');

export const passwordField = (minLength = 8) =>
  z
    .string()
    .min(minLength, `Password must be at least ${minLength} characters`);

export const nameField = (label = 'Name') =>
  z
    .string()
    .min(1, `${label} is required`)
    .regex(
      NAME_REGEX,
      'Only letters, spaces, hyphens, and apostrophes allowed',
    );

export const phoneField = () =>
  z
    .string()
    .min(1, 'Phone is required')
    .regex(PHONE_REGEX, 'Invalid phone number');

export const otpField = () =>
  z
    .string()
    .length(6, 'OTP must be 6 digits')
    .regex(OTP_REGEX, 'OTP must be numeric');

export const safeTextField = (label = 'Field') =>
  z
    .string()
    .min(1, `${label} is required`)
    .regex(NO_ANGLE_BRACKETS_REGEX, 'Invalid characters detected');
```

**Usage:**

```ts
import {emailField, passwordField, nameField} from '../utils/validation';

const signupSchema = z.object({
  firstName: nameField('First name'),
  lastName: nameField('Last name'),
  email: emailField(),
  password: passwordField(8),
});
```

---

## 7. Error Handling Patterns

### Server-Side Errors

Set field-specific or root errors from API responses.

```tsx
const onSubmit = async (data: LoginFormValues) => {
  try {
    await loginMutation.mutateAsync(data);
  } catch (error: any) {
    if (error.response?.data?.field) {
      // Field-specific error
      setError(error.response.data.field, {
        message: error.response.data.message,
      });
    } else {
      // Root-level error
      setError('root', {
        message:
          error.response?.data?.message || 'Login failed. Please try again.',
      });
    }
  }
};

// Display root error
{
  errors.root && <Text style={styles.rootError}>{errors.root.message}</Text>;
}
```

### Sanitization Before Submit

Trim whitespace and sanitize inputs before sending to API.

```tsx
import {sanitizeInput} from '../utils/sanitize'; // simple trim + XSS cleanup

const onSubmit = (data: FormValues) => {
  const sanitized = {
    email: sanitizeInput(data.email),
    name: sanitizeInput(data.name),
  };

  apiMutation.mutate(sanitized);
};
```

---

## 8. Mobile Keyboard Management

### KeyboardAvoidingView Pattern

```tsx
import {KeyboardAvoidingView, Platform, ScrollView} from 'react-native';

<KeyboardAvoidingView
  behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
  style={{flex: 1}}>
  <ScrollView
    keyboardShouldPersistTaps="handled"
    contentContainerStyle={{paddingBottom: 100}}>
    {/* Form fields */}
  </ScrollView>
</KeyboardAvoidingView>;
```

### Dismiss Keyboard on Submit

```tsx
import {Keyboard} from 'react-native';

const onSubmit = (data: FormValues) => {
  Keyboard.dismiss();
  // Handle submit
};
```

### Auto-Focus Next Field

Use `returnKeyType` and `onSubmitEditing` to navigate between fields.

```tsx
const emailRef = useRef<TextInput>(null);
const passwordRef = useRef<TextInput>(null);

<FormInput
  ref={emailRef}
  returnKeyType="next"
  onSubmitEditing={() => passwordRef.current?.focus()}
  {...field}
/>

<FormInput
  ref={passwordRef}
  returnKeyType="done"
  onSubmitEditing={handleSubmit(onSubmit)}
  {...field}
/>
```

---

## 9. Complete Example: Registration Form

```tsx
import React, {useRef} from 'react';
import {
  StyleSheet,
  View,
  TouchableOpacity,
  KeyboardAvoidingView,
  Platform,
  ScrollView,
  TextInput,
  Keyboard,
} from 'react-native';
import {useForm} from 'react-hook-form';
import {zodResolver} from '@hookform/resolvers/zod';
import {z} from 'zod';
import {Screen} from '../../components/layout/Screen';
import {Content} from '../../components/layout/Content';
import {Text} from '../../components/typography/Text';
import {FormField} from '../../components/form/FormField';
import {FormItem} from '../../components/form/FormItem';
import {FormLabel} from '../../components/form/FormLabel';
import {FormInput} from '../../components/form/FormInput';
import {FormError} from '../../components/form/FormError';
import {spacing} from '../../theme/spacing';
import {sizes} from '../../theme/sizes';

const registerSchema = z
  .object({
    firstName: z
      .string()
      .min(1, 'First name is required')
      .regex(/^[A-Za-z\s]+$/, 'Only letters and spaces'),
    lastName: z
      .string()
      .min(1, 'Last name is required')
      .regex(/^[A-Za-z\s]+$/, 'Only letters and spaces'),
    email: z.string().min(1, 'Email is required').email('Invalid email'),
    phone: z
      .string()
      .min(1, 'Phone is required')
      .regex(/^\+?[0-9]{10,15}$/, 'Invalid phone number'),
    password: z.string().min(8, 'Password must be at least 8 characters'),
    confirmPassword: z.string().min(1, 'Please confirm your password'),
  })
  .refine(data => data.password === data.confirmPassword, {
    path: ['confirmPassword'],
    message: 'Passwords do not match',
  });

type RegisterFormValues = z.infer<typeof registerSchema>;

export const RegisterScreen: React.FC = () => {
  const lastNameRef = useRef<TextInput>(null);
  const emailRef = useRef<TextInput>(null);
  const phoneRef = useRef<TextInput>(null);
  const passwordRef = useRef<TextInput>(null);
  const confirmPasswordRef = useRef<TextInput>(null);

  const {
    control,
    handleSubmit,
    formState: {errors, isSubmitting},
  } = useForm<RegisterFormValues>({
    resolver: zodResolver(registerSchema),
    defaultValues: {
      firstName: '',
      lastName: '',
      email: '',
      phone: '',
      password: '',
      confirmPassword: '',
    },
  });

  const onSubmit = async (data: RegisterFormValues) => {
    Keyboard.dismiss();
    console.log('Register data:', data);
    // Call registration mutation
  };

  return (
    <Screen>
      <KeyboardAvoidingView
        behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
        style={{flex: 1}}>
        <ScrollView
          keyboardShouldPersistTaps="handled"
          contentContainerStyle={{paddingBottom: 50}}>
          <Content>
            <View style={styles.header}>
              <Text variant="h1">Create Account</Text>
              <Text variant="body" style={styles.subtitle}>
                Sign up to get started
              </Text>
            </View>

            <FormField
              control={control}
              name="firstName"
              render={({field}) => (
                <FormItem>
                  <FormLabel required>First Name</FormLabel>
                  <FormInput
                    {...field}
                    placeholder="John"
                    returnKeyType="next"
                    onSubmitEditing={() => lastNameRef.current?.focus()}
                    error={!!errors.firstName}
                    onChangeText={field.onChange}
                  />
                  <FormError message={errors.firstName?.message} />
                </FormItem>
              )}
            />

            <FormField
              control={control}
              name="lastName"
              render={({field}) => (
                <FormItem>
                  <FormLabel required>Last Name</FormLabel>
                  <FormInput
                    {...field}
                    ref={lastNameRef}
                    placeholder="Doe"
                    returnKeyType="next"
                    onSubmitEditing={() => emailRef.current?.focus()}
                    error={!!errors.lastName}
                    onChangeText={field.onChange}
                  />
                  <FormError message={errors.lastName?.message} />
                </FormItem>
              )}
            />

            <FormField
              control={control}
              name="email"
              render={({field}) => (
                <FormItem>
                  <FormLabel required>Email</FormLabel>
                  <FormInput
                    {...field}
                    ref={emailRef}
                    placeholder="you@example.com"
                    keyboardType="email-address"
                    autoCapitalize="none"
                    autoCorrect={false}
                    returnKeyType="next"
                    onSubmitEditing={() => phoneRef.current?.focus()}
                    error={!!errors.email}
                    onChangeText={field.onChange}
                  />
                  <FormError message={errors.email?.message} />
                </FormItem>
              )}
            />

            <FormField
              control={control}
              name="phone"
              render={({field}) => (
                <FormItem>
                  <FormLabel required>Phone</FormLabel>
                  <FormInput
                    {...field}
                    ref={phoneRef}
                    placeholder="+1234567890"
                    keyboardType="phone-pad"
                    returnKeyType="next"
                    onSubmitEditing={() => passwordRef.current?.focus()}
                    error={!!errors.phone}
                    onChangeText={field.onChange}
                  />
                  <FormError message={errors.phone?.message} />
                </FormItem>
              )}
            />

            <FormField
              control={control}
              name="password"
              render={({field}) => (
                <FormItem>
                  <FormLabel required>Password</FormLabel>
                  <FormInput
                    {...field}
                    ref={passwordRef}
                    placeholder="••••••••"
                    secureTextEntry
                    autoCapitalize="none"
                    autoCorrect={false}
                    returnKeyType="next"
                    onSubmitEditing={() => confirmPasswordRef.current?.focus()}
                    error={!!errors.password}
                    onChangeText={field.onChange}
                  />
                  <FormError message={errors.password?.message} />
                </FormItem>
              )}
            />

            <FormField
              control={control}
              name="confirmPassword"
              render={({field}) => (
                <FormItem>
                  <FormLabel required>Confirm Password</FormLabel>
                  <FormInput
                    {...field}
                    ref={confirmPasswordRef}
                    placeholder="••••••••"
                    secureTextEntry
                    autoCapitalize="none"
                    autoCorrect={false}
                    returnKeyType="done"
                    onSubmitEditing={handleSubmit(onSubmit)}
                    error={!!errors.confirmPassword}
                    onChangeText={field.onChange}
                  />
                  <FormError message={errors.confirmPassword?.message} />
                </FormItem>
              )}
            />

            <TouchableOpacity
              style={styles.button}
              onPress={handleSubmit(onSubmit)}
              disabled={isSubmitting}>
              <Text
                variant="button"
                weight="semibold"
                style={styles.buttonText}>
                {isSubmitting ? 'Creating account...' : 'Sign Up'}
              </Text>
            </TouchableOpacity>
          </Content>
        </ScrollView>
      </KeyboardAvoidingView>
    </Screen>
  );
};

const styles = StyleSheet.create({
  header: {
    paddingTop: spacing.lg,
    paddingBottom: spacing.xl,
  },
  subtitle: {
    marginTop: spacing.xs,
    color: '#666',
  },
  button: {
    height: sizes.buttons.lg.height,
    backgroundColor: '#111',
    borderRadius: sizes.borderRadius.md,
    justifyContent: 'center',
    alignItems: 'center',
    marginTop: spacing.lg,
  },
  buttonText: {
    color: '#fff',
  },
});
```

---

## 10. Best Practices Summary

### Schema Design

- Use `z.infer<typeof schema>` for type inference (single source of truth).
- Provide user-friendly error messages in validation rules.
- Centralize common patterns in `validation.ts`.
- Use `z.coerce.number()` for numeric fields from text inputs.
- Use `.refine()` for cross-field validation.
- Use `.transform()` for data normalization (e.g., trim, lowercase).

### Form Setup

- Always use `zodResolver(schema)` in `useForm`.
- Provide `defaultValues` for all fields (prevents uncontrolled → controlled warnings).
- Use `Controller` via `FormField` wrapper for consistent rendering.

### Mobile-Specific

- Set appropriate `keyboardType` for each input type.
- Use `autoCapitalize="none"` and `autoCorrect={false}` for emails/passwords.
- Wrap forms in `KeyboardAvoidingView` + `ScrollView`.
- Use `returnKeyType` and `onSubmitEditing` for field navigation.
- Call `Keyboard.dismiss()` on submit.

### Error Handling

- Display field errors via `FormError` component.
- Use `setError('root', {message})` for server/form-level errors.
- Show loading state (`isSubmitting`) to disable submit button.

### Reusability

- Build form components (`FormField`, `FormInput`, `FormLabel`, `FormError`) once.
- Create a `validation.ts` module for shared regex/schema patterns.
- Keep form logic in custom hooks when forms are complex (e.g., `useLoginForm`).

---

## 11. File Structure Template

```
src/
  components/
    form/
      FormField.tsx          # Controller wrapper
      FormInput.tsx          # Styled TextInput
      FormLabel.tsx          # Label with required indicator
      FormError.tsx          # Error message display
      FormItem.tsx           # Container for field group
    layout/
      Screen.tsx
      Content.tsx
    typography/
      Text.tsx
  utils/
    validation.ts            # Centralized regex & schema builders
    sanitize.ts              # Input sanitization helpers
  features/
    auth/
      screens/
        LoginScreen.tsx      # Login form
        RegisterScreen.tsx   # Registration form
      schemas/
        authSchemas.ts       # Auth-related Zod schemas
    products/
      screens/
        AddProductScreen.tsx
      schemas/
        productSchemas.ts
```

---

## TL;DR

- **Schema-first**: Define Zod schemas with `z.infer` for TypeScript types.
- **Reusable components**: Build `FormField`, `FormInput`, `FormLabel`, `FormError` once; use everywhere.
- **Mobile inputs**: Set correct `keyboardType`, `autoCapitalize`, `secureTextEntry`, `returnKeyType`.
- **Keyboard handling**: Use `KeyboardAvoidingView` + `ScrollView` + `Keyboard.dismiss()`.
- **Validation**: Centralize common patterns in `validation.ts`; use clear error messages.
- **Dynamic fields**: Use `useFieldArray` for arrays of objects.
- **Error handling**: Use `setError` for server errors; display with `FormError`.

---

## Next Steps

1. Install dependencies:

   ```bash
   npm install react-hook-form zod @hookform/resolvers
   ```

2. Create form components in `src/components/form/`.

3. Create `src/utils/validation.ts` with common schemas.

4. Build your first form (e.g., login) following the patterns above.

5. Integrate with React Query mutations for API calls.

6. Test on real devices to tune keyboard behavior and input types.

7. Refer to this guide when building new forms or refactoring existing ones.
