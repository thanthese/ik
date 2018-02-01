package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
)

var (
	upperLeft  = vector{-1, 1}
	upperRight = vector{0, 1}
	right      = vector{1, 0}
	left       = vector{-1, 0}
	lowerLeft  = vector{0, -1}
	lowerRight = vector{1, -1}

	// notice clockwise order
	sixDirs = []vector{upperLeft, upperRight, right, lowerRight, lowerLeft, left}
)

type polarity int

const (
	black polarity = iota
	white
)

type entity interface {
	At() point
	MoveTo(point)
	Glyph() rune
	Style() tcell.Style
}

type point struct{ x, y int }
type vector struct{ xdiff, ydiff int }

func (p point) add(v vector) point {
	return point{
		x: p.x + v.xdiff,
		y: p.y + v.ydiff}
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
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	s.Clear()
	return s
}

func main() {
	s := newScreen()
	defer s.Fini()
	w := newWorld(point{15, 15})
	w.buildScent() // remove
	w.render(s)

	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				return
			case tcell.KeyRune:
				var v vector
				switch ev.Rune() {
				case 'q', 'Q', ' ':
					return
				case 'h', 'H':
					v = left
				case 'l', 'L':
					v = right
				case 'y', 'Y':
					v = upperLeft
				case 'u', 'U':
					v = upperRight
				case 'b', 'B':
					v = lowerLeft
				case 'n', 'N':
					v = lowerRight
				}
				w.moveEntityTo(w.player, w.player.At().add(v))
				w.buildScent()
				w.moveEnemies()
				w.render(s)
			}
		}
	}
}
