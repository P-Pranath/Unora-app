# Complex Screen Architecture Guide

This guide outlines the patterns and practices for building complex screens in the Cenomi app. It follows SOLID principles, clean architecture, and React best practices.

## Overview

When a screen grows beyond ~300 lines or handles multiple concerns (form state, API calls, navigation, UI state), it should be refactored into a clean architecture with separated hooks.

### Before (Anti-pattern)

```
screens/
└── LoginScreen.tsx  # 700+ lines with mixed concerns
```

### After (Clean Architecture)

```
features/auth/
├── hooks/
│   ├── index.ts              # Barrel exports
│   ├── useLogin.ts           # Orchestrator (facade)
│   ├── useLoginAuth.ts       # Auth logic
│   ├── useLoginForm.ts       # Form state
│   ├── useLoginError.ts      # Error handling
│   └── useLoginBottomSheets.ts  # UI state
├── screens/
│   └── LoginScreen/
│       ├── index.tsx         # Presentational only
│       └── styles.ts         # Extracted styles
```

## Core Principles

### 1. Single Responsibility Principle (SRP)

Each hook should handle ONE concern:

- **Form Hook**: Form state, validation, input handlers
- **Auth Hook**: API calls, authentication flows, session management
- **Error Hook**: Error popup state, show/close handlers
- **UI State Hook**: Bottom sheets, modals, toggles

### 2. Separation of Concerns

- **Hooks**: Business logic, state management, side effects
- **Screen Component**: Pure UI rendering based on hook state
- **Styles**: Extracted to separate file

### 3. Facade Pattern

The orchestrator hook (`useLogin`) combines all sub-hooks and exposes a clean, organized API to the screen component.

## Step-by-Step Guide

### Step 1: Identify Concerns

Before refactoring, identify the different concerns in your screen:

| Concern            | Example                          | Hook Name                              |
| ------------------ | -------------------------------- | -------------------------------------- |
| Form state         | Input values, validation, submit | `use{Screen}Form`                      |
| API/Business logic | Fetch, submit, auth flows        | `use{Screen}Auth` or `use{Screen}Data` |
| Error handling     | Error popups, messages           | `use{Screen}Error`                     |
| UI state           | Modals, bottom sheets, toggles   | `use{Screen}UI` or specific hooks      |
| Navigation         | Route changes, deep links        | Usually in auth/data hook              |

### Step 2: Create the Hook Structure

```
src/features/{feature}/
├── hooks/
│   ├── index.ts
│   ├── use{Screen}.ts           # Orchestrator
│   ├── use{Screen}Form.ts       # Form logic
│   ├── use{Screen}Auth.ts       # Business logic
│   ├── use{Screen}Error.ts      # Error state
│   └── use{Screen}UI.ts         # UI state (optional)
```

### Step 3: Implement Sub-Hooks

#### Error Hook (Simplest - start here)

```typescript
// hooks/useLoginError.ts
import {useState, useCallback} from 'react';
import type {AuthError} from '../types';

export interface UseLoginErrorReturn {
  showErrorPopup: boolean;
  errorInfo: AuthError | null;
  showError: (error: AuthError) => void;
  closeErrorPopup: () => void;
}

export const useLoginError = (): UseLoginErrorReturn => {
  const [showErrorPopup, setShowErrorPopup] = useState(false);
  const [errorInfo, setErrorInfo] = useState<AuthError | null>(null);

  const showError = useCallback((error: AuthError) => {
    setErrorInfo(error);
    setShowErrorPopup(true);
  }, []);

  const closeErrorPopup = useCallback(() => {
    setShowErrorPopup(false);
    setErrorInfo(null);
  }, []);

  return {
    showErrorPopup,
    errorInfo,
    showError,
    closeErrorPopup,
  };
};
```

#### Form Hook

