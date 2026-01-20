import React from 'react';
import {View, StyleSheet, Platform} from 'react-native';
import {
  MenuBookIcon,
  FitnessCenterIcon,
  PaletteIcon,
} from '@/components/UI/Icons';
import {colors} from '@/theme';
import {RfW, RfH} from '@/utils/helpers';
import {HEIGHT, WIDTH, BorderRadius} from '@/theme/sizes';

export const CardStack: React.FC = () => {
  return (
    <View style={styles.cardStackContainer}>
      {/* Left Back Card */}
      <View style={[styles.card, styles.leftBackCard]}>
        <MenuBookIcon size={32} color="rgba(196, 167, 125, 0.4)" />
        <View style={styles.skeletonLine} />
        <View style={[styles.skeletonLine, styles.skeletonLineShort]} />
      </View>

      {/* Right Back Card */}
      <View style={[styles.card, styles.rightBackCard]}>
        <FitnessCenterIcon size={32} color="rgba(196, 167, 125, 0.4)" />
        <View style={styles.skeletonLine} />
        <View style={[styles.skeletonLine, styles.skeletonLineShort]} />
      </View>

      {/* Front Card */}
      <View style={[styles.card, styles.frontCard]}>
        <View style={styles.iconCircle}>
          <PaletteIcon size={48} color={colors.onboarding.primary} />
        </View>
        <View style={styles.skeletonContainer}>
          <View style={styles.skeletonLineFull} />
          <View
            style={[styles.skeletonLineFull, styles.skeletonLineTwoThirds]}
          />
        </View>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  cardStackContainer: {
    width: RfW(256), // w-64
    height: RfW(256), // h-64
    justifyContent: 'center',
    alignItems: 'center',
    transform: [{scale: 1.1}], // scale-110
  },
  card: {
    position: 'absolute',
    alignItems: 'center',
    justifyContent: 'center',
    gap: HEIGHT.H8, // gap-2
    borderWidth: 1,
    borderColor: '#F5F5F4', // border-stone-100
    ...Platform.select({
      ios: {
        // shadow-card: "0 10px 20px -5px rgba(28, 28, 25, 0.05)"
        shadowColor: 'rgba(28, 28, 25, 0.05)',
        shadowOffset: {width: 0, height: 5},
        shadowOpacity: 1,
        shadowRadius: 10,
      },
      android: {
        elevation: 4,
      },
    }),
  },
  leftBackCard: {
    width: RfW(128), // w-32
    height: RfH(176), // h-44
    left: RfW(16), // left-4
    top: RfH(32), // top-8
    backgroundColor: colors.onboarding.cardInnerBack, // bg-card-inner-back
    borderRadius: BorderRadius.BR12, // rounded-xl
    paddingVertical: HEIGHT.H16,
    paddingHorizontal: WIDTH.W16,
    transform: [{rotate: '-12deg'}, {translateY: RfH(8)}], // -rotate-12 translate-y-2
    zIndex: 10,
  },
  rightBackCard: {
    width: RfW(128), // w-32
    height: RfH(176), // h-44
    right: RfW(16), // right-4
    top: RfH(32), // top-8
    backgroundColor: colors.onboarding.cardInnerBack, // bg-card-inner-back
    borderRadius: BorderRadius.BR12, // rounded-xl
    paddingVertical: HEIGHT.H16,
    paddingHorizontal: WIDTH.W16,
    transform: [{rotate: '12deg'}, {translateY: RfH(8)}], // rotate-12 translate-y-2
    zIndex: 10,
  },
  frontCard: {
    width: RfW(144), // w-36
    height: RfH(192), // h-48
    backgroundColor: colors.onboarding.cardInnerFront, // bg-card-inner-front
    borderRadius: BorderRadius.BR16, // rounded-2xl
    paddingVertical: HEIGHT.H16,
    paddingHorizontal: WIDTH.W16,
    transform: [{translateY: RfH(-10)}], // translate-y-[-10px]
    zIndex: 20,
    gap: HEIGHT.H12, // gap-3
    ...Platform.select({
      ios: {
        // shadow-soft: "0 20px 40px -10px rgba(28, 28, 25, 0.08)"
        shadowColor: 'rgba(28, 28, 25, 0.08)',
        shadowOffset: {width: 0, height: 10},
        shadowOpacity: 1,
        shadowRadius: 20,
      },
      android: {
        elevation: 8,
      },
    }),
  },
  iconCircle: {
    width: RfW(64), // w-16
    height: RfW(64), // h-16
    borderRadius: RfW(32), // rounded-full
    backgroundColor: 'rgba(196, 167, 125, 0.1)', // bg-primary/10
    alignItems: 'center',
    justifyContent: 'center',
  },
  skeletonContainer: {
    width: '100%',
    paddingHorizontal: WIDTH.W24, // px-6
    gap: HEIGHT.H8, // gap-2
    alignItems: 'center',
  },
  skeletonLine: {
    width: RfW(48), // w-12
    height: RfH(4), // h-1
    backgroundColor: colors.onboarding.skeleton, // bg-skeleton
    borderRadius: BorderRadius.BRFull, // rounded-full
  },
  skeletonLineShort: {
    width: RfW(32), // w-8
  },
  skeletonLineFull: {
    width: '100%',
    height: RfH(6), // h-1.5
    backgroundColor: colors.onboarding.skeleton, // bg-skeleton
    borderRadius: BorderRadius.BRFull, // rounded-full
  },
  skeletonLineTwoThirds: {
    width: '66.67%', // w-2/3
  },
});
