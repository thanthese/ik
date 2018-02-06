package main

type Polarity bool

const (
	Black Polarity = true
	White Polarity = false
)

type Point struct{ X, Y int }

func (p Point) add(v Vector) Point {
	return Point{
		X: p.X + v.Xdiff,
		Y: p.Y + v.Ydiff}
}

type Vector struct{ Xdiff, Ydiff int }

var (
	UpperLeft  = Vector{-1, 1}
	UpperRight = Vector{0, 1}
	Right      = Vector{1, 0}
	Left       = Vector{-1, 0}
	LowerLeft  = Vector{0, -1}
	LowerRight = Vector{1, -1}

	// notice clockwise order
	SixDirs = []Vector{UpperLeft, UpperRight, Right, LowerRight, LowerLeft, Left}
)

func main() {
	ui := NewUI(10, 10)
	defer ui.Fini()
	w := NewWorld()
	p := &Player{
		Point:    Point{0, 0},
		Polarity: White}
	w.Player = p
	w.Board[p.Point] = nil
	ui.Write(w)

	for {
		ui.Write(w)
		v, quit := ui.GetInput()
		if quit {
			return
		}
		p.Move(w, v)
		MoveEnemies(w)
	}
}
