module.exports = {
  presets: ['module:@react-native/babel-preset'],
  plugins: [
    // Module resolver for path aliases
    [
      'module-resolver',
      {
        root: ['./src'],
        extensions: ['.ios.js', '.android.js', '.js', '.ts', '.tsx', '.json'],
        alias: {
          '@': './src',
          '@app': './src/app',
          '@assets': './src/assets',
          '@components': './src/components',
          '@constants': './src/constants',
          '@features': './src/features',
          '@hooks': './src/hooks',
          '@navigation': './src/navigation',
          '@services': './src/services',
          '@stores': './src/stores',
          '@theme': './src/theme',
          '@types': './src/types',
          '@utils': './src/utils',
        },
      },
    ],
    // React Native Reanimated - MUST be last
    'react-native-reanimated/plugin',
  ],
};
