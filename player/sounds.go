package player

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const Ext = ".ogg"

type SoundFile struct {
	Name string
	Path string
}

func (p *Player) SystemSounds() map[string][]SoundFile {
	return p.systemSounds
}

func (p *Player) loadSystemSounds() (err error) {
	slog.Debug("Load system sounds")
	kinds, err := soundsFS.ReadDir("sounds")
	if err != nil {
		return err
	}
	for _, kind := range kinds {
		if !kind.IsDir() {
			continue
		}

		srcs := make([]SoundFile, 0)
		sounds, err := soundsFS.ReadDir("sounds/" + kind.Name())
		if err != nil {
			return err
		}

		for _, sound := range sounds {
			if sound.IsDir() {
				continue
			}

			srcs = append(srcs, SoundFile{
				Name: convPathToName(sound.Name()),
				Path: fmt.Sprintf("sounds/%s/%s", kind.Name(), sound.Name()),
			})
		}
		p.systemSounds[convPathToName(kind.Name())] = srcs
	}

	return nil
}

func (p *Player) findSound(name string) (fs.File, error) {
	path := ""

	for _, sounds := range p.systemSounds {
		for _, sound := range sounds {
			if sound.Name == name {
				path = sound.Path
				break
			}
		}
	}

	if path == "" {
		return nil, errors.New("not found sound")
	}

	return soundsFS.Open(path)
}

func convPathToName(path string) string {
	return cases.Title(language.Dutch).String(
		strings.ReplaceAll(
			strings.Trim(path, filepath.Ext(path)),
			"-", " "),
	)
}
