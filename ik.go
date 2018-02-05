package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
	"github.com/thanthese/ik/color"
)

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

type Polarity bool

const (
	Black Polarity = true
	White Polarity = false
)

type Point struct{ X, Y int }
type Vector struct{ Xdiff, Ydiff int }

func (p Point) add(v Vector) Point {
	return Point{
		X: p.X + v.Xdiff,
		Y: p.Y + v.Ydiff}
}

type Entity interface {
	At() Point
	GetPolarity() Polarity
	MoveTo(Point)
}

func newScreen() tcell.Screen {
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e = s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	s.SetStyle(tcell.StyleDefault.
		Foreground(color.White).
		Background(color.Background))
	s.Clear()
	return s
}

func main() {
	s := newScreen()
	defer s.Fini()
	w := NewWorld()
	w.BuildScent() // remove
	w.Render(s)

	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				return
			case tcell.KeyRune:
				var v Vector
				switch ev.Rune() {
				case 'q', 'Q', ' ':
					return
				case 'h', 'H':
					v = Left
				case 'l', 'L':
					v = Right
				case 'y', 'Y':
					v = UpperLeft
				case 'u', 'U':
					v = UpperRight
				case 'b', 'B':
					v = LowerLeft
				case 'n', 'N':
					v = LowerRight
				}
				w.MoveEntityTo(w.Player, w.Player.At().add(v))
				w.BuildScent()
				w.MoveEnemies()
				w.Render(s)
			}
		}
	}
}
