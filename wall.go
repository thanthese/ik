package main

import "github.com/gdamore/tcell"

type Wall struct {
	point
	symbol rune
	style  tcell.Style
}

func NewWall(pt point) *Wall {
	w := &Wall{point: pt,
		symbol: '#',
		style:  tcell.StyleDefault}
	return w
}

func (w *Wall) At() point {
	return w.point
}

func (w *Wall) MoveTo(pt point) {
	w.point = pt
}

func (w *Wall) Glyph() rune {
	return w.symbol
}

func (w *Wall) Style() tcell.Style {
	return w.style
}
