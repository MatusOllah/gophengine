// original: https://github.com/yohamta/ganim8

package anim

import (
	"image"
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const Dur24FPS time.Duration = 41 * time.Millisecond

var intervalMatcher = *regexp.MustCompile("^([0-9]+)-([0-9]+)$")

func parseInterval(val any) (int, int, int) {
	switch v := val.(type) {
	case int:
		return v, v, 1
	case float64:
		return int(v), int(v), 1
	case string:
		if n, err := strconv.Atoi(v); err == nil {
			return n, n, 1
		}
		matches := intervalMatcher.FindStringSubmatch(strings.TrimSpace(v))
		if len(matches) != 3 {
			log.Fatalf("Could not parse interval from %s", v)
		}
		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])
		if min > max {
			return min, max, -1
		} else {
			return min, max, 1
		}
	default:
		log.Fatalf("Could not parse interval from %v", val)
	}
	panic("unreachable")
}

var DefaultDelta = time.Millisecond * 16

func parseDurations(durations any, frameCount int) []time.Duration {
	result := make([]time.Duration, frameCount)
	switch val := durations.(type) {
	case time.Duration:
		for i := 0; i < frameCount; i++ {
			result[i] = val
		}
	case []time.Duration:
		for i := range val {
			result[i] = val[i]
		}
	case []any:
		for i := range val {
			result[i] = parseDurationValue(val[i])
		}
	case map[string]time.Duration:
		for key, duration := range val {
			min, max, step := parseInterval(key)
			for i := min; i <= max; i += step {
				result[i-1] = duration
			}
		}
	case map[string]any:
		for key, duration := range val {
			min, max, step := parseInterval(key)
			for i := min; i <= max; i += step {
				result[i-1] = parseDurationValue(duration)
			}
		}
	case any:
		for i := 0; i < frameCount; i++ {
			result[i] = parseDurationValue(val)
		}
	default:
		log.Fatalf("failed to parse durations: type=%T val=%+v", durations, durations)
	}
	return result
}

func parseDurationValue(value any) time.Duration {
	switch val := value.(type) {
	case time.Duration:
		return val
	case int:
		return time.Millisecond * time.Duration(val)
	case float64:
		return time.Millisecond * time.Duration(val)
	default:
		log.Fatalf("failed to parse duration value: %+v", value)
	}
	return 0
}

func parseIntervals(durations []time.Duration) ([]time.Duration, time.Duration) {
	result := []time.Duration{0}
	var time time.Duration = 0
	for _, v := range durations {
		time += v
		result = append(result, time)
	}
	return result, time
}

// Status represents the animation status.
type Status int

const (
	Playing = iota
	Paused
)

// Animation represents an animation created from specified frames
// and an *ebiten.Image
type Animation struct {
	frames        []*ebiten.Image
	position      int
	timer         time.Duration
	durations     []time.Duration
	intervals     []time.Duration
	totalDuration time.Duration
	onLoop        OnLoop
	status        Status
}

// OnLoop is callback function which representing
// one of the animation methods.
// it will be called every time an animation "loops".
//
// It will have two parameters: the animation instance,
// and how many loops have been elapsed.
//
// The value would be Nop (No operation) if there's nothing
// to do except for looping the animation.
//
// The most usual value (apart from none) is the string 'pauseAtEnd'.
// It will make the animation loop once and then pause
// and stop on the last frame.
type OnLoop func(anim *Animation, loops int)

// Nop does nothing.
func Nop(anim *Animation, loops int) {}

// Pause pauses the animation on loop finished.
func Pause(anim *Animation, loops int) {
	anim.Pause()
}

// PauseAtEnd pauses the animation and set the position to
// the last frame.
func PauseAtEnd(anim *Animation, loops int) {
	anim.PauseAtEnd()
}

// PauseAtStart pauses the animation and set the position to
// the first frame.
func PauseAtStart(anim *Animation, loops int) {
	anim.PauseAtStart()
}

