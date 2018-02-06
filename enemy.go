package main

type Enemy struct {
	Polarity Polarity
}

func MoveEnemies(w *World) {
	// 	w.Scent = map[Point]int{w.Player.At(): 0}
	// 	for q, c := []Point{w.Player.At()}, 1; len(q) > 0; c++ {
	// 		r := []Point{}
	// 		for _, p := range q {
	// 			for _, n := range w.ExistingNeighbors(p) {
	// 				if _, ok := w.Scent[n]; !ok {
	// 					w.Scent[n] = c
	// 					r = append(r, n)
	// 				}
	// 			}
	// 		}
	// 		q = r
	// 	}

}
