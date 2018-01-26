package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/gdamore/tcell"
)

const boardXDim = 7
const boardYDim = 5

var (
	upperLeft  = vector{-1, 1}
	upperRight = vector{0, 1}
	right      = vector{1, 0}
	left       = vector{-1, 0}
	lowerLeft  = vector{0, -1}
	lowerRight = vector{1, -1}
	neighbors  = []vector{upperLeft, upperRight, right, left, lowerLeft, lowerRight}
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

type world struct {
	screen tcell.Screen
	board  [][]entity
	player *Player
}

func newWorld(screen tcell.Screen) *world {
	w := &world{}
	w.screen = screen
	w.board = make([][]entity, boardYDim)
	for x := range w.board {
		w.board[x] = make([]entity, boardXDim)
		for y := range w.board[x] {
			w.board[x][y] = NewGround(point{x, y})
		}
	}
	p := NewPlayer(point{3, 3})
	w.board[3][3] = p
	w.player = p
	return w
}

func (w *world) render() {
	w.screen.Clear()
	for x := 0; x < boardYDim; x++ {
		for y := 0; y < boardXDim; y++ {
			cx := x*4 + y*2
			cy := boardXDim - y - 1
			en := w.board[x][y]
			w.screen.SetContent(cx, cy, en.Glyph(), nil, en.Style())
		}
	}
	w.screen.Show()
}

func (w *world) moveEntityBy(e entity, v vector) error {
	a := e.At()
	b := e.At().add(v)
	if !inBounds(b) {
		return errors.New("movement out of bounds")
	}
	w.board[a.x][a.y] = NewGround(point{a.x, a.y})
	w.board[b.x][b.y] = e
	e.MoveTo(b)
	return nil
}

func inBounds(p point) bool {
	return 0 <= p.x && p.x < boardYDim && 0 <= p.y && p.y < boardXDim
}

func main() {
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

	w := newWorld(s)
	w.render()

	quit := make(chan struct{})
	go func() {
		for {
			switch ev := s.PollEvent().(type) {
			case *tcell.EventResize:
				s.Sync()
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyCtrlC:
					close(quit)
					return
				case tcell.KeyRune:
					var v vector
					switch ev.Rune() {
					case 'q', 'Q':
						close(quit)
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
					w.moveEntityBy(w.player, v)
					w.render()
				}
			}
		}
	}()
	<-quit
	s.Fini()
}
