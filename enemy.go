package main

import (
	"errors"
)

type Enemy struct {
	Point
	Polarity Polarity
}

func (e *Enemy) At() Point             { return e.Point }
func (e *Enemy) GetPolarity() Polarity { return e.Polarity }
func (p *Enemy) MoveTo(pt Point)       { p.Point = pt }

func (e *Enemy) Move(w *World) (Point, error) {
	min := e.Point
	for _, n := range w.ExistingNeighbors(e.Point) {
		if w.Scent[n] < w.Scent[min] {
			min = n
		}
	}
	if min == e.Point {
		return e.Point, errors.New("no place/reason for enemy to move")
	}
	return min, nil
}
