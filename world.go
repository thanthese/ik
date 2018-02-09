package main

type World struct {
	Board  map[Point]*Enemy
	Player *Player
}

func NewLvl1() *World {
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
	w.Board[Point{6, 6}] = &Enemy{Polarity: Black}
	w.Player = &Player{
		Point:    Point{0, 0},
		Polarity: White}
	return w
}

func (w *World) Exists(p Point) bool {
	_, ok := w.Board[p]
	return ok
}

func (w *World) IsEmpty(p Point) bool {
	return w.Exists(p) && w.Board[p] == nil
}

func (w *World) HasEnemy(p Point) bool {
	return w.Exists(p) && w.Board[p] != nil
}

func (w *World) ExistingNeighbors(p Point) []Point {
	ps := []Point{}
	for _, v := range SixDirs {
		n := p.Add(v)
		if w.Exists(n) {
			ps = append(ps, n)
		}
	}
	return ps
}
