# Monitoring & Error Handling Setup

This project includes comprehensive monitoring and error handling based on the [10 checks for stable React Native apps](./error-handling.md).

## 🎯 Implemented Checks

- ✅ **Check #1**: Crash Rate Monitoring (Sentry)
- ✅ **Check #2**: Centralized Error Handling
- ✅ **Check #3**: Versioning and Build Tracking
- ✅ **Check #5**: Contextual Logging
- ✅ **Check #6**: Performance Tracking

## 🚀 Quick Start

### 1. Get Your Sentry DSN

1. Sign up at [sentry.io](https://sentry.io) (free tier available)
2. Create a new React Native project
3. Copy your DSN from Settings → Projects → [Your Project] → Client Keys (DSN)

**Sentry Free Tier includes:**

- 5,000 errors/month
- 10,000 performance units/month
- 1 project
- 30-day data retention

### 2. Configure Sentry DSN

Update `index.js` with your Sentry DSN:

```javascript
const SENTRY_DSN = 'YOUR_SENTRY_DSN_HERE';
```

Or set it as an environment variable:

```bash
export SENTRY_DSN="your-dsn-here"
```

### 3. Build and Run

The monitoring services will initialize automatically when the app starts.

```bash
# iOS
yarn ios

# Android
yarn android
```

## 📦 Services Overview

### Sentry Service

Handles crash monitoring and performance tracking.

```typescript
import {sentryService} from '@/services/monitoring';

// Track user (call after login)
sentryService.setUser('user-123', 'user@example.com');

// Clear user (call on logout)
sentryService.clearUser();

// Add custom context
sentryService.setContext('checkout', {
  cartValue: 149.99,
  itemCount: 3,
});

// Manually capture an error
try {
  await riskyOperation();
} catch (error) {
  sentryService.captureException(error, {
    tags: {section: 'checkout'},
    extra: {userId: user.id},
  });
}
```

### Logger Service

Centralized logging with automatic Sentry integration and data sanitization.

```typescript
import {logger} from '@/services/monitoring';

// Different log levels
logger.debug('Debug info', {detail: 'value'});
logger.info('User logged in', {userId: '123'});
logger.warn('Slow operation detected', {duration: 5000});
logger.error('Operation failed', error, {context: 'data'});

// Specialized logging
logger.logUserAction('Clicked checkout button', {cartTotal: 99.99});
logger.logNavigation('ProfileScreen', {userId: '123'});
logger.logApiCall('POST', '/api/checkout', 200, 450);
logger.logApiError('/api/payment', 500, error, 2000);

// Sensitive data is automatically redacted
logger.info('Login attempt', {
  email: 'user@example.com',
  password: 'secret123', // Will be logged as '***REDACTED***'
});
```

### Error Handler

Global error catching for JavaScript, Native, and Promise errors.

```typescript
import {errorHandler} from '@/services/monitoring';

// Manually report a caught error
try {
  await dangerousOperation();
} catch (error) {
  errorHandler.reportError(error, {
    section: 'PaymentFlow',
    action: 'ProcessPayment',
    extra: {amount: 99.99},
  });
}
```

## 🎨 Components

### AppVersion

Display app version and build number in your UI.

```typescript
import {
  AppVersion,
  AppVersionCompact,
  useAppVersion,
} from '@/services/monitoring';

// Full component with copy button
function AboutScreen() {
  return (
    <View>
      <AppVersion onCopy={() => Alert.alert('Version copied!')} />
    </View>
  );
}

// Compact version for footers
function Footer() {
  return <AppVersionCompact prefix="v" />;
}

// Hook for accessing version in logic
function MyComponent() {
  const {version, buildNumber, displayVersion} = useAppVersion();

  return <Text>App {displayVersion}</Text>;
}
```

## 🔥 Hooks

### useScreenLoadTime

Track screen load performance automatically.

```typescript
import {useScreenLoadTime} from '@/services/monitoring';

function ProfileScreen() {
  // Automatically tracks load time and reports if > 3 seconds
  useScreenLoadTime({
    screenName: 'ProfileScreen',
    slowThreshold: 3000, // ms
  });

  return <View>...</View>;
}
```

### usePerformanceTracker

Track custom operations performance.

```typescript
import {usePerformanceTracker} from '@/services/monitoring';

function CheckoutScreen() {
  const tracker = usePerformanceTracker();

  const handleSubmit = async () => {
    const span = tracker.startSpan('form.submit', 'Submit checkout form');

    try {
      await submitOrder();
      span.finish();
    } catch (error) {
      span.finish();
      throw error;
    }
  };

  return <Button onPress={handleSubmit} />;
}
```

## 📊 Usage Examples

### Complete Screen Example

```typescript
import React, {useEffect} from 'react';
import {View, Text, Button} from 'react-native';
import {
  useScreenLoadTime,
  logger,
  errorHandler,
  AppVersionCompact,
} from '@/services/monitoring';

function ProfileScreen({route}) {
  const {userId} = route.params;

  // Track screen load performance
  useScreenLoadTime({screenName: 'ProfileScreen'});

  useEffect(() => {
    // Log navigation
    logger.logNavigation('ProfileScreen', {userId});

    // Load user data
    loadUserProfile();
  }, [userId]);

  const loadUserProfile = async () => {
    try {
      logger.info('Loading user profile', {userId});
      const profile = await api.getUserProfile(userId);
      logger.info('Profile loaded successfully', {userId});
    } catch (error) {
      logger.error('Failed to load profile', error, {userId});
      errorHandler.reportError(error, {
        section: 'ProfileScreen',
        action: 'LoadProfile',
        extra: {userId},
      });
    }
  };

  return (
    <View>
      <Text>Profile Screen</Text>
      <AppVersionCompact style={{marginTop: 20}} />
    </View>
  );
}
```

### API Client Integration

```typescript
import axios from 'axios';
import {logger} from '@/services/monitoring';

const api = axios.create({
  baseURL: 'https://api.example.com',
});

// Request interceptor
api.interceptors.request.use(
  config => {
    logger.info('API Request', {
      method: config.method,
      url: config.url,
    });
    config.metadata = {startTime: Date.now()};
    return config;
  },
  error => {
    logger.error('API Request Error', error);
    return Promise.reject(error);
  },
);

// Response interceptor
api.interceptors.response.use(
  response => {
    const duration = Date.now() - response.config.metadata.startTime;
    logger.logApiSuccess(response.config.url, duration);
    return response;
  },
  error => {
    const duration = error.config?.metadata
      ? Date.now() - error.config.metadata.startTime
      : undefined;
    logger.logApiError(
      error.config?.url || 'unknown',
      error.response?.status || 0,
      error,
      duration,
    );
    return Promise.reject(error);
  },
);
```

## 🔐 Security

All sensitive data is automatically sanitized from logs:

- `password`
- `token`, `accessToken`, `refreshToken`
- `apiKey`, `secret`
- `ssn`, `creditCard`, `cvv`, `pin`

To add custom sensitive keys:

```typescript
import {logger} from '@/services/monitoring';

logger.addSensitiveKeys(['customSecret', 'internalToken']);
```

## 🎯 Best Practices

1. **Initialize Early**: Monitoring is initialized in `index.js` before app registration
2. **Track User Context**: Call `sentryService.setUser()` after login
3. **Clear on Logout**: Call `sentryService.clearUser()` when user logs out
4. **Add Screen Tracking**: Use `useScreenLoadTime` in all major screens
5. **Log User Actions**: Use `logger.logUserAction()` for important user interactions
6. **Track Navigation**: Use `logger.logNavigation()` or integrate with navigation listeners
7. **API Logging**: Integrate logger with your API client (axios/fetch)
8. **Don't Over-Log**: Use appropriate log levels (debug for development, error for production issues)

## 📈 Monitoring Dashboard

Once configured, you can view:

1. **Sentry Dashboard**: [sentry.io](https://sentry.io)

   - Crash rate and error trends
   - Performance metrics (screen load times)
   - User impact analysis
   - Release comparison

2. **Set Up Alerts** (in Sentry):
   - Crash rate > 0.5%
   - New errors detected
   - Performance degradation
   - Error spike detection

## 🐛 Troubleshooting

### Sentry not tracking errors

1. Check DSN is correctly set in `index.js`
2. Verify network connectivity
3. Check Sentry dashboard for incoming events
4. In development, errors might not be sent (check `beforeSend` in sentry service)

### Performance not showing

1. Ensure `tracesSampleRate` is set (default: 1.0 in dev, 0.2 in production)
2. Check Sentry Performance tab in dashboard
3. Verify transactions are being created

### Build errors

Run pod install for iOS after adding Sentry:

```bash
cd ios && pod install
```

## 📚 Additional Resources

- [Sentry React Native Docs](https://docs.sentry.io/platforms/react-native/)
- [Error Handling Best Practices](./error-handling.md)
- [Sentry Performance Monitoring](https://docs.sentry.io/platforms/react-native/performance/)

## 🎉 What's Next?

Consider implementing additional checks from the [error-handling.md](./error-handling.md) guide:

- Check #4: Feature Flags
- Check #7: CI/CD Automation
- Check #8: Offline State Handling
- Check #9: Security and Sensitive Data
- Check #10: Real-time Alerts

---

**Need Help?** Check the [error-handling.md](./error-handling.md) guide or open an issue.
