import React from 'react';
import {View, StyleSheet, TouchableOpacity, Platform} from 'react-native';
import {CustomText} from '@/components/UI/CustomText';
import {CheckIcon, LockOpenIcon} from '@/components/UI/Icons';
import {colors, FontSize} from '@/theme';
import {RfW, RfH} from '@/utils/helpers';
import {HEIGHT, WIDTH, BorderRadius} from '@/theme/sizes';
import {LinearGradient} from 'react-native-linear-gradient';

interface OnboardingSlide2Props {
  onNext?: () => void;
  onSkip?: () => void;
}

export const OnboardingSlide2: React.FC<OnboardingSlide2Props> = ({
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

          {/* Content Container */}
          <View style={styles.contentContainer}>
            {/* User Profiles at Top */}
            <View style={styles.userProfilesContainer}>
              <View style={styles.userProfile}>
                <View style={styles.profileImagePlaceholder} />
              </View>
              <View style={styles.connectorLine} />
              <View style={styles.userProfile}>
                <View style={styles.profileImagePlaceholder} />
              </View>
            </View>

            {/* Progression Path */}
            <View style={styles.progressionContainer}>
              {/* Vertical gradient line */}
              <LinearGradient
                colors={[
                  'transparent',
                  'rgba(196, 167, 125, 0.3)',
                  'transparent',
                ]}
                start={{x: 0, y: 0}}
                end={{x: 0, y: 1}}
                style={styles.verticalLine}
              />

              {/* Dots Grid */}
              <View style={styles.dotsGrid}>
                {/* Row 1 */}
                <View style={styles.gridRow}>
                  <View style={[styles.dot, styles.dotActive]} />
                  <View
                    style={[styles.dot, styles.dotActive, styles.dotOffsetTop]}
                  />
                  <View style={[styles.dot, styles.dotActive]} />
                </View>

                {/* Row 2 */}
                <View style={styles.gridRow}>
                  <View
                    style={[styles.dot, styles.dotActive, styles.dotOffsetLeft]}
                  />
                  <View style={styles.checkCircleContainer}>
                    <View style={styles.checkCircle}>
                      <CheckIcon size={16} color={colors.onboarding.primary} />
                    </View>
                  </View>
                  <View
                    style={[
                      styles.dot,
                      styles.dotActive,
                      styles.dotOffsetRight,
                    ]}
                  />
                </View>

                {/* Row 3 */}
                <View style={styles.gridRow}>
                  <View style={styles.dot} />
                  <View style={[styles.dot, styles.dotOffsetTop]} />
                  <View style={styles.dot} />
                </View>
              </View>
            </View>

            {/* Lock Icon and Text */}
            <View style={styles.lockContainer}>
              <View style={styles.lockCard}>
                <LockOpenIcon size={24} color={colors.onboarding.primary} />
              </View>
              <CustomText
                fontSize={10}
                fontWeight="semibold"
                color={colors.onboarding.primary}
                fontFamily="body"
                textAlign="center"
                style={styles.unlockText}>
                DAY 15 UNLOCKED
              </CustomText>
            </View>
          </View>
        </View>
      </View>

      {/* Bottom Content */}
      <View style={styles.bottomContent}>
        {/* Text Content */}
        <View style={styles.textContent}>
          <CustomText
            fontSize={28}
            fontWeight="bold"
            color={colors.onboarding.textHead}
            fontFamily="display"
            textAlign="center"
            style={styles.title}>
            15 days of showing up.
          </CustomText>
          <CustomText
            fontSize={16}
            fontWeight="regular"
            color={colors.onboarding.textBody}
            fontFamily="body"
            textAlign="center"
            style={styles.subtitle}>
            Chat unlocks after 15 days of mutual consistency.
          </CustomText>
        </View>
      </View>

      {/* Footer */}
      <View style={styles.footer}>
        {/* Pagination Dots */}
        <View style={styles.paginationDots}>
          <View style={styles.paginationDot} />
          <View style={[styles.paginationDot, styles.paginationDotActive]} />
          <View style={styles.paginationDot} />
        </View>

        {/* Action Buttons */}
        <View style={styles.actions}>
          <TouchableOpacity
            style={styles.nextButton}
            onPress={onNext}
            activeOpacity={0.9}>
            <CustomText
              fontSize={17}
              fontWeight="semibold"
              color="#FFFFFF"
              fontFamily="display"
              textAlign="center"
              style={styles.buttonText}>
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
    paddingTop: HEIGHT.H48,
    paddingBottom: HEIGHT.H40,
    justifyContent: 'space-between',
  },
  cardContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    marginBottom: HEIGHT.H32,
  },
  mainCard: {
    width: '100%',
    maxWidth: WIDTH.W343,
    aspectRatio: 4 / 5,
    maxHeight: RfH(420),
    backgroundColor: colors.onboarding.surfaceCard,
    borderRadius: RfW(24),
    overflow: 'hidden',
    padding: WIDTH.W24,
    ...Platform.select({
      ios: {
        // shadow-soft: "0 20px 40px -10px rgba(0, 0, 0, 0.05)"
        shadowColor: 'rgba(0, 0, 0, 0.05)',
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
    width: RfW(192),
    height: RfW(192),
    marginLeft: RfW(-96),
    marginTop: RfW(-96),
    borderRadius: RfW(96),
    backgroundColor: 'rgba(196, 167, 125, 0.1)',
    ...Platform.select({
      ios: {
        shadowColor: 'rgba(196, 167, 125, 0.1)',
        shadowOffset: {width: 0, height: 0},
        shadowOpacity: 1,
        shadowRadius: 60,
      },
    }),
  },
  contentContainer: {
    flex: 1,
    justifyContent: 'space-between',
    alignItems: 'center',
    paddingVertical: HEIGHT.H16,
    zIndex: 10,
  },
  userProfilesContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    gap: WIDTH.W16,
    marginBottom: HEIGHT.H8,
    zIndex: 10,
  },
  userProfile: {
    width: RfW(48),
    height: RfW(48),
    borderRadius: RfW(24),
    backgroundColor: '#F9FAFB',
    borderWidth: 1,
    borderColor: '#F3F4F6',
    overflow: 'hidden',
    ...Platform.select({
      ios: {
        // shadow-sm: small shadow
        shadowColor: 'rgba(0, 0, 0, 0.1)',
        shadowOffset: {width: 0, height: 1},
        shadowOpacity: 1,
        shadowRadius: 3,
      },
      android: {
        elevation: 2,
      },
    }),
  },
  profileImagePlaceholder: {
    width: '100%',
    height: '100%',
    backgroundColor: '#E5E7EB',
    opacity: 0.9,
  },
  connectorLine: {
    width: RfW(32),
    height: 1,
    backgroundColor: '#E5E7EB',
  },
  progressionContainer: {
    flex: 1,
    width: '100%',
    justifyContent: 'center',
    alignItems: 'center',
    paddingVertical: HEIGHT.H16,
    position: 'relative',
  },
  verticalLine: {
    position: 'absolute',
    top: 0,
    bottom: 0,
    width: 1,
    left: '50%',
    marginLeft: -0.5,
  },
  dotsGrid: {
    width: RfW(120),
    justifyContent: 'center',
    alignItems: 'center',
    gap: RfH(12),
    zIndex: 10,
  },
  gridRow: {
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
    gap: RfW(32),
  },
  dot: {
    width: RfW(8),
    height: RfW(8),
    borderRadius: RfW(4),
    backgroundColor: '#E5E7EB',
  },
  dotActive: {
    backgroundColor: 'rgba(196, 167, 125, 0.6)',
  },
  dotOffsetTop: {
    marginTop: RfH(16),
  },
  dotOffsetLeft: {
    marginLeft: RfW(8),
  },
  dotOffsetRight: {
    marginRight: RfW(8),
  },
  checkCircleContainer: {
    width: RfW(32),
    height: RfW(32),
    justifyContent: 'center',
    alignItems: 'center',
  },
  checkCircle: {
    width: RfW(32),
    height: RfW(32),
    borderRadius: RfW(16),
    backgroundColor: '#FFFFFF',
    borderWidth: 1,
    borderColor: 'rgba(196, 167, 125, 0.2)',
    justifyContent: 'center',
    alignItems: 'center',
    ...Platform.select({
      ios: {
        // shadow-lg shadow-primary/10: "0 4px 20px rgba(196, 167, 125, 0.15)"
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
  lockContainer: {
    alignItems: 'center',
    marginTop: HEIGHT.H8,
    zIndex: 10,
  },
  lockCard: {
    width: RfW(56),
    height: RfW(56),
    borderRadius: BorderRadius.BR12,
    backgroundColor: '#FFFFFF',
    borderWidth: 1,
    borderColor: '#F3F4F6',
    justifyContent: 'center',
    alignItems: 'center',
    ...Platform.select({
      ios: {
        // shadow-[0_4px_20px_-5px_rgba(0,0,0,0.05)]: "0 4px 20px -5px rgba(0, 0, 0, 0.05)"
        shadowColor: 'rgba(0, 0, 0, 0.05)',
        shadowOffset: {width: 0, height: 4},
        shadowOpacity: 1,
        shadowRadius: 20,
      },
      android: {
        elevation: 4,
      },
    }),
  },
  unlockText: {
    marginTop: HEIGHT.H12,
    letterSpacing: 0.15 * 10, // 0.15em * 10px
    textTransform: 'uppercase',
  },
  bottomContent: {
    alignItems: 'center',
    gap: HEIGHT.H12,
    marginBottom: HEIGHT.H16,
  },
  textContent: {
    alignItems: 'center',
    gap: HEIGHT.H12,
  },
  title: {
    lineHeight: FontSize[28] * 1.2,
    letterSpacing: -0.5,
  },
  subtitle: {
    maxWidth: RfW(280),
    lineHeight: FontSize[16] * 1.6,
  },
  footer: {
    width: '100%',
    alignItems: 'center',
    gap: HEIGHT.H32,
    paddingTop: HEIGHT.H16,
  },
  paginationDots: {
    flexDirection: 'row',
    gap: WIDTH.W12,
    alignItems: 'center',
  },
  paginationDot: {
    width: RfW(6),
    height: RfW(6),
    borderRadius: RfW(3),
    backgroundColor: colors.onboarding.dotInactive,
  },
  paginationDotActive: {
    backgroundColor: colors.onboarding.primary,
  },
  actions: {
    width: '100%',
    gap: HEIGHT.H16,
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
  buttonText: {
    letterSpacing: 0.5,
  },
  skipButton: {
    width: '100%',
    alignItems: 'center',
    justifyContent: 'center',
    paddingVertical: HEIGHT.H8,
    opacity: 0.7,
  },
  skipText: {
    // Skip text styling
  },
});
