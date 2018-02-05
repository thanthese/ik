package main

type Player struct {
	Point
	Polarity
}

func (p *Player) At() Point             { return p.Point }
func (p *Player) GetPolarity() Polarity { return p.Polarity }
func (p *Player) MoveTo(pt Point)       { p.Point = pt }
