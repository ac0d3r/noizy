package ui

import (
	"noizy/player"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func buildLibraryView(soundSet map[string][]player.SoundFile, onPlay, onStop func(string), onStopAll func()) fyne.CanvasObject {
	content := container.NewVBox()

	for category, sounds := range soundSet {
		var soundCards []fyne.CanvasObject
		for _, s := range sounds {
			soundCards = append(soundCards, buildSoundCard(s, onPlay, onStop))
		}
		container.NewVBox(
			widget.NewLabelWithStyle(category, fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			container.NewGridWithColumns(3, soundCards...),
		)

		content.Add(container.NewVBox(
			widget.NewLabelWithStyle(category, fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			container.NewGridWithColumns(4, soundCards...),
		))
		content.Add(widget.NewSeparator())
	}

	scroll := container.NewVScroll(content)
	scroll.SetMinSize(fyne.NewSize(400, 600))

	return scroll
}

func buildSoundCard(s player.SoundFile, onPlay, onStop func(string)) fyne.CanvasObject {
	var isPlaying bool
	var playBtn *widget.Button
	playBtn = widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {
		isPlaying = !isPlaying

		if isPlaying {
			playBtn.SetIcon(theme.MediaPauseIcon())
			onPlay(s.Name)
		} else {
			playBtn.SetIcon(theme.MediaPlayIcon())
			onStop(s.Name)
		}
	})

	card := widget.NewCard("", "", container.NewHBox(widget.NewLabel(s.Name), layout.NewSpacer(), playBtn))
	card.Resize(fyne.NewSize(150, 80))

	return card
}
