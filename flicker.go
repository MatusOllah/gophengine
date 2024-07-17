package gophengine

import (
	"log/slog"
	"time"
)

type Flicker struct {
	Sprite             *Sprite
	dur                time.Duration
	interval           time.Duration
	startTime          time.Time
	flickerTimer       time.Time
	flicker            bool
	OnCompleteCallback func()
}

func NewFlicker(spr *Sprite, dur time.Duration, interval time.Duration) *Flicker {
	return &Flicker{
		Sprite:             spr,
		dur:                dur,
		interval:           interval,
		flicker:            false,
		OnCompleteCallback: func() {},
	}
}

func (f *Flicker) Flicker() {
	slog.Debug("[Flicker] starting flicker")
	f.flicker = true
	f.startTime = time.Now()
	f.flickerTimer = time.Now()
}

func (f *Flicker) Update() {
	if !f.flicker {
		return
	}

	// Check if the total flicker duration has passed
	if time.Since(f.startTime) < f.dur {
		// Handle the flicker logic
		if time.Since(f.flickerTimer) > f.interval {
			slog.Debug("[Flicker] flickering", "flickerTimer", f.flickerTimer, "startTime", f.startTime, "visible", f.Sprite.Visible)
			f.Sprite.Visible = !f.Sprite.Visible
			f.flickerTimer = time.Now()
		}
	} else {
		f.OnCompleteCallback()
		f.Sprite.Visible = false
		f.flicker = false
	}
}
