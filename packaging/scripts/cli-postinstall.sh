#!/bin/sh
if command -v systemctl >/dev/null 2>&1; then
  echo "To enable and start the service, run:"
  echo "  sudo systemctl enable chowkidar"
  echo "  sudo systemctl start chowkidar"
  echo ""
  echo "First, configure chowkidar by running:"
  echo "  chowkidar setup"
fi
