package main

import "github.com/gdamore/tcell"

type Ground struct {
	BaseEntity
}

func NewGround(pt point) *Ground {
	g := &Ground{}
	g.point = pt
	g.symbol = '.'
	g.style = tcell.StyleDefault
	return g
}
