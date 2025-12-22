#!/bin/sh
# Update icon cache so icons show up immediately
if command -v gtk-update-icon-cache >/dev/null 2>&1; then
    gtk-update-icon-cache -q -t -f /usr/share/icons/hicolor 2>/dev/null || true
fi

# Update desktop database so application appears in menu
if command -v update-desktop-database >/dev/null 2>&1; then
    update-desktop-database -q /usr/share/applications 2>/dev/null || true
fi

echo "Internet Chowkidar GUI installed successfully!"
echo ""
echo "Launch from:"
echo "  - Application menu (Network/Utilities)"
echo "  - Or run: chowkidar-gui"
echo ""
echo "On first run, you'll be guided through setup."
echo "After setup, look for the system tray icon."
