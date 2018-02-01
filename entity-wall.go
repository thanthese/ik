package main

import "github.com/gdamore/tcell"

type Wall struct {
	BaseEntity
}

func NewWall(pt point) *Wall {
	w := &Wall{}
	w.point = pt
	w.symbol = '#'
	w.style = tcell.StyleDefault
	return w
}
