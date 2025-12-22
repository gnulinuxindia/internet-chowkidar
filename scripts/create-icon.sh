#!/bin/bash
# Create icons for all platforms (macOS .icns + Linux PNGs) from source image
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

echo "Creating macOS icon from $SOURCE_IMAGE..."}

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

    cp "$temp_resized" "$ICONSET_DIR/$filename"
done

# Convert to .icns
echo "Creating .icns file..."
iconutil -c icns "$ICONSET_DIR" -o "$OUTPUT_ICON"

# Create Linux PNG icons (for deb/rpm/apk packages)
echo ""
echo "Creating Linux PNG icons..."
LINUX_ICON_DIR="packaging/icons/hicolor"
mkdir -p "$LINUX_ICON_DIR"/{16x16,32x32,48x48,64x64,128x128,256x256,512x512}/apps
mkdir -p packaging/pixmaps

for size in 16 32 48 64 128 256 512; do
    sips -z $size $size "$SOURCE_IMAGE" --out "$LINUX_ICON_DIR/${size}x${size}/apps/chowkidar-gui.png" > /dev/null 2>&1
done

# Fallback icon for /usr/share/pixmaps
sips -z 48 48 "$SOURCE_IMAGE" --out "packaging/pixmaps/chowkidar-gui.png" > /dev/null 2>&1

# Move macOS icon to assets directory
mkdir -p packaging/assets
mv "$OUTPUT_ICON" "packaging/assets/$OUTPUT_ICON"

# Clean up
rm -rf "$ICONSET_DIR"
rm -rf "$TEMP_DIR"

echo ""
echo "✅ Icons created for all platforms:"
echo ""
echo "macOS:"
echo "  - packaging/assets/$OUTPUT_ICON"
if [ "$USE_ROUNDING" = true ]; then
    echo "    (with Big Sur rounded corners)"
fi
echo ""
echo "Linux:"
echo "  - packaging/icons/hicolor/{16-512}x{16-512}/apps/chowkidar-gui.png"
echo "  - packaging/pixmaps/chowkidar-gui.png"
echo ""
echo "All icons ready for packaging. Rebuild to include them:"
echo "  - macOS: make package-dmg"
echo "  - Linux: make release-gui (requires Linux)"
echo "  - All platforms: git tag vX.Y.Z && git push origin vX.Y.Z"