```typescript
// hooks/useLoginForm.ts
import {useCallback} from 'react';
import {useForm, Control, UseFormHandleSubmit} from 'react-hook-form';
import {zodResolver} from '@hookform/resolvers/zod';
import {
  loginSchema,
  LoginFormValues,
  defaultLoginFormValues,
} from '../schemas/loginSchema';

export interface UseLoginFormReturn {
  control: Control<LoginFormValues>;
  handleSubmit: UseFormHandleSubmit<LoginFormValues>;
  username: string;
  password: string;
  handleUsernameChange: (value: string) => void;
  handlePasswordChange: (value: string) => void;
  resetPassword: () => void;
}

export const useLoginForm = (): UseLoginFormReturn => {
  const {control, handleSubmit, watch, setValue} = useForm<LoginFormValues>({
    resolver: zodResolver(loginSchema),
    defaultValues: defaultLoginFormValues,
    mode: 'onBlur',
  });

  const username = watch('username');
  const password = watch('password');

  const handleUsernameChange = useCallback(
    (value: string) => {
      const cleanValue = value.replace(/\s/g, '');
      setValue('username', cleanValue);
    },
    [setValue],
  );

  const handlePasswordChange = useCallback(
    (value: string) => {
      setValue('password', value);
    },
    [setValue],
  );

  const resetPassword = useCallback(() => {
    setValue('password', '');
  }, [setValue]);

  return {
    control,
    handleSubmit,
    username,
    password,
    handleUsernameChange,
    handlePasswordChange,
    resetPassword,
  };
};
```

#### Auth/Business Logic Hook

```typescript
// hooks/useLoginAuth.ts
import {useState, useCallback, useEffect} from 'react';
import {useNavigation} from '@react-navigation/native';
import {useAuthStore} from '../store/authStore';
import {checkUser, login} from '../api/authApi';
import type {AuthError} from '../types';

export interface UseLoginAuthProps {
  showError: (error: AuthError) => void; // Injected dependency
}

export interface UseLoginAuthReturn {
  isLoading: boolean;
  showPasswordField: boolean;
  handleCheckUser: (username: string) => Promise<void>;
  handleEmailLogin: (data: LoginFormValues) => Promise<void>;
  handleGuestLogin: () => void;
  resetAuthState: () => void;
}

export const useLoginAuth = ({
  showError,
}: UseLoginAuthProps): UseLoginAuthReturn => {
  const navigation = useNavigation();
  const {setLoading, setSession, ...store} = useAuthStore();

  const [showPasswordField, setShowPasswordField] = useState(false);

  // ... implementation

  return {
    isLoading: store.isLoading,
    showPasswordField,
    handleCheckUser,
    handleEmailLogin,
    handleGuestLogin,
    resetAuthState,
  };
};
```

### Step 4: Create Orchestrator Hook (Facade)

```typescript
// hooks/useLogin.ts
import {useCallback, useState} from 'react';
import {Keyboard} from 'react-native';
import {useAuthStore} from '../store/authStore';
import {useLoginError} from './useLoginError';
import {useLoginForm} from './useLoginForm';
import {useLoginAuth} from './useLoginAuth';

export interface UseLoginReturn {
  form: {
    control: ReturnType<typeof useLoginForm>['control'];
    username: string;
    handleUsernameChange: (value: string) => void;
    // ... other form props
  };
  auth: {
    isLoading: boolean;
    showPasswordField: boolean;
    handleContinue: () => void;
    // ... other auth props
  };
  error: {
    showErrorPopup: boolean;
    errorInfo: ReturnType<typeof useLoginError>['errorInfo'];
    closeErrorPopup: () => void;
  };
  // ... other grouped props
}

export const useLogin = (): UseLoginReturn => {
  // Initialize sub-hooks
  const errorHook = useLoginError();
  const formHook = useLoginForm();
  const authHook = useLoginAuth({showError: errorHook.showError});

  // Coordinate between hooks
  const handleUsernameChange = useCallback(
    (value: string) => {
      formHook.handleUsernameChange(value);
      formHook.resetPassword();
      authHook.resetAuthState();
    },
    [formHook, authHook],
  );

  const handleContinue = useCallback(() => {
    Keyboard.dismiss();
    formHook.handleSubmit(data => {
      if (authHook.showPasswordField) {
        authHook.handleEmailLogin(data);
      } else {
        authHook.handleCheckUser(data.username);
      }
    })();
  }, [formHook, authHook]);

  // Return organized API
  return {
    form: {
      control: formHook.control,
      username: formHook.username,
      handleUsernameChange,
      // ...
    },
    auth: {
      isLoading: authHook.isLoading,
      showPasswordField: authHook.showPasswordField,
      handleContinue,
      // ...
    },
    error: {
      showErrorPopup: errorHook.showErrorPopup,
      errorInfo: errorHook.errorInfo,
      closeErrorPopup: errorHook.closeErrorPopup,
    },
  };
};
```

