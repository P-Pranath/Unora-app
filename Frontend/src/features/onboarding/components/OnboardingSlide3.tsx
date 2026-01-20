import React from 'react';
import {View, StyleSheet, TouchableOpacity, Platform} from 'react-native';
import {CustomText} from '@/components/UI/CustomText';
import {ProgressiveRevealIcon} from '@/components/UI/Icons';
import {colors, FontSize} from '@/theme';
import {RfW, RfH} from '@/utils/helpers';
import {HEIGHT, WIDTH, BorderRadius} from '@/theme/sizes';

interface OnboardingSlide3Props {
  onGetStarted?: () => void;
}

export const OnboardingSlide3: React.FC<OnboardingSlide3Props> = ({
  onGetStarted,
}) => {
  return (
    <View style={styles.container}>
      {/* Main Content */}
      <View style={styles.mainContent}>
        {/* Illustration Card */}
        <View style={styles.illustrationCard}>
          <ProgressiveRevealIcon
            width={RfW(220)}
            height={RfW(220)}
            primaryColor={colors.onboarding.primary}
            secondaryColor={colors.onboarding.textBody}
          />
        </View>

        {/* Text Content */}
        <View style={styles.textContent}>
          <CustomText
            fontSize={28}
            fontWeight="bold"
            color={colors.onboarding.textHead}
            fontFamily="display"
            textAlign="center"
            style={styles.title}>
            Earned, not given.
          </CustomText>
          <CustomText
            fontSize={16}
            fontWeight="regular"
            color={colors.onboarding.textBody}
            fontFamily="body"
            textAlign="center"
            style={styles.subtitle}>
            Identity is revealed gradually as trust builds.
          </CustomText>
        </View>
      </View>

      {/* Footer */}
      <View style={styles.footer}>
        {/* Pagination Dots */}
        <View style={styles.paginationDots}>
          <View style={styles.dot} />
          <View style={styles.dot} />
          <View style={[styles.dot, styles.dotActive]} />
        </View>

        {/* Get Started Button */}
        <TouchableOpacity
          style={styles.getStartedButton}
          onPress={onGetStarted}
          activeOpacity={0.9}>
          <CustomText
            fontSize={17}
            fontWeight="semibold"
            color="#FFFFFF"
            fontFamily="display"
            textAlign="center"
            style={styles.buttonText}>
            Get Started
          </CustomText>
        </TouchableOpacity>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: colors.onboarding.backgroundLight,
    paddingHorizontal: WIDTH.W24,
    paddingTop: HEIGHT.H32,
    paddingBottom: HEIGHT.H40,
    justifyContent: 'space-between',
  },
  mainContent: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    paddingTop: HEIGHT.H32,
  },
  illustrationCard: {
    width: '100%',
    maxWidth: WIDTH.W343,
    aspectRatio: 4 / 5,
    maxHeight: RfH(420),
    backgroundColor: colors.onboarding.surfaceCard,
    borderRadius: BorderRadius.BR16,
    overflow: 'hidden',
    justifyContent: 'center',
    alignItems: 'center',
    marginBottom: HEIGHT.H40,
    ...Platform.select({
      ios: {
        // box-shadow: 0 4px 20px -2px rgba(0, 0, 0, 0.04)
        shadowColor: 'rgba(0, 0, 0, 0.04)',
        shadowOffset: {width: 0, height: 4},
        shadowOpacity: 1,
        shadowRadius: 20,
      },
      android: {
        elevation: 4,
      },
    }),
  },
  textContent: {
    alignItems: 'center',
    gap: HEIGHT.H12,
    paddingHorizontal: WIDTH.W8,
  },
  title: {
    lineHeight: FontSize[28] * 1.1,
    letterSpacing: -0.5,
  },
  subtitle: {
    maxWidth: RfW(260),
    lineHeight: FontSize[16] * 1.6,
  },
  footer: {
    width: '100%',
    alignItems: 'center',
    gap: HEIGHT.H40,
    paddingTop: HEIGHT.H16,
  },
  paginationDots: {
    flexDirection: 'row',
    gap: WIDTH.W12,
    alignItems: 'center',
  },
  dot: {
    width: RfW(6),
    height: RfW(6),
    borderRadius: RfW(3),
    backgroundColor: colors.onboarding.dotInactive,
  },
  dotActive: {
    backgroundColor: colors.onboarding.primary,
  },
  getStartedButton: {
    width: '100%',
    height: HEIGHT.H52,
    backgroundColor: colors.onboarding.primary,
    borderRadius: BorderRadius.BR12,
    alignItems: 'center',
    justifyContent: 'center',
    ...Platform.select({
      ios: {
        // shadow-lg shadow-[#C4A77D]/10: "0 4px 20px rgba(196, 167, 125, 0.15)"
        shadowColor: 'rgba(196, 167, 125, 0.15)',
        shadowOffset: {width: 0, height: 4},
        shadowOpacity: 1,
        shadowRadius: 20,
      },
      android: {
        elevation: 6,
      },
    }),
  },
  buttonText: {
    letterSpacing: 0.5,
  },
});
