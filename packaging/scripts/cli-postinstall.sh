#!/bin/sh
# Reload systemd to pick up new service file
if command -v systemctl >/dev/null 2>&1; then
  systemctl daemon-reload >/dev/null 2>&1 || true

  # Apply system preset policy (respects systemd presets)
  systemctl preset chowkidar.service >/dev/null 2>&1 || true

  echo "Internet Chowkidar service installed!"
  echo ""
  echo "First, configure chowkidar by running:"
  echo "  chowkidar setup"
  echo ""
  echo "Then enable and start the service:"
  echo "  sudo systemctl enable chowkidar"
  echo "  sudo systemctl start chowkidar"
  echo ""
  echo "Check status with:"
  echo "  sudo systemctl status chowkidar"
fi
