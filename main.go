package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"image/color"
	"scouting_app/dataBase"
	"scouting_app/tabs"
)

type myTheme struct {
}

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.NRGBA{
				R: 255,
				G: 165,
				B: 0,
				A: 255,
			}
		}
		return color.Black
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {

	return theme.DefaultTheme().Icon(name)
}

func main() {
	a := app.New()
	w := a.NewWindow("SSM")
	db := &dataBase.DB{}
	a.Settings().SetTheme(&myTheme{})
	w.SetContent(tabs.CreateTabs(db))
	w.Show()
	w.Resize(fyne.NewSize(800, 600))

	a.Run()
}
