package main

type Enemy struct {
	Polarity Polarity
}

func MoveEnemies(w *World) {
	scent := buildScent(w)

	es := []Point{}
	for p := range w.Board {
		if w.HasEnemy(p) {
			es = append(es, p)
		}
	}

	for _, e := range es {
		min := e
		for _, n := range w.ExistingNeighbors(e) {
			if w.IsEmpty(n) && scent[n] < scent[min] {
				min = n
			}
		}
		if min != e {
			w.Board[min], w.Board[e] = w.Board[e], w.Board[min]
		}
	}
}

func buildScent(w *World) map[Point]int {
	scent := map[Point]int{w.Player.Point: 0}
	for q, c := []Point{w.Player.Point}, 1; len(q) > 0; c++ {
		q2 := []Point{}
		for _, p := range q {
			for _, n := range w.ExistingNeighbors(p) {
				if _, ok := scent[n]; !ok {
					scent[n] = c
					q2 = append(q2, n)
				}
			}
		}
		q = q2
	}
	return scent
}
