#!/bin/bash
# Create macOS .icns icon from source image with proper rounded corners
# Usage: ./scripts/create-icon.sh <source-image.png>

set -e

SOURCE_IMAGE=$1

if [ -z "$SOURCE_IMAGE" ] || [ ! -f "$SOURCE_IMAGE" ]; then
    echo "Usage: $0 <source-image.png>"
    echo ""
    echo "The source image should be:"
    echo "  - At least 1024x1024 pixels"
    echo "  - PNG format with transparency"
    echo "  - Square aspect ratio"
    echo ""
    echo "The script will apply macOS Big Sur style rounded corners."
    exit 1
fi

ICONSET_DIR="chowkidar-icon.iconset"
OUTPUT_ICON="chowkidar.icns"
TEMP_DIR="$(mktemp -d)"

echo "Creating macOS icon from $SOURCE_IMAGE..."

# Function to apply rounded corners using ImageMagick
apply_rounded_corners() {
    local input=$1
    local output=$2
    local size=$3

    # Calculate corner radius (approximately 22.37% of size for Big Sur style)
    local radius=$(echo "$size * 0.2237" | bc | cut -d. -f1)

    if command -v magick >/dev/null 2>&1 || command -v convert >/dev/null 2>&1; then
        # Use ImageMagick if available
        local magick_cmd="convert"
        if command -v magick >/dev/null 2>&1; then
            magick_cmd="magick"
        fi

        $magick_cmd "$input" \
            \( +clone -alpha extract \
            -draw "fill black polygon 0,0 0,$radius $radius,0 fill white circle $radius,$radius $radius,0" \
            \( +clone -flip \) -compose Multiply -composite \
            \( +clone -flop \) -compose Multiply -composite \
            \) -alpha off -compose CopyOpacity -composite "$output"
    else
        # Fallback: just resize without rounding
        cp "$input" "$output"
    fi
}

# Create iconset directory
mkdir -p "$ICONSET_DIR"

# Check if ImageMagick is available
if ! command -v magick >/dev/null 2>&1 && ! command -v convert >/dev/null 2>&1; then
    echo "⚠️  ImageMagick not found - icons will not have rounded corners"
    echo "   Install with: brew install imagemagick"
    echo ""
    USE_ROUNDING=false
else
    USE_ROUNDING=true
    echo "✓ ImageMagick found - applying macOS rounded corners"
fi

# Generate all required icon sizes with rounded corners
declare -a sizes=(
    "16:icon_16x16.png"
    "32:icon_16x16@2x.png"
    "32:icon_32x32.png"
    "64:icon_32x32@2x.png"
    "128:icon_128x128.png"
    "256:icon_128x128@2x.png"
    "256:icon_256x256.png"
    "512:icon_256x256@2x.png"
    "512:icon_512x512.png"
    "1024:icon_512x512@2x.png"
)

for size_file in "${sizes[@]}"; do
    IFS=':' read -r size filename <<< "$size_file"
    temp_resized="$TEMP_DIR/resized_${size}.png"

    # First resize
    sips -z $size $size "$SOURCE_IMAGE" --out "$temp_resized" > /dev/null 2>&1

    # Then apply rounded corners if available
    if [ "$USE_ROUNDING" = true ]; then
        apply_rounded_corners "$temp_resized" "$ICONSET_DIR/$filename" "$size"
    else
        cp "$temp_resized" "$ICONSET_DIR/$filename"
    fi
done

# Convert to .icns
echo "Creating .icns file..."
iconutil -c icns "$ICONSET_DIR" -o "$OUTPUT_ICON"

# Clean up
rm -rf "$ICONSET_DIR"
rm -rf "$TEMP_DIR"

echo ""
echo "✅ Created $OUTPUT_ICON"
if [ "$USE_ROUNDING" = true ]; then
    echo "✓ Applied macOS Big Sur style rounded corners"
fi
echo ""
echo "Next steps:"
echo "  1. Move the .icns file to your project:"
echo "     mkdir -p packaging/assets"
echo "     mv $OUTPUT_ICON packaging/assets/"
echo ""
echo "  2. Rebuild the DMG - the script will automatically use it:"
echo "     make package-dmg"
echo ""
echo "Note: For best results, design your icon following Apple's guidelines:"
echo "  https://developer.apple.com/design/human-interface-guidelines/app-icons"
