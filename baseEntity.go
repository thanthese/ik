package main

import "github.com/gdamore/tcell"

type BaseEntity struct {
	point
	symbol rune
	style  tcell.Style
}

func (e *BaseEntity) At() point {
	return e.point
}

func (e *BaseEntity) MoveTo(pt point) {
	e.point = pt
}

func (e *BaseEntity) Glyph() rune {
	return e.symbol
}

func (e *BaseEntity) Style() tcell.Style {
	return e.style
}
