#!/bin/bash
# Build Windows MSI installer for Internet Chowkidar GUI
# Requires WiX Toolset (https://wixtoolset.org/)
set -e

VERSION=${1:-dev}
MSI_NAME="chowkidar-gui-${VERSION}-windows-amd64.msi"

echo "Building Windows MSI for version ${VERSION}..."

# Check if WiX is installed
if ! command -v candle.exe >/dev/null 2>&1 && ! command -v candle >/dev/null 2>&1; then
    echo "Error: WiX Toolset not found!"
    echo "Install from: https://wixtoolset.org/"
    echo ""
    echo "Alternative: Use scripts/build-msi.ps1 on Windows with WiX installed"
    exit 1
fi

# Build the binary first if not exists
if [ ! -f "bin/chowkidar-gui.exe" ]; then
    echo "Building GUI binary for Windows..."
    GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -ldflags="-s -w" -o bin/chowkidar-gui.exe ./gui/chowkidar
fi

# Create WiX source file
mkdir -p dist
cat > "dist/chowkidar-gui.wxs" << EOF
<?xml version='1.0' encoding='windows-1252'?>
<Wix xmlns='http://schemas.microsoft.com/wix/2006/wi'>
  <Product Name='Internet Chowkidar'
           Id='*'
           UpgradeCode='12345678-1234-1234-1234-123456789012'
           Language='1033'
           Codepage='1252'
           Version='${VERSION}'
           Manufacturer='GNU/Linux India'>

    <Package Id='*'
             Keywords='Installer'
             Description='Internet Chowkidar GUI Installer'
             Manufacturer='GNU/Linux India'
             InstallerVersion='200'
             Languages='1033'
             Compressed='yes'
             SummaryCodepage='1252' />

    <Media Id='1' Cabinet='chowkidar.cab' EmbedCab='yes' />

    <Directory Id='TARGETDIR' Name='SourceDir'>
      <Directory Id='ProgramFilesFolder'>
        <Directory Id='INSTALLDIR' Name='Internet Chowkidar'>
          <Component Id='MainExecutable' Guid='*'>
            <File Id='ChowkidarGUIEXE'
                  Name='chowkidar-gui.exe'
                  Source='../bin/chowkidar-gui.exe'
                  KeyPath='yes' />
          </Component>

          <Component Id='LicenseFile' Guid='*'>
            <File Id='LICENSE'
                  Name='LICENSE'
                  Source='../LICENSE' />
          </Component>

          <Component Id='ReadmeFile' Guid='*'>
            <File Id='README'
                  Name='README.md'
                  Source='../README.md' />
          </Component>
        </Directory>
      </Directory>

      <Directory Id='ProgramMenuFolder'>
        <Directory Id='ApplicationProgramsFolder' Name='Internet Chowkidar'>
          <Component Id='ApplicationShortcut' Guid='*'>
            <Shortcut Id='ApplicationStartMenuShortcut'
                      Name='Internet Chowkidar'
                      Description='Monitor Internet Censorship'
                      Target='[INSTALLDIR]chowkidar-gui.exe'
                      WorkingDirectory='INSTALLDIR' />
            <RemoveFolder Id='ApplicationProgramsFolder' On='uninstall' />
            <RegistryValue Root='HKCU'
                          Key='Software\GNU-Linux-India\InternetChowkidar'
                          Name='installed'
                          Type='integer'
                          Value='1'
                          KeyPath='yes' />
          </Component>
        </Directory>
      </Directory>
    </Directory>

    <Feature Id='Complete' Level='1'>
      <ComponentRef Id='MainExecutable' />
      <ComponentRef Id='LicenseFile' />
      <ComponentRef Id='ReadmeFile' />
      <ComponentRef Id='ApplicationShortcut' />
    </Feature>

    <Property Id='WIXUI_INSTALLDIR' Value='INSTALLDIR' />
    <UIRef Id='WixUI_InstallDir' />
    <UIRef Id='WixUI_ErrorProgressText' />

  </Product>
</Wix>
EOF

# Compile WiX source
echo "Compiling WiX source..."
cd dist
candle chowkidar-gui.wxs

# Link to create MSI
echo "Linking MSI..."
light -ext WixUIExtension chowkidar-gui.wixobj -out "${MSI_NAME}"

cd ..

echo "✅ MSI created: dist/${MSI_NAME}"
echo ""
echo "To distribute:"
echo "1. Test the MSI on a Windows machine"
echo "2. (Optional) Code sign with signtool.exe"
echo "3. Distribute via website or package manager"
