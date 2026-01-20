import React from 'react';
import Svg, {Rect, Circle, Path} from 'react-native-svg';

interface CardStackIconProps {
  width?: number;
  height?: number;
  primaryColor?: string;
  cardBackColor?: string;
  cardFrontColor?: string;
  skeletonColor?: string;
}

export const CardStackIcon: React.FC<CardStackIconProps> = ({
  width = 200,
  height = 240,
  primaryColor = '#C4A77D',
  cardBackColor = '#F5F5F4',
  cardFrontColor = '#FAFAF8',
  skeletonColor = '#E7E5E4',
}) => {
  // Calculate responsive sizes
  const cardWidth = width * 0.6;
  const cardHeight = height * 0.7;
  const cardRadius = 12;
  const iconSize = width * 0.25;
  const iconRadius = iconSize / 2;

  // Center positions
  const centerX = width / 2;
  const centerY = height / 2;

  // Card positions (stacked vertically, slightly offset to the left)
  const backCardX = centerX - cardWidth / 2 - 12;
  const backCardY = centerY - cardHeight / 2 - 8;
  const middleCardX = centerX - cardWidth / 2 - 6;
  const middleCardY = centerY - cardHeight / 2 - 4;
  const frontCardX = centerX - cardWidth / 2;
  const frontCardY = centerY - cardHeight / 2;

  // Icon position on front card
  const iconX = centerX;
  const iconY = centerY - cardHeight / 2 + 40;

  // Palette shape (kidney bean shape)
  const palettePath = `M ${iconX - iconRadius * 0.3} ${iconY - iconRadius * 0.2}
    C ${iconX - iconRadius * 0.5} ${iconY - iconRadius * 0.4},
      ${iconX - iconRadius * 0.6} ${iconY},
      ${iconX - iconRadius * 0.3} ${iconY + iconRadius * 0.3}
    C ${iconX} ${iconY + iconRadius * 0.4},
      ${iconX + iconRadius * 0.3} ${iconY + iconRadius * 0.2},
      ${iconX + iconRadius * 0.4} ${iconY}
    C ${iconX + iconRadius * 0.5} ${iconY - iconRadius * 0.2},
      ${iconX + iconRadius * 0.3} ${iconY - iconRadius * 0.4},
      ${iconX} ${iconY - iconRadius * 0.3}
    C ${iconX - iconRadius * 0.2} ${iconY - iconRadius * 0.3},
      ${iconX - iconRadius * 0.3} ${iconY - iconRadius * 0.2},
      ${iconX - iconRadius * 0.3} ${iconY - iconRadius * 0.2} Z`;

  // Text placeholder lines position
  const lineY1 = iconY + iconRadius + 20;
  const lineY2 = lineY1 + 12;
  const lineWidth1 = cardWidth * 0.7;
  const lineWidth2 = cardWidth * 0.5;

  return (
    <Svg width={width} height={height} viewBox={`0 0 ${width} ${height}`}>
      {/* Back Card (farthest) */}
      <Rect
        x={backCardX}
        y={backCardY}
        width={cardWidth}
        height={cardHeight}
        rx={cardRadius}
        fill={cardBackColor}
        opacity={0.6}
      />

      {/* Middle Card */}
      <Rect
        x={middleCardX}
        y={middleCardY}
        width={cardWidth}
        height={cardHeight}
        rx={cardRadius}
        fill={cardBackColor}
        opacity={0.8}
      />

      {/* Front Card */}
      <Rect
        x={frontCardX}
        y={frontCardY}
        width={cardWidth}
        height={cardHeight}
        rx={cardRadius}
        fill={cardFrontColor}
      />

      {/* Icon Circle on Front Card */}
      <Circle
        cx={iconX}
        cy={iconY}
        r={iconRadius}
        fill={primaryColor}
        opacity={0.1}
      />

      {/* Palette Shape */}
      <Path d={palettePath} fill={primaryColor} />

      {/* Three dots inside palette */}
      <Circle
        cx={iconX - iconRadius * 0.15}
        cy={iconY - iconRadius * 0.15}
        r={iconRadius * 0.15}
        fill={cardFrontColor}
      />
      <Circle
        cx={iconX + iconRadius * 0.15}
        cy={iconY - iconRadius * 0.15}
        r={iconRadius * 0.15}
        fill={cardFrontColor}
      />
      <Circle
        cx={iconX}
        cy={iconY + iconRadius * 0.2}
        r={iconRadius * 0.15}
        fill={cardFrontColor}
      />

      {/* Text placeholder lines */}
      <Rect
        x={centerX - lineWidth1 / 2}
        y={lineY1}
        width={lineWidth1}
        height={4}
        rx={2}
        fill={skeletonColor}
      />
      <Rect
        x={centerX - lineWidth2 / 2}
        y={lineY2}
        width={lineWidth2}
        height={4}
        rx={2}
        fill={skeletonColor}
      />
    </Svg>
  );
};