### Step 5: Create Presentational Screen

```typescript
// screens/LoginScreen/index.tsx
import React from 'react';
import {View, KeyboardAvoidingView} from 'react-native';
import {CustomText, CustomTextInput, AppPrimaryButton} from '@components/UI';
import {FormField} from '@components/form';
import {useColors} from '@theme/colors';
import {useT} from '@services/i18n';
import {useLogin} from '../../hooks/useLogin';
import {styles} from './styles';

export const LoginScreen: React.FC = () => {
  const colors = useColors();
  const t = useT();
  const {form, auth, error, bottomSheets} = useLogin();

  return (
    <View style={styles.container}>
      {/* Pure UI rendering - no logic */}
      <FormField
        control={form.control}
        name="username"
        render={({field}) => (
          <CustomTextInput
            value={field.value}
            onChangeText={form.handleUsernameChange}
          />
        )}
      />

      <AppPrimaryButton
        onPress={auth.handleContinue}
        isLoading={auth.isLoading}
      />

      <ErrorPopup
        isVisible={error.showErrorPopup}
        onClose={error.closeErrorPopup}
      />
    </View>
  );
};
```

### Step 6: Extract Styles

```typescript
// screens/LoginScreen/styles.ts
import {StyleSheet} from 'react-native';
import {spacing} from '@theme/spacing';
import {moderateScale} from '@utils/responsive';

export const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: spacing.lg,
  },
  // ... all styles
});
```

### Step 7: Create Barrel Export

```typescript
// hooks/index.ts
export {useLogin} from './useLogin';
export type {UseLoginReturn} from './useLogin';

export {useLoginForm} from './useLoginForm';
export {useLoginAuth} from './useLoginAuth';
export {useLoginError} from './useLoginError';
// ... types
```

## Best Practices

### DO ✅

- Keep the screen component under 250 lines
- Use TypeScript interfaces for all hook returns
- Group related state in the orchestrator return object
- Use `useCallback` for all handlers
- Inject dependencies (like `showError`) into hooks that need them
- Add JSDoc comments to exported interfaces
- Create barrel exports for clean imports

### DON'T ❌

- Put business logic in the screen component
- Have hooks directly depend on each other (use orchestrator)
- Mix UI state with business logic in the same hook
- Forget to memoize handlers
- Create circular dependencies between hooks

## When to Apply This Pattern

Apply this architecture when:

- Screen exceeds ~300 lines
- Screen has 3+ distinct concerns (form, API, modals, etc.)
- Multiple developers work on the same screen
- Screen needs extensive unit testing
- Logic will be reused across screens

Keep it simple when:

- Screen is under 200 lines
- Screen has 1-2 simple concerns
- It's a static/display-only screen

## Testing Strategy

```typescript
// hooks/__tests__/useLoginForm.test.ts
import {renderHook, act} from '@testing-library/react-hooks';
import {useLoginForm} from '../useLoginForm';

describe('useLoginForm', () => {
  it('should sanitize username input', () => {
    const {result} = renderHook(() => useLoginForm());

    act(() => {
      result.current.handleUsernameChange('test @email.com');
    });

    expect(result.current.username).toBe('test@email.com');
  });
});
```

## Real Example: LoginScreen

See the implementation at:

- `src/features/auth/hooks/` - All login hooks
- `src/features/auth/screens/LoginScreen/` - Presentational component

This refactoring reduced the LoginScreen from ~770 lines to ~220 lines while improving testability and maintainability.
