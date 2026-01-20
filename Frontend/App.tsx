/**
 * Unora App
 *
 * @format
 */

import React from 'react';
import {StatusBar, StyleSheet, View, Platform} from 'react-native';
import {AppProviders} from './src/app';
import {OnboardingCarousel} from './src/features/onboarding/components';
import {colors} from './src/theme';

function AppContent() {
  const handleOnboardingComplete = () => {
    // TODO: Navigate to main app or mark onboarding as complete
    console.log('Onboarding completed');
  };

  const handleOnboardingSkip = () => {
    // TODO: Handle skip action (e.g., mark as skipped, navigate to main app)
    console.log('Onboarding skipped');
  };

  return (
    <View style={styles.container}>
      <StatusBar
        barStyle={Platform.OS === 'ios' ? 'dark-content' : 'light-content'}
        backgroundColor={colors.onboarding.backgroundLight}
      />
      <OnboardingCarousel
        onComplete={handleOnboardingComplete}
        onSkip={handleOnboardingSkip}
      />
    </View>
  );
}

function App() {
  return (
    <AppProviders>
      <AppContent />
    </AppProviders>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: colors.onboarding.backgroundLight,
  },
});

export default App;
