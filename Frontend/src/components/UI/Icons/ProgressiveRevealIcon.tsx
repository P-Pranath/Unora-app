import React from 'react';
import Svg, {Rect, Path} from 'react-native-svg';

interface ProgressiveRevealIconProps {
  width?: number;
  height?: number;
  primaryColor?: string;
  secondaryColor?: string;
}

export const ProgressiveRevealIcon: React.FC<ProgressiveRevealIconProps> = ({
  width = 220,
  height = 220,
  primaryColor = '#C4A77D',
  secondaryColor = '#4A4A4A',
}) => (
  <Svg width={width} height={height} viewBox="0 0 240 240" fill="none">
    {/* Left bars - gray */}
    <Rect
      x="50"
      y="80"
      width="14"
      height="80"
      rx="7"
      fill={secondaryColor}
      fillOpacity="0.15"
    />
    <Rect
      x="74"
      y="60"
      width="14"
      height="120"
      rx="7"
      fill={secondaryColor}
      fillOpacity="0.4"
    />
    <Rect x="98" y="45" width="14" height="150" rx="7" fill={secondaryColor} />

    {/* Right bars - primary gold */}
    <Rect x="122" y="45" width="14" height="150" rx="7" fill={primaryColor} />
    <Rect
      x="146"
      y="60"
      width="14"
      height="120"
      rx="7"
      fill={primaryColor}
      fillOpacity="0.6"
    />
    <Rect
      x="170"
      y="80"
      width="14"
      height="80"
      rx="7"
      fill={primaryColor}
      fillOpacity="0.3"
    />

    {/* Curved path (P shape) */}
    <Path
      d="M136 70 C 145 70, 155 80, 155 100 C 155 120, 145 130, 136 130"
      fill="none"
      opacity="0.8"
      stroke={primaryColor}
      strokeLinecap="round"
      strokeWidth="3"
    />
  </Svg>
);
