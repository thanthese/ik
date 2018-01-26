package main

import (
	"errors"
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
	sixDirs    = []vector{upperLeft, upperRight, right, left, lowerLeft, lowerRight}
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
	maxPt  point
	board  [][]entity
	player *Player
}

func newWorld(maxPt point) *world {
	w := &world{}
	w.maxPt = maxPt
	w.board = make([][]entity, maxPt.x+1)
	for x := 0; x <= w.maxPt.x; x++ {
		w.board[x] = make([]entity, maxPt.y+1)
		for y := 0; y <= w.maxPt.y; y++ {
			w.board[x][y] = NewGround(point{x, y})
		}
	}
	w.addEntity(NewPlayer(point{0, 0}))
	for _, p := range w.inBoundsNeighbors(point{3, 3}) {
		w.addEntity(NewWall(p))
	}
	for _, p := range w.inBoundsNeighbors(point{7, 7}) {
		w.addEntity(NewWall(p))
	}
	return w
}

func (w *world) addEntity(e entity) {
	if p, ok := e.(*Player); ok {
		w.player = p
	}
	w.board[e.At().x][e.At().y] = e
}

func (w *world) render(s tcell.Screen) {
	s.Clear()
	for x := 0; x <= w.maxPt.x; x++ {
		for y := 0; y <= w.maxPt.y; y++ {
			cx := x*4 + y*2
			cy := w.maxPt.y - y
			en := w.board[x][y]
			s.SetContent(cx, cy, en.Glyph(), nil, en.Style())
		}
	}
	s.Show()
}

func (w *world) moveEntityBy(e entity, v vector) error {
	a := e.At()
	b := e.At().add(v)
	if !w.inBounds(b) {
		return errors.New("movement out of bounds")
	}
	if _, ok := w.board[b.x][b.y].(*Ground); !ok {
		return errors.New("can only move onto empty ground")
	}
	w.board[a.x][a.y] = NewGround(point{a.x, a.y})
	w.board[b.x][b.y] = e
	e.MoveTo(b)
	return nil
}

func (w *world) inBounds(p point) bool {
	return 0 <= p.x && p.x <= w.maxPt.x && 0 <= p.y && p.y <= w.maxPt.y
}

func (w *world) inBoundsNeighbors(p point) []point {
	ps := []point{}
	for _, v := range sixDirs {
		n := p.add(v)
		if w.inBounds(n) {
			ps = append(ps, n)
		}
	}
	return ps
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
	defer s.Fini()
	s.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	s.Clear()

	w := newWorld(point{15, 15})
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
				case 'q', 'Q':
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
				w.render(s)
			}
		}
	}
}
