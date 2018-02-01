package main

import (
	"errors"

	"github.com/gdamore/tcell"
)

type enemy struct {
	BaseEntity
	polarity polarity
}

func NewEnemy(pt point, p polarity) *enemy {
	g := &enemy{}
	g.point = pt
	g.symbol = 'e'
	g.style = tcell.StyleDefault
	g.polarity = p
	return g
}

func (e *enemy) Move(w *world) (point, error) {
	m := e.point
	for _, p := range w.inBoundsNeighbors(e.point) {
		switch w.board[p.x][p.y].(type) {
		case *Ground:
			if w.scent[p] < w.scent[m] {
				m = p
			}
		}
	}
	if m == e.point {
		return m, errors.New("no place/reason for enemy to move")
	}
	return m, nil
}
