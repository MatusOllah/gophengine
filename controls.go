package gophengine

import (
	"github.com/MatusOllah/gophengine/internal/config"
	"github.com/hajimehoshi/ebiten/v2"
)

type Controls struct {
	Up, Down, Left, Right, Accept, Back, Pause, Reset, Fullscreen ebiten.Key
}

func GetControlsFromConfig(cfg *config.Config) (*Controls, error) {
	ctl := &Controls{}

	// Up
	up, err := cfg.Get("Controls.Up")
	if err != nil {
		return nil, err
	}
	if err := ctl.Up.UnmarshalText(up.([]byte)); err != nil {
		return nil, err
	}

	// Down
	down, err := cfg.Get("Controls.Down")
	if err != nil {
		return nil, err
	}
	if err := ctl.Down.UnmarshalText(down.([]byte)); err != nil {
		return nil, err
	}

	// Left
	left, err := cfg.Get("Controls.Left")
	if err != nil {
		return nil, err
	}
	if err := ctl.Left.UnmarshalText(left.([]byte)); err != nil {
		return nil, err
	}

	// Right
	right, err := cfg.Get("Controls.Right")
	if err != nil {
		return nil, err
	}
	if err := ctl.Right.UnmarshalText(right.([]byte)); err != nil {
		return nil, err
	}

	// Accept
	accept, err := cfg.Get("Controls.Accept")
	if err != nil {
		return nil, err
	}
	if err := ctl.Accept.UnmarshalText(accept.([]byte)); err != nil {
		return nil, err
	}

	// Back
	back, err := cfg.Get("Controls.Back")
	if err != nil {
		return nil, err
	}
	if err := ctl.Back.UnmarshalText(back.([]byte)); err != nil {
		return nil, err
	}

	// Pause
	pause, err := cfg.Get("Controls.Pause")
	if err != nil {
		return nil, err
	}
	if err := ctl.Pause.UnmarshalText(pause.([]byte)); err != nil {
		return nil, err
	}

	// Reset
	reset, err := cfg.Get("Controls.Reset")
	if err != nil {
		return nil, err
	}
	if err := ctl.Reset.UnmarshalText(reset.([]byte)); err != nil {
		return nil, err
	}

	// Fullscreen
	fullscreen, err := cfg.Get("Controls.Fullscreen")
	if err != nil {
		return nil, err
	}
	if err := ctl.Fullscreen.UnmarshalText(fullscreen.([]byte)); err != nil {
		return nil, err
	}

	return ctl, nil
}
