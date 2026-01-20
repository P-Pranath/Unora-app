# Logo Setup for Google Search Results

## Important: Two Different Places

Your logo needs to be in **two places** for it to appear everywhere:

1. **Favicon** (for Google search results) - Goes in `/public` folder
2. **Header logo** (for your website UI) - Already set up in Header component

## Step 1: Add Logo Files to Public Folder

1. Create a `public` folder in your project root (if it doesn't exist)
2. Add these logo files:

### Required Files:
- `favicon.ico` - 32x32px or 16x16px (classic favicon)
- `favicon.png` - 32x32px (modern browsers)
- `logo.png` - Your full logo, ideally 512x512px (for structured data)
- `apple-touch-icon.png` - 180x180px (for iOS devices)

### File Sizes:
- **favicon.ico**: 16x16 or 32x32 pixels
- **favicon.png**: 32x32 pixels  
- **logo.png**: 512x512 pixels (or square, high resolution)
- **apple-touch-icon.png**: 180x180 pixels

## Step 2: Files Are Already Configured!

Your `app/layout.tsx` already references:
- `/favicon.ico` ✅
- `/apple-touch-icon.png` ✅

## Step 3: Update Organization Schema (for Google)

The code has been updated to include your logo in structured data, which helps Google recognize your brand.

## What Shows Where:

- **Browser tab icon**: Uses `favicon.ico` or `favicon.png`
- **Google search results**: Uses `favicon.ico` or `favicon.png`
- **iOS home screen**: Uses `apple-touch-icon.png`
- **Structured data / Knowledge Graph**: Uses `logo.png` (from Organization schema)

## After Adding Files:

1. Place your logo files in the `public` folder
2. The favicon will automatically appear in:
   - Browser tabs
   - Google search results (after Google re-crawls)
   - Bookmarks
3. The header logo can be updated in `components/Header.tsx`

