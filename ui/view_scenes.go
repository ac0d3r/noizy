package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func buildScenesView() *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("🌅 这里是自定义场景列表"),
	)
}
