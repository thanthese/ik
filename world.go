package main

type World struct {

	// not existant entry: cannot pass there
	// nil entry: empty, passable ground
	// entity there: entity over passable ground
	Board map[Point]*Enemy

	Player *Player
}

func NewWorld() *World {
	w := &World{}
	w.Board = make(map[Point]*Enemy)
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			w.Board[Point{x, y}] = nil
		}
	}
	for i, m := 2, 8; i < m; i++ {
		delete(w.Board, Point{i, m - i})
	}
	w.Board[Point{5, 5}] = &Enemy{Polarity: Black}
	return w
}

func (w *World) Exists(p Point) bool {
	_, ok := w.Board[p]
	return ok
}

func (w *World) ExistingNeighbors(p Point) []Point {
	ps := []Point{}
	for _, v := range SixDirs {
		n := p.add(v)
		if w.Exists(n) {
			ps = append(ps, n)
		}
	}
	return ps
}
