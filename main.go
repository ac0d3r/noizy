package main

import (
	"embed"

	"noizy/player"
	"noizy/ui"
)

//go:embed assets/*
var assetsFS embed.FS

var (
	AppName    = "Noizy"
	AppVersion = "v0.0.1"
)

func main() {
	player := player.New()
	if err := player.Init(); err != nil {
		panic(err)
	}

	app := ui.NewApp(AppName, player, assetsFS)
	app.Run()
}
