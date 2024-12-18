package audio

import (
	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/effects"
)

// MixerChannel is a mixer channel with adjustable volume.
// It is a wrapper around *beep.Mixer and *effects.Volume.
type MixerChannel struct {
	beepMixer *beep.Mixer
	volume    *effects.Volume
}

// NewMixerChannel creates a new [MixerChannel].
func NewMixerChannel() *MixerChannel {
	ch := &MixerChannel{}
	ch.beepMixer = &beep.Mixer{}
	ch.volume = &effects.Volume{
		Streamer: ch.beepMixer,
		Base:     2,
		Volume:   0,
		Silent:   false,
	}
	return ch
}

// Add adds Streamers to the wrapped Mixer.
func (ch *MixerChannel) Add(s ...beep.Streamer) {
	ch.beepMixer.Add(s...)
}

// Clear clears all Streamers from the wrapped Mixer.
func (ch *MixerChannel) Clear() {
	ch.beepMixer.Clear()
}

// KeepAlive configures the wrapped Mixer to either keep playing silence when all its Streamers have
// drained (keepAlive == true) or stop playing (keepAlive == false).
func (ch *MixerChannel) KeepAlive(keepAlive bool) {
	ch.beepMixer.KeepAlive(keepAlive)
}

// Len returns the number of Streamers currently playing in the wrapped Mixer.
func (ch *MixerChannel) Len() int {
	return ch.beepMixer.Len()
}

// Stream streams the wrapped Volume.
func (ch *MixerChannel) Stream(samples [][2]float64) (n int, ok bool) {
	return ch.volume.Stream(samples)
}

// Err propagates the wrapped Volume's errors.
func (ch *MixerChannel) Err() error {
	return ch.volume.Err()
}

// Volume returns the current volume level.
func (ch *MixerChannel) Volume() float64 {
	return ch.volume.Volume
}

// SetVolume sets the volume level. If the volume is lower or equal to -10, the channel is muted (set to silent).
func (ch *MixerChannel) SetVolume(vol float64) {
	ch.volume.Volume = vol
	if vol <= -10 {
		ch.volume.Silent = true
	} else {
		ch.volume.Silent = false
	}
}

// Mixer is a collection of mixer channels, including a master, music and SFX channel.
type Mixer struct {
	Master *MixerChannel
	SFX    *MixerChannel
	Music  *MixerChannel
	Extra  map[string]*MixerChannel
}

// NewMixer creates a new [Mixer].
func NewMixer() *Mixer {
	m := &Mixer{
		Master: NewMixerChannel(),
		SFX:    NewMixerChannel(),
		Music:  NewMixerChannel(),
		Extra:  make(map[string]*MixerChannel),
	}

	m.Master.Add(m.SFX)
	m.Master.Add(m.Music)

	return m
}

// AddChannel adds an extra channel to the mixer and assigns it to the master channel.
func (m *Mixer) AddChannel(name string, ch *MixerChannel) {
	m.Extra[name] = ch
	m.Master.Add(ch)
}

// Stream streams the master channel.
func (m *Mixer) Stream(samples [][2]float64) (n int, ok bool) {
	return m.Master.Stream(samples)
}

// Err propagates the master channel's errors.
func (m *Mixer) Err() error {
	return m.Master.Err()
}
