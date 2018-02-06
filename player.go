package main

type Player struct {
	Point
	Polarity
}

func (p *Player) Move(w *World, v Vector) (worked bool) {
	b := p.Add(v)
	if !w.IsEmpty(b) {
		return false
	}
	p.Point = b
	return true
}
