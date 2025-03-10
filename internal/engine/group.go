package engine

import (
	"iter"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
)

// Group represents a group of objects (aka. FlxTypedGroup).
type Group[T any] struct {
	objs []T
}

// NewGroup creates a new [Group].
func NewGroup[T any](objs ...T) *Group[T] {
	return &Group[T]{objs: objs}
}

// Len returns the length of the group i.e. the number of objects in the slice.
func (g *Group[T]) Len() int {
	return len(g.objs)
}

// Cap returns the maximum capacity of the group.
func (g *Group[T]) Cap() int {
	return cap(g.objs)
}

// Add adds new a new object or multiple objects to the group.
func (g *Group[T]) Add(objs ...T) {
	g.objs = append(g.objs, objs...)
}

// Remove removes an object at the index and returns it.
func (g *Group[T]) Remove(index int) T {
	obj := g.objs[index]
	g.objs = slices.Delete(g.objs, index, index+1)
	return obj
}

// Reset wipes and clears the group.
func (g *Group[T]) Reset() {
	g.objs = nil
}

// Get returns an object at the index.
func (g *Group[T]) Get(index int) T {
	return g.objs[index]
}

// Set sets the object at the index.
func (g *Group[T]) Set(index int, obj T) {
	g.objs[index] = obj
}

// Iterate returns an iterator over the objects.
func (g *Group[T]) Iterate() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, obj := range g.objs {
			if !yield(i, obj) {
				return
			}
		}
	}
}

// Update iterates and updates all objects.
func (g *Group[T]) Update() error {
	for i := range g.objs {
		switch obj := any(g.objs[i]).(type) {
		case interface{ Update() }:
			obj.Update()
		case interface{ Update() error }:
			if err := obj.Update(); err != nil {
				return err
			}
			/*
				case interface{ Update(dt float32) }:
					obj.Update(1 / float32(ebiten.TPS()))
				case interface{ Update(dt float32) error }:
					if err := obj.Update(1 / float32(ebiten.TPS())); err != nil {
						return err
					}
				case interface{ Update(dt float64) }:
					obj.Update(1 / float64(ebiten.TPS()))
				case interface{ Update(dt float64) error }:
					if err := obj.Update(1 / float64(ebiten.TPS())); err != nil {
						return err
					}
			*/
		}
	}
	return nil
}

// Draw iterates and draws all objects onto the image.
func (g *Group[T]) Draw(img *ebiten.Image) {
	for i := range g.objs {
		if obj, ok := any(g.objs[i]).(interface{ Draw(*ebiten.Image) }); ok {
			obj.Draw(img)
		}
	}
}
