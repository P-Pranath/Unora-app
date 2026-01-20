# Test Images Directory

This directory contains sample images for testing the Vision Validation Service.

## Naming Conventions

The mock AI client uses filename patterns to determine responses:

| Pattern | Response |
|---------|----------|
| `*_pass.*` or `*valid*` | Usable, high confidence |
| `*no_face*` or `*noface*` | Issue: no_face |
| `*blur*` | Issue: blurred |
| `*dark*` | Issue: too_dark |
| `*bright*` | Issue: too_bright |
| `*obstructed*` or `*masked*` | Issue: obstructed_face |
| `*screenshot*` or `*meme*` | Issue: non_photographic |
| `*lowres*` or `*small*` | Issue: low_resolution |
| `*multi_issue*` | Multiple issues: blurred, too_dark |
| `*borderline*` | Usable, medium confidence |
| Default | Random (70% pass rate) |

## CV Validation Thresholds

The CV validation module uses these thresholds (configurable via .env):

| Check | Threshold | Issue Tag |
|-------|-----------|-----------|
| Blur score | < 100 | `blurred` |
| Brightness | < 30 | `too_dark` |
| Brightness | > 225 | `too_bright` |
| Width | < 640px | `low_resolution` |
| Height | < 480px | `low_resolution` |

## Creating Test Images

For actual testing, create images that match these criteria:

### Valid Photos
- Clear, well-lit images
- At least 640x480 resolution
- Contains a visible human face
- Natural/candid photo style

### Invalid Photos (for testing rejection)
- **Blur test**: Apply heavy blur filter
- **Dark test**: Reduce brightness significantly
- **Bright test**: Overexpose the image
- **Low res**: Resize to below 640x480
- **No face**: Landscape or object photos
- **Screenshot**: Screen capture of UI elements
