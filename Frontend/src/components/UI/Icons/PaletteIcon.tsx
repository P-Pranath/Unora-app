import React from 'react';
import Svg, {Path} from 'react-native-svg';

interface PaletteIconProps {
  size?: number;
  color?: string;
}

export const PaletteIcon: React.FC<PaletteIconProps> = ({
  size = 48,
  color = '#C4A77D',
}) => (
  <Svg width={size} height={size} viewBox="0 0 24 24" fill="none">
    <Path
      d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10c.55 0 1-.45 1-1v-3.5c0-.28-.22-.5-.5-.5-.83 0-1.5-.67-1.5-1.5 0-.28.22-.5.5-.5.28 0 .5-.22.5-.5 0-.83.67-1.5 1.5-1.5.28 0 .5-.22.5-.5 0-.83.67-1.5 1.5-1.5.28 0 .5-.22.5-.5V3c0-.55.45-1 1-1zm-2 13c-1.1 0-2-.9-2-2s.9-2 2-2 2 .9 2 2-.9 2-2 2zm2-8c-1.1 0-2-.9-2-2s.9-2 2-2 2 .9 2 2-.9 2-2 2zm4 4c-1.1 0-2-.9-2-2s.9-2 2-2 2 .9 2 2-.9 2-2 2z"
      fill={color}
    />
  </Svg>
);
