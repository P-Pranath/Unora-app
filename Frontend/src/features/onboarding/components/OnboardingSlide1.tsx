import React from 'react';
import {View, StyleSheet, TouchableOpacity, Platform} from 'react-native';
import {CustomText} from '@/components/UI/CustomText';
import {CardStack} from './CardStack';
import {colors, FontSize} from '@/theme';
import {RfW, RfH} from '@/utils/helpers';
import {HEIGHT, WIDTH, BorderRadius} from '@/theme/sizes';

interface OnboardingSlide1Props {
  onNext?: () => void;
  onSkip?: () => void;
}

export const OnboardingSlide1: React.FC<OnboardingSlide1Props> = ({
  onNext,
  onSkip,
}) => {
  return (
    <View style={styles.container}>
      {/* Main Card Container */}
      <View style={styles.cardContainer}>
        <View style={styles.mainCard}>
          {/* Blurred background circle */}
          <View style={styles.blurCircle} />

          {/* Card Stack */}
          <View style={styles.cardStackWrapper}>
            <CardStack />
          </View>

          {/* Subtle pagination dots inside card */}
          <View style={styles.internalDots}>
            <View style={styles.internalDot} />
            <View style={styles.internalDot} />
            <View style={styles.internalDot} />
          </View>
        </View>
      </View>

      {/* Bottom Content */}
      <View style={styles.bottomContent}>
        {/* Pagination Dots */}
        <View style={styles.paginationDots}>
          <View style={[styles.dot, styles.dotActive]} />
          <View style={styles.dot} />
          <View style={styles.dot} />
        </View>

        {/* Text Content */}
        <View style={styles.textContent}>
          <CustomText
            fontSize={26}
            fontWeight="bold"
            color={colors.onboarding.textHead}
            fontFamily="display"
            textAlign="center"
            style={styles.title}>
            No photos. Just hobbies.
          </CustomText>
          <CustomText
            fontSize={16}
            fontWeight="regular"
            color={colors.onboarding.textBody}
            fontFamily="body"
            textAlign="center"
            style={styles.subtitle}>
            Discover people through what they do, not how they look.
          </CustomText>
        </View>

        {/* Action Buttons */}
        <View style={styles.actions}>
          <TouchableOpacity
            style={styles.nextButton}
            onPress={onNext}
            activeOpacity={0.9}>
            <CustomText
              fontSize={16}
              fontWeight="semibold"
              color="#FFFFFF"
              fontFamily="display"
              textAlign="center">
              Next
            </CustomText>
          </TouchableOpacity>

          <TouchableOpacity
            style={styles.skipButton}
            onPress={onSkip}
            activeOpacity={0.7}>
            <CustomText
              fontSize={14}
              fontWeight="medium"
              color={colors.onboarding.primary}
              fontFamily="body"
              textAlign="center"
              style={styles.skipText}>
              Skip →
            </CustomText>
          </TouchableOpacity>
        </View>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: colors.onboarding.backgroundLight,
    paddingHorizontal: WIDTH.W24,
    paddingTop: HEIGHT.H16,
    paddingBottom: HEIGHT.H40,
    justifyContent: 'space-between',
  },
  cardContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    paddingVertical: HEIGHT.H32,
  },
  mainCard: {
    width: '100%',
    maxWidth: WIDTH.W343,
    aspectRatio: 4 / 5,
    maxHeight: RfH(420),
    backgroundColor: colors.onboarding.surfaceCard,
    borderRadius: RfW(40),
    borderWidth: 1,
    borderColor: '#F5F5F4',
    overflow: 'hidden',
    ...Platform.select({
      ios: {
        // shadow-soft: "0 20px 40px -10px rgba(28, 28, 25, 0.08)"
        shadowColor: 'rgba(28, 28, 25, 0.08)',
        shadowOffset: {width: 0, height: 20},
        shadowOpacity: 1,
        shadowRadius: 40,
      },
      android: {
        elevation: 12,
      },
    }),
  },
  blurCircle: {
    position: 'absolute',
    top: '50%',
    left: '50%',
    width: RfW(256),
    height: RfW(256),
    marginLeft: RfW(-128),
    marginTop: RfW(-128),
    borderRadius: RfW(128),
    backgroundColor: 'rgba(196, 167, 125, 0.05)',
    ...Platform.select({
      ios: {
        shadowColor: 'rgba(196, 167, 125, 0.1)',
        shadowOffset: {width: 0, height: 0},
        shadowOpacity: 1,
        shadowRadius: 80,
      },
    }),
  },
  internalDots: {
    position: 'absolute',
    bottom: HEIGHT.H32,
    left: 0,
    right: 0,
    flexDirection: 'row',
    justifyContent: 'center',
    gap: WIDTH.W8,
    opacity: 0.1,
  },
  internalDot: {
    width: RfW(6),
    height: RfW(6),
    borderRadius: RfW(3),
    backgroundColor: colors.onboarding.textHead,
  },
  bottomContent: {
    alignItems: 'center',
    gap: HEIGHT.H32,
  },
  cardStackWrapper: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    zIndex: 10,
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
  textContent: {
    alignItems: 'center',
    gap: HEIGHT.H12,
    paddingHorizontal: WIDTH.W8,
  },
  title: {
    lineHeight: FontSize[26] * 1.1,
    letterSpacing: -0.5,
  },
  subtitle: {
    maxWidth: RfW(280),
    lineHeight: FontSize[16] * 1.6,
  },
  actions: {
    width: '100%',
    gap: HEIGHT.H16,
    paddingTop: HEIGHT.H8,
  },
  nextButton: {
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
  skipButton: {
    width: '100%',
    alignItems: 'center',
    justifyContent: 'center',
    paddingVertical: HEIGHT.H8,
    opacity: 0.7,
  },
  skipText: {
    // Skip text styling - consistent across slides
  },
});
