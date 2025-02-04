package gophengine

import (
	"time"
)

type Flicker struct {
	Sprite             *Sprite
	dur                time.Duration
	interval           time.Duration
	startTime          time.Time
	flickerTimer       time.Time
	flicker            bool
	OnCompleteCallback func() error
}

func NewFlicker(spr *Sprite, dur time.Duration, interval time.Duration) *Flicker {
	return &Flicker{
		Sprite:             spr,
		dur:                dur,
		interval:           interval,
		flicker:            false,
		OnCompleteCallback: func() error { return nil },
	}
}

func (f *Flicker) Flicker() {
	f.flicker = true
	f.startTime = time.Now()
	f.flickerTimer = time.Now()
}

func (f *Flicker) Update() error {
	if !f.flicker {
		return nil
	}

	// Check if the total flicker duration has passed
	if time.Since(f.startTime) < f.dur {
		// Handle the flicker logic
		if time.Since(f.flickerTimer) > f.interval {
			f.Sprite.Visible = !f.Sprite.Visible
			f.flickerTimer = time.Now()
		}
	} else {
		f.Sprite.Visible = false
		f.flicker = false
		return f.OnCompleteCallback()
	}

	return nil
}
