package player

import (
	"embed"
	"fmt"
	"sync"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/vorbis"
)

//go:embed sounds/*
var soundsFS embed.FS

type Player struct {
	mux          sync.RWMutex
	running      map[string]*beep.Ctrl
	systemSounds map[string][]SoundFile
}

func New() *Player {
	return &Player{
		systemSounds: make(map[string][]SoundFile),
		running:      make(map[string]*beep.Ctrl),
	}
}

const bufSize = 44100

func (p *Player) Init() error {
	if err := p.loadSystemSounds(); err != nil {
		return err
	}

	// setup speaker
	fs, err := p.findSound("Brook")
	if err != nil {
		return err
	}

	streamer, format, err := vorbis.Decode(fs)
	if err != nil {
		return err
	}
	defer streamer.Close()

	return speaker.Init(format.SampleRate, bufSize)
}

func (p *Player) Play(name string) error {
	p.mux.Lock()
	defer p.mux.Unlock()

	if ctrl, exist := p.running[name]; exist {
		if ctrl.Paused {
			ctrl.Paused = false
		}
		return nil
	}

	fs, err := p.findSound(name)
	if err != nil {
		return err
	}

	streamer, _, err := vorbis.Decode(fs)
	if err != nil {
		return err
	}

	// TODO
	// streamer := &effects.Volume{
	// 	Streamer: beep.Loop(-1, streamer),
	// 	Base:     2,     // 对数底（通常为2）
	// 	Volume:   -2,    // 分贝单位，负数表示降低音量
	// 	Silent:   false, // 是否静音
	// }

	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
	p.running[name] = ctrl

	speaker.Play(ctrl)
	return nil
}

func (p *Player) Stop(name string) error {
	p.mux.RLock()
	defer p.mux.RUnlock()

	if ctrl, exist := p.running[name]; exist {
		if !ctrl.Paused {
			ctrl.Paused = true
		}
		return nil
	}

	return nil
}

func (p *Player) StopAll() error {
	p.mux.RLock()
	defer p.mux.RUnlock()

	for name, ctrl := range p.running {
		fmt.Println(name, ctrl.Paused)
		if !ctrl.Paused {
			ctrl.Paused = true
		}
		fmt.Println(name, ctrl.Paused)
		return nil
	}
	return nil
}
