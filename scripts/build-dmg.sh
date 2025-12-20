#!/bin/bash
# Build macOS DMG installer for Internet Chowkidar GUI
# Creates a drag-and-drop installer with custom background
set -e

VERSION=${1:-dev}
SIGN_IDENTITY=${2:-""}
APP_NAME="Internet Chowkidar"
BUNDLE_ID="watch.inet.gui"
DMG_NAME="chowkidar-gui-${VERSION}-macos"
ENTITLEMENTS_FILE="$(mktemp)"

echo "Building macOS DMG installer for version ${VERSION}..."
if [ -n "$SIGN_IDENTITY" ]; then
    echo "Will code sign with identity: $SIGN_IDENTITY"
fi

# Check if create-dmg is installed
if ! command -v create-dmg >/dev/null 2>&1; then
    echo "Installing create-dmg..."
    brew install create-dmg
fi

# Create entitlements file for hardened runtime
cat > "$ENTITLEMENTS_FILE" << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>com.apple.security.cs.allow-jit</key>
    <true/>
    <key>com.apple.security.cs.allow-unsigned-executable-memory</key>
    <true/>
    <key>com.apple.security.cs.disable-library-validation</key>
    <true/>
</dict>
</plist>
EOF

# Build the binary first if not exists
if [ ! -f "bin/chowkidar-gui" ]; then
    echo "Building GUI binary..."
    CGO_ENABLED=1 go build -ldflags="-s -w" -o bin/chowkidar-gui ./gui/chowkidar
fi

# Create app bundle structure
APP_DIR="dist/${APP_NAME}.app"
rm -rf "$APP_DIR"
mkdir -p "$APP_DIR/Contents/MacOS"
mkdir -p "$APP_DIR/Contents/Resources"

# Copy binary
cp bin/chowkidar-gui "$APP_DIR/Contents/MacOS/${APP_NAME}"
chmod +x "$APP_DIR/Contents/MacOS/${APP_NAME}"

# Copy icon if it exists
if [ -f "packaging/assets/chowkidar.icns" ]; then
    echo "Adding app icon..."
    cp packaging/assets/chowkidar.icns "$APP_DIR/Contents/Resources/chowkidar.icns"
fi

# Code sign the binary if signing identity provided
if [ -n "$SIGN_IDENTITY" ]; then
    echo "Code signing binary with hardened runtime..."
    codesign --sign "$SIGN_IDENTITY" \
        --entitlements "$ENTITLEMENTS_FILE" \
        --options runtime \
        --timestamp \
        --force \
        --verbose \
        "$APP_DIR/Contents/MacOS/${APP_NAME}"
fi

# Create Info.plist
cat > "$APP_DIR/Contents/Info.plist" << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>CFBundleExecutable</key>
    <string>${APP_NAME}</string>
    <key>CFBundleIdentifier</key>
    <string>${BUNDLE_ID}</string>
    <key>CFBundleName</key>
    <string>${APP_NAME}</string>
    <key>CFBundlePackageType</key>
    <string>APPL</string>
    <key>CFBundleShortVersionString</key>
    <string>${VERSION}</string>
    <key>CFBundleVersion</key>
    <string>${VERSION}</string>
    <key>CFBundleIconFile</key>
    <string>chowkidar.icns</string>
    <key>LSMinimumSystemVersion</key>
    <string>10.15</string>
    <key>NSHighResolutionCapable</key>
    <true/>
    <key>LSUIElement</key>
    <string>1</string>
    <key>NSHumanReadableCopyright</key>
    <string>Copyright © 2025 GNU/Linux India</string>
</dict>
</plist>
EOF

# Code sign the app bundle if signing identity provided
if [ -n "$SIGN_IDENTITY" ]; then
    echo "Code signing app bundle..."
    codesign --sign "$SIGN_IDENTITY" \
        --entitlements "$ENTITLEMENTS_FILE" \
        --options runtime \
        --timestamp \
        --deep \
        --force \
        --verbose \
        "$APP_DIR"

    # Verify the signature
    echo "Verifying app bundle signature..."
    codesign --verify --deep --strict --verbose=2 "$APP_DIR"
fi

# Create background image for DMG installer
mkdir -p dist/dmg-assets

# Create DMG with create-dmg
echo "Creating DMG installer..."

# Remove old DMG if exists
rm -f "dist/${DMG_NAME}.dmg"

