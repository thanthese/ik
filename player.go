package main

type Player struct {
	Point
	Polarity
}

func (p *Player) Move(w *World, v Vector) (worked bool) {
	b := p.add(v)
	if !w.Exists(b) {
		return false
	}
	p.Point = b
	return true
}
