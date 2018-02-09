package main

type Player struct {
	Point
	Polarity
	Kills int
}

func (p *Player) Move(w *World, v Vector) (worked bool) {
	from := p.Point
	to := p.Add(v)
	if !w.IsEmpty(to) {
		return false
	}
	p.Point = to

	for _, n1 := range w.ExistingNeighbors(from) {
		for _, n2 := range w.ExistingNeighbors(to) {
			if n1 == n2 {
				if w.HasEnemy(n1) {
					w.Board[n1] = nil
					p.Kills++
				}
			}
		}
	}

	return true
}
