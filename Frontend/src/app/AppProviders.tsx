import React from 'react';
import {QueryClientProvider} from '@tanstack/react-query';
import {SafeAreaProvider} from 'react-native-safe-area-context';
import {GestureHandlerRootView} from 'react-native-gesture-handler';
import {StyleSheet} from 'react-native';
import {queryClient} from '../services/http/queryClient';

type AppProvidersProps = {
  children: React.ReactNode;
};

/**
 * Root providers wrapper for the application
 * Includes:
 * - GestureHandlerRootView (required for react-native-gesture-handler)
 * - SafeAreaProvider (safe area context)
 * - QueryClientProvider (React Query)
 */
export const AppProviders: React.FC<AppProvidersProps> = ({children}) => {
  return (
    <GestureHandlerRootView style={styles.container}>
      <SafeAreaProvider>
        <QueryClientProvider client={queryClient}>
          {children}
        </QueryClientProvider>
      </SafeAreaProvider>
    </GestureHandlerRootView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
});

export default AppProviders;
