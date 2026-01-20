import React, {useRef, useState} from 'react';
import {View, StyleSheet, Dimensions} from 'react-native';
import Carousel, {ICarouselInstance} from 'react-native-reanimated-carousel';
import {OnboardingSlide1} from './OnboardingSlide1';
import {OnboardingSlide2} from './OnboardingSlide2';
import {OnboardingSlide3} from './OnboardingSlide3';
import {deviceWidth} from '@/utils/helpers';

interface OnboardingCarouselProps {
  onComplete?: () => void;
  onSkip?: () => void;
}

const SLIDES = [
  {id: 1, component: OnboardingSlide1},
  {id: 2, component: OnboardingSlide2},
  {id: 3, component: OnboardingSlide3},
];

export const OnboardingCarousel: React.FC<OnboardingCarouselProps> = ({
  onComplete,
  onSkip,
}) => {
  const carouselRef = useRef<ICarouselInstance>(null);
  const [currentIndex, setCurrentIndex] = useState(0);
  const screenWidth = deviceWidth();
  const screenHeight = Dimensions.get('window').height;

  const handleNext = () => {
    if (currentIndex < SLIDES.length - 1) {
      carouselRef.current?.next();
    } else {
      // Last slide - call onComplete
      onComplete?.();
    }
  };

  const handleSkip = () => {
    onSkip?.();
  };

  const handleGetStarted = () => {
    onComplete?.();
  };

  const handleSnapToItem = (index: number) => {
    setCurrentIndex(index);
  };

  const renderSlide = ({
    item,
    index,
  }: {
    item: (typeof SLIDES)[0];
    index: number;
  }) => {
    const SlideComponent = item.component;

    // Pass appropriate handlers based on slide index
    if (index === 0) {
      // First slide
      return <SlideComponent onNext={handleNext} onSkip={handleSkip} />;
    } else if (index === 1) {
      // Second slide
      return <SlideComponent onNext={handleNext} onSkip={handleSkip} />;
    } else {
      // Third slide (last)
      return <SlideComponent onGetStarted={handleGetStarted} />;
    }
  };

  return (
    <View style={styles.container}>
      <Carousel
        ref={carouselRef}
        width={screenWidth}
        height={screenHeight}
        data={SLIDES}
        renderItem={renderSlide}
        onSnapToItem={handleSnapToItem}
        loop={false}
        enabled={true}
        pagingEnabled={true}
        snapEnabled={true}
        windowSize={3}
      />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
});
