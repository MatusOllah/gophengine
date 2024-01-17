package gophengine

import (
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type IntroText struct {
	textLock sync.RWMutex
	text     []string
}

func NewIntroText() *IntroText {
	return &IntroText{
		text: []string{},
	}
}

func (it *IntroText) Update(dt float64) error {
	return nil
}

func (it *IntroText) Draw(img *ebiten.Image) {
	it.textLock.RLock()
	for i, s := range it.text {
		ebitenutil.DebugPrintAt(img, s, img.Bounds().Dx()/2, (i*60)+200)
	}
	it.textLock.RUnlock()
}

func (it *IntroText) CreateText(text ...string) {
	it.textLock.Lock()
	it.text = append(it.text, text...)
	it.textLock.Unlock()
}

func (it *IntroText) AddText(text string) {
	it.textLock.Lock()
	it.text = append(it.text, text)
	it.textLock.Unlock()
}

func (it *IntroText) DeleteText() {
	it.textLock.Lock()
	it.text = []string{}
	it.textLock.Unlock()
}
