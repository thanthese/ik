package main

import "github.com/gdamore/tcell"

type Player struct {
	BaseEntity
}

func NewPlayer(pt point) *Player {
	p := &Player{}
	p.point = pt
	p.symbol = '@'
	p.style = tcell.StyleDefault
	return p
}
