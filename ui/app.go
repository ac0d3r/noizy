package ui

import (
	"bytes"
	"embed"
	"image/png"
	"noizy/player"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	app      fyne.App
	win      fyne.Window
	player   *player.Player
	assetsFS embed.FS
}

func NewApp(name string, player *player.Player, assetsFS embed.FS) *App {
	app := app.New()
	win := app.NewWindow(name)

	return &App{app: app, win: win, player: player, assetsFS: assetsFS}
}

func (a *App) Run() {
	a.buildMainView()
	a.win.Resize(fyne.NewSize(800, 600))
	a.win.ShowAndRun()
}

func (a *App) buildMainView() {
	mainContent := container.NewStack(buildLibraryView(a.player.SystemSounds(),
		func(name string) {
			a.player.Play(name)
		},
		func(name string) {
			a.player.Stop(name)
		},
		func() {
			a.player.StopAll()
		},
	),
		buildScenesView())

	var (
		btnLibrary *widget.Button
		btnScenes  *widget.Button
	)
	setActive := func(index int) {
		if index == 0 {
			btnLibrary.Importance = widget.HighImportance
			btnScenes.Importance = widget.MediumImportance
		} else {
			btnLibrary.Importance = widget.HighImportance
			btnScenes.Importance = widget.MediumImportance
		}
		btnLibrary.Refresh()
		btnScenes.Refresh()

		for i := range mainContent.Objects {
			mainContent.Objects[i].Hide()
		}
		mainContent.Objects[index].Show()
	}

	btnLibrary = widget.NewButtonWithIcon("Sounds", theme.GridIcon(), func() { setActive(0) })
	btnScenes = widget.NewButtonWithIcon("Scenes", theme.BrokenImageIcon(), func() { setActive(1) })

	logoImg, err := a.makeLogoImg()
	if err != nil {
		return
	}

	nav := container.NewVBox(
		container.NewCenter(
			container.NewHBox(
				logoImg,
				widget.NewLabelWithStyle("Noizy", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			),
		),
		btnLibrary, btnScenes,
	)

	setActive(0)

	split := container.NewHSplit(nav, mainContent)
	split.SetOffset(0.2)
	a.win.SetContent(split)
}

func (a *App) makeLogoImg() (*canvas.Image, error) {
	data, err := a.assetsFS.ReadFile("assets/icon.png")
	if err != nil {
		return nil, err
	}

	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	logo := canvas.NewImageFromImage(img)
	logo.SetMinSize(fyne.NewSize(32, 32))
	return logo, nil
}
