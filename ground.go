package main

import "github.com/gdamore/tcell"

type Ground struct {
	point
	symbol rune
	style  tcell.Style
}

func NewGround(pt point) *Ground {
	g := &Ground{point: pt,
		symbol: '.',
		style:  tcell.StyleDefault}
	return g
}

func (g *Ground) At() point {
	return g.point
}

func (g *Ground) MoveTo(pt point) {
	g.point = pt
}

func (g *Ground) Glyph() rune {
	return g.symbol
}

func (g *Ground) Style() tcell.Style {
	return g.style
}
