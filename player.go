package main

import "github.com/gdamore/tcell"

type Player struct {
	point
	symbol rune
	style  tcell.Style
}

func NewPlayer(pt point) *Player {
	return &Player{point: pt,
		symbol: '@',
		style:  tcell.StyleDefault}
}

func (p *Player) At() point {
	return p.point
}

func (p *Player) MoveTo(pt point) {
	p.point = pt
}

func (p *Player) Glyph() rune {
	return p.symbol
}

func (p *Player) Style() tcell.Style {
	return p.style
}