// NewAnimation returns a new animation object
//
// durations is a time.Duration or a []time.Duration or
// a map[string]time.Duration.
// When it's a time.Duration, it represents the duration of
// all frames in the animation.
// When it's a []time.Duration, it can represent different
// durations for different frames.
// You can specify durations for all frames individually,
// like this: []time.Duration { 100 * time.Millisecond,
// 100 * time.Millisecond } or you can specify durations for
// ranges of frames: map[string]time.Duration { "1-2":
// 100 * time.Millisecond, "3-5": 200 * time.Millisecond }.
func NewAnimation(frames []*ebiten.Image, durations any, onLoop ...OnLoop) *Animation {
	_durations := parseDurations(durations, len(frames))
	intervals, totalDuration := parseIntervals(_durations)
	ol := Nop
	if len(onLoop) > 0 {
		ol = onLoop[0]
	}
	anim := &Animation{
		frames:        frames,
		position:      0,
		timer:         0,
		durations:     _durations,
		intervals:     intervals,
		totalDuration: totalDuration,
		onLoop:        ol,
		status:        Playing,
	}
	return anim
}

// Clone return a copied animation object.
func (anim *Animation) Clone() *Animation {
	new := *anim
	return &new
}

// SetOnLoop sets the callback function which representing
func (anim *Animation) SetOnLoop(onLoop OnLoop) {
	anim.onLoop = onLoop
}

func (anim *Animation) IsEnd() bool {
	if anim.status == Paused && anim.position == len(anim.frames)-1 {
		return true
	}
	return false
}

func seekFrameIndex(intervals []time.Duration, timer time.Duration) int {
	high, low, i := len(intervals)-2, 0, 0
	for low <= high {
		i = (low + high) / 2
		if timer >= intervals[i+1] {
			low = i + 1
		} else if timer < intervals[i] {
			high = i - 1
		} else {
			return i
		}
	}
	return i
}

// Update updates the animation.
func (anim *Animation) Update() {
	anim.UpdateWithDelta(DefaultDelta)
}

// UpdateWithDelta updates the animation with the specified delta.
func (anim *Animation) UpdateWithDelta(elapsedTime time.Duration) {
	if anim.status != Playing || len(anim.frames) <= 1 {
		return
	}
	anim.timer += elapsedTime
	loops := anim.timer / anim.totalDuration
	if loops != 0 {
		anim.timer = anim.timer - anim.totalDuration*loops
		(anim.onLoop)(anim, int(loops))
	}
	anim.position = seekFrameIndex(anim.intervals, anim.timer)
}

// SetDurations sets the durations of the animation.
func (anim *Animation) SetDurations(durations any) {
	_durations := parseDurations(durations, len(anim.frames))
	anim.durations = _durations
	anim.intervals, anim.totalDuration = parseIntervals(_durations)
	anim.timer = 0
}

// Status returns the status of the animation.
func (anim *Animation) Status() Status {
	return anim.status
}

// Pause pauses the animation.
func (anim *Animation) Pause() {
	anim.status = Paused
}

// Position returns the current position of the frame.
// The position counts from 1 (not 0).
func (anim *Animation) Position() int {
	return anim.position + 1
}

// Duration returns the current durations of each frames.
func (anim *Animation) Durations() []time.Duration {
	return anim.durations
}

// TotalDuration returns the total duration of the animation.
func (anim *Animation) TotalDuration() time.Duration {
	return anim.totalDuration
}

// Timer returns the current accumulated times of current frame.
func (anim *Animation) Timer() time.Duration {
	return anim.timer
}

// GoToFrame sets the position of the animation and
// sets the timer at the start of the frame.
func (anim *Animation) GoToFrame(position int) {
	anim.position = position - 1
	anim.timer = anim.intervals[anim.position]
}

// PauseAtEnd pauses the animation and set the position
// to the last frame.
func (anim *Animation) PauseAtEnd() {
	anim.position = len(anim.frames) - 1
	anim.timer = anim.totalDuration
	anim.Pause()
}

// PauseAtStart pauses the animation and set the position
// to the first frame.
func (anim *Animation) PauseAtStart() {
	anim.position = 0
	anim.timer = 0
	anim.status = Paused
}

// Resume resumes the animation
func (anim *Animation) Resume() {
	anim.status = Playing
}

// Draw draws the animation with the specified option parameters.
func (anim *Animation) Draw(screen *ebiten.Image, pt image.Point) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(pt.X), float64(pt.Y))
	screen.DrawImage(anim.frames[anim.position], opts)
}

// Frames returns a copy of the frames slice.
func (anim *Animation) Frames() []*ebiten.Image {
	return slices.Clone(anim.frames)
}
