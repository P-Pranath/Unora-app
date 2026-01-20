import React from 'react';
import {Text, TextProps, TextStyle} from 'react-native';
import {FontSize, fontFamily, fontWeight} from '../../theme';

export interface CustomTextProps extends Omit<TextProps, 'style'> {
  /**
   * Font size - use numeric key from FontSize (e.g., 12, 14, 16)
   * Uses responsive font sizing based on standard screen size (812px)
   * @default 14
   */
  fontSize?: keyof typeof FontSize;

  /**
   * Font weight - use fontWeight constants or string values
   * @default 'regular'
   */
  fontWeight?: keyof typeof fontWeight | TextStyle['fontWeight'];

  /**
   * Font style
   */
  fontStyle?: TextStyle['fontStyle'];

  /**
   * Text color - hex value or theme color path (e.g., '#000000' or 'text.primary')
   */
  color?: string;

  /**
   * Font family - use fontFamily constants
   * @default 'regular'
   */
  fontFamily?: keyof typeof fontFamily;

  /**
   * Text alignment
   * @default 'left'
   */
  textAlign?: TextStyle['textAlign'];

  /**
   * Additional style overrides
   */
  style?: TextStyle | TextStyle[];

  /**
   * Number of lines to display (0 = unlimited)
   */
  numberOfLines?: number;

  /**
   * Whether text is selectable
   * @default false
   */
  selectable?: boolean;

  /**
   * Children content
   */
  children?: React.ReactNode;
}

/**
 * CustomText - A responsive text component following the app's design system
 *
 * Features:
 * - Responsive font sizing using FontSize constants (RFValue-based)
 * - Simple color support (hex values or theme color paths)
 * - TypeScript support with proper types
 *
 * @example
 * ```tsx
 * <CustomText fontSize={16} fontWeight="bold" color="#000000">
 *   Hello World
 * </CustomText>
 * ```
 */
export const CustomText: React.FC<CustomTextProps> = ({
  fontSize = 14,
  fontWeight: fontWeightProp = 'regular',
  fontStyle,
  color,
  fontFamily: fontFamilyProp = 'regular',
  textAlign = 'left',
  style,
  numberOfLines,
  selectable = false,
  children,
  ...restProps
}) => {
  // Get responsive font size from FontSize constants
  const getFontSize = (): number => {
    if (fontSize in FontSize) {
      return FontSize[fontSize as keyof typeof FontSize];
    }
    // Fallback to default if invalid key
    return FontSize[14];
  };

  // Get font weight
  const getFontWeight = (): TextStyle['fontWeight'] => {
    if (typeof fontWeightProp === 'string') {
      // Check if it's a key in fontWeight constants
      if (fontWeightProp in fontWeight) {
        return fontWeight[
          fontWeightProp as keyof typeof fontWeight
        ] as TextStyle['fontWeight'];
      }
      // Otherwise use the string value directly (e.g., '400', 'bold')
      return fontWeightProp as TextStyle['fontWeight'];
    }
    return fontWeightProp;
  };

  const textStyle: TextStyle = {
    fontSize: getFontSize(),
    fontWeight: getFontWeight(),
    fontStyle,
    fontFamily: fontFamily[fontFamilyProp],
    color: color,
    textAlign,
  };

  return (
    <Text
      numberOfLines={numberOfLines}
      style={[textStyle, style]}
      selectable={selectable}
      {...restProps}>
      {children}
    </Text>
  );
};

export default CustomText;
