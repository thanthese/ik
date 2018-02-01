package main

import "github.com/gdamore/tcell"

type Pit struct {
	BaseEntity
}

func NewPit(pt point) *Pit {
	p := &Pit{}
	p.point = pt
	p.symbol = ' '
	p.style = tcell.StyleDefault
	return p
}
