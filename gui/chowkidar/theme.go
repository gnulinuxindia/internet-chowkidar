package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// CustomTheme provides a nicer looking theme
type CustomTheme struct{}

var _ fyne.Theme = (*CustomTheme)(nil)

func (m CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	// Use system default colors but override primary color
	if name == theme.ColorNamePrimary {
		return color.NRGBA{R: 0x3b, G: 0x82, B: 0xf6, A: 0xff} // Material Blue
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m CustomTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m CustomTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m CustomTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNamePadding:
		return 8
	case theme.SizeNameInnerPadding:
		return 8
	default:
		return theme.DefaultTheme().Size(name)
	}
}
