package main

import (
	"errors"

	"github.com/gdamore/tcell"
)

type world struct {
	maxPt   point
	board   [][]entity
	player  *Player
	enemies []*enemy
	scent   map[point]int
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
	for i, m := 2, 8; i < m; i++ {
		w.addEntity(NewWall(point{i, m - i}))
	}
	for i, m := 1, 15; i < m; i++ {
		w.addEntity(NewWall(point{i, m - i}))
	}
	w.addEntity(NewEnemy(point{15, 15}, white))
	w.addEntity(NewEnemy(point{10, 10}, white))
	return w
}

func (w *world) addEntity(e entity) {
	if p, ok := e.(*Player); ok {
		w.player = p
	}
	if b, ok := e.(*enemy); ok {
		w.enemies = append(w.enemies, b)
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

func (w *world) moveEntityTo(e entity, p point) error {
	a := e.At()
	if !w.inBounds(p) {
		return errors.New("movement out of bounds")
	}
	if _, ok := w.board[p.x][p.y].(*Ground); !ok {
		return errors.New("can only move onto empty ground")
	}
	w.board[a.x][a.y] = NewGround(point{a.x, a.y})
	w.board[p.x][p.y] = e
	e.MoveTo(p)
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

func (w *world) moveEnemies() {
	for _, e := range w.enemies {
		p, err := e.Move(w)
		if err == nil {
			w.moveEntityTo(e, p)
		}
	}
}

func (w *world) buildScent() {
	w.scent = map[point]int{w.player.At(): 0}
	for q, c := []point{w.player.At()}, 1; len(q) > 0; c++ {
		r := []point{}
		for _, p := range q {
			for _, n := range w.inBoundsNeighbors(p) {
				switch w.board[n.x][n.y].(type) {
				case *Ground, *enemy:
					if _, ok := w.scent[n]; !ok {
						w.scent[n] = c
						r = append(r, n)
					}
				}
			}
		}
		q = r
	}
}
