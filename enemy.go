package main

import "github.com/gdamore/tcell"

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
