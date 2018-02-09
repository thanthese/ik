package main

type Player struct {
	Point
	Polarity
}

func (p *Player) Move(w *World, v Vector) (worked bool) {
	to := p.Add(v)
	if !w.IsEmpty(to) {
		return false
	}
	p.Point = to
	return true
}