# Prepare DMG icon if app icon exists
DMG_ICON_ARG=""
if [ -f "packaging/assets/chowkidar.icns" ]; then
    DMG_ICON_ARG="--volicon packaging/assets/chowkidar.icns"
fi

if [ -f "packaging/assets/background.png" ]; then
    # Create background image for DMG installer
    mkdir -p dist/dmg-assets
    cp packaging/assets/background.png dist/dmg-assets/background.png
fi

# Create the DMG with drag-and-drop layout
if [ -f "dist/dmg-assets/background.png" ]; then
    # With custom background
    create-dmg \
        --volname "${APP_NAME}" \
        $DMG_ICON_ARG \
        --background "dist/dmg-assets/background.png" \
        --window-pos 200 120 \
        --window-size 700 450 \
        --icon-size 100 \
        --icon "${APP_NAME}.app" 160 220 \
        --hide-extension "${APP_NAME}.app" \
        --app-drop-link 550 220 \
        --no-internet-enable \
        "dist/${DMG_NAME}.dmg" \
        "$APP_DIR" || {
            echo "⚠️  create-dmg with background failed, trying without..."
            create-dmg \
                --volname "${APP_NAME}" \
                $DMG_ICON_ARG \
                --window-pos 200 120 \
                --window-size 600 400 \
                --icon-size 100 \
                --icon "${APP_NAME}.app" 150 200 \
                --hide-extension "${APP_NAME}.app" \
                --app-drop-link 450 200 \
                "dist/${DMG_NAME}.dmg" \
                "$APP_DIR"
        }
else
    # Without custom background
    create-dmg \
        --volname "${APP_NAME}" \
        $DMG_ICON_ARG \
        --window-pos 200 120 \
        --window-size 600 400 \
        --icon-size 100 \
        --icon "${APP_NAME}.app" 150 200 \
        --hide-extension "${APP_NAME}.app" \
        --app-drop-link 450 200 \
        "dist/${DMG_NAME}.dmg" \
        "$APP_DIR"
fi

# Code sign the DMG if signing identity provided
if [ -n "$SIGN_IDENTITY" ]; then
    echo "Code signing DMG..."
    codesign --sign "$SIGN_IDENTITY" \
        --timestamp \
        --force \
        --verbose \
        "dist/${DMG_NAME}.dmg"

    # Verify the DMG signature
    echo "Verifying DMG signature..."
    codesign --verify --verbose=2 "dist/${DMG_NAME}.dmg"

    echo ""
    echo "✅ DMG created and code signed: dist/${DMG_NAME}.dmg"
else
    echo ""
    echo "✅ DMG created: dist/${DMG_NAME}.dmg"
    echo "⚠️  DMG is not code signed (no signing identity provided)"
fi

# Clean up temporary entitlements file
rm -f "$ENTITLEMENTS_FILE"
echo ""
echo "The DMG includes:"
echo "  - ${APP_NAME}.app bundle"
if [ -f "packaging/assets/chowkidar.icns" ]; then
    echo "  - App icon (chowkidar.icns)"
else
    echo "  ⚠️  No app icon (add packaging/assets/chowkidar.icns)"
fi
echo "  - Symlink to /Applications folder"
echo "  - Drag-and-drop layout"
if [ -f "dist/dmg-assets/background.png" ]; then
    echo "  - Custom background image"
fi
echo ""
if [ -n "$SIGN_IDENTITY" ]; then
    echo "Code signing complete! Ready for notarization."
    echo ""
    echo "To notarize with Apple:"
    echo "  xcrun notarytool submit dist/${DMG_NAME}.dmg \\"
    echo "    --keychain-profile \"notarization-profile\" \\"
    echo "    --wait"
    echo ""
    echo "After notarization, staple the ticket:"
    echo "  xcrun stapler staple dist/${DMG_NAME}.dmg"
else
    echo "To build with code signing for distribution:"
    echo "  ./scripts/build-dmg.sh ${VERSION} \"Developer ID Application: Your Name\""
    echo ""
    echo "Or set up a keychain profile for notarization:"
    echo "  xcrun notarytool store-credentials \"notarization-profile\" \\"
    echo "    --apple-id your@email.com \\"
    echo "    --team-id YOUR_TEAM_ID"
fi
echo ""
echo "To test:"
echo "  1. Double-click the DMG"
echo "  2. Drag the app to Applications"
echo "  3. Launch from Applications or Launchpad"
