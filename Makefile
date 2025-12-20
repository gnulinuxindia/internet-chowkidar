# Makefile for Internet Chowkidar
# Handles local builds and packaging

.PHONY: all cli gui clean test install help
.PHONY: package-dmg package-msi
.PHONY: release-cli release-gui release-local

# Variables
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS := -s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)

CLI_BINARY := chowkidar
GUI_BINARY := chowkidar-gui

# Platform detection
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
    PLATFORM := macos
endif
ifeq ($(UNAME_S),Linux)
    PLATFORM := linux
endif
ifeq ($(OS),Windows_NT)
    PLATFORM := windows
endif

## help: Show this help message
help:
	@echo "Internet Chowkidar Build System"
	@echo ""
	@echo "Usage: make [target]"
	@echo ""
	@echo "Common targets:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## all: Build both CLI and GUI for current platform
all: cli gui

## cli: Build CLI version
cli:
	@echo "Building CLI..."
	@mkdir -p bin
	go build -ldflags="$(LDFLAGS)" -o bin/$(CLI_BINARY) ./cmd/chowkidar

## gui: Build GUI version for current platform
gui:
	@echo "Building GUI for $(PLATFORM)..."
	@mkdir -p bin
ifeq ($(PLATFORM),macos)
	CGO_ENABLED=1 go build -ldflags="$(LDFLAGS)" -o bin/$(GUI_BINARY) ./gui/chowkidar
else ifeq ($(PLATFORM),linux)
	CGO_ENABLED=1 go build -ldflags="$(LDFLAGS)" -o bin/$(GUI_BINARY) ./gui/chowkidar
else ifeq ($(PLATFORM),windows)
	CGO_ENABLED=1 go build -ldflags="$(LDFLAGS)" -o bin/$(GUI_BINARY).exe ./gui/chowkidar
endif

## package-dmg: Create macOS DMG installer (macOS only)
## Usage: make package-dmg [SIGN_IDENTITY="Developer ID Application: Your Name"]
package-dmg: gui
	@echo "Creating macOS DMG..."
	@if [ "$(PLATFORM)" != "macos" ]; then \
		echo "Error: DMG creation only supported on macOS"; \
		exit 1; \
	fi
	@./scripts/build-dmg.sh $(VERSION) "$(SIGN_IDENTITY)"

## package-msi: Create Windows MSI installer (Windows only, requires WiX)
package-msi:
	@echo "Creating Windows MSI..."
	@if [ "$(PLATFORM)" != "windows" ]; then \
		echo "Error: MSI creation only supported on Windows"; \
		echo "Use scripts/build-msi.sh on Windows with WiX installed"; \
		exit 1; \
	fi
	@./scripts/build-msi.sh $(VERSION)

## test: Run tests
test:
	go test -v -race -coverprofile=coverage.out ./...

## clean: Remove build artifacts
clean:
	rm -rf bin/
	rm -rf dist/
	rm -f coverage.out

## install: Install binaries to system (requires sudo)
install: all
	@echo "Installing binaries..."
	sudo install -m 755 bin/$(CLI_BINARY) /usr/local/bin/$(CLI_BINARY)
ifeq ($(PLATFORM),linux)
	sudo install -m 755 bin/$(GUI_BINARY) /usr/local/bin/$(GUI_BINARY)
	sudo install -m 644 packaging/chowkidar-gui.desktop /usr/share/applications/
else ifeq ($(PLATFORM),macos)
	sudo install -m 755 bin/$(GUI_BINARY) /usr/local/bin/$(GUI_BINARY)
endif
	@echo "Installation complete!"

## release-cli: Build CLI release with GoReleaser (Linux packages: deb, rpm, apk)
release-cli:
	@if ! command -v goreleaser >/dev/null 2>&1; then \
		echo "Installing goreleaser..."; \
		go install github.com/goreleaser/goreleaser/v2@latest; \
	fi
	goreleaser release --clean --snapshot

## release-gui: Build GUI Linux packages with GoReleaser (deb, rpm, apk)
## NOTE: Only works on Linux due to CGO cross-compilation limitations
release-gui:
	@if [ "$(PLATFORM)" != "linux" ]; then \
		echo "⚠️  Warning: GUI package building requires Linux"; \
		echo "   Cross-compiling CGO code is not supported"; \
		echo "   Use GitHub Actions or build on a Linux machine"; \
		exit 1; \
	fi
	@if ! command -v goreleaser >/dev/null 2>&1; then \
		echo "Installing goreleaser..."; \
		go install github.com/goreleaser/goreleaser/v2@latest; \
	fi
	goreleaser release --clean --snapshot --config .goreleaser-gui.yaml

## release-local: Test both CLI and GUI releases locally
release-local: release-cli release-gui
	@echo ""
	@echo "✅ Local release build complete!"
	@echo "Check dist/ directory for packages"

.DEFAULT_GOAL := help
