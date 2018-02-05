package main

import (
	"errors"

	"github.com/gdamore/tcell"
	"github.com/thanthese/ik/color"
)

type World struct {

	// not existant entry: cannot pass there
	// nil entry: empty, passable ground
	// entity there: entity over passable ground
	Board map[Point]Entity

	Player  *Player
	Enemies []*Enemy
	Scent   map[Point]int
}

func NewWorld() *World {
	w := &World{}
	w.Board = make(map[Point]Entity)
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			w.Board[Point{x, y}] = nil
		}
	}
	for i, m := 2, 8; i < m; i++ {
		delete(w.Board, Point{i, m - i})
	}
	w.Player = &Player{
		Point:    Point{0, 0},
		Polarity: White}
	w.Enemies = append(w.Enemies,
		&Enemy{
			Point:    Point{5, 5},
			Polarity: Black})
	return w
}

func (w *World) Render(s tcell.Screen) {
	s.Clear()
	for p, e := range w.Board {
		cx := p.X*4 + p.Y*2
		cy := 20 - p.Y
		if e == nil {
			s.SetContent(cx, cy, '.', nil, color.WhiteStyle)
		} else {
			switch e.(type) {
			case *Player:
				s.SetContent(cx, cy, '@', nil, getStyle(e.GetPolarity()))
			case *Enemy:
				s.SetContent(cx, cy, 'e', nil, getStyle(e.GetPolarity()))
			}
		}
	}
	s.Show()
}

func getStyle(p Polarity) tcell.Style {
	if p == White {
		return color.WhiteStyle
	}
	return color.BlackStyle
}

func (w *World) MoveEntityTo(e Entity, p Point) error {
	a := e.At()
	if _, ok := w.Board[p]; !ok {
		return errors.New("movement out of bounds")
	}
	if w.Board[p] != nil {
		return errors.New("can only move onto empty ground")
	}
	w.Board[a] = nil
	w.Board[p] = e
	e.MoveTo(p)
	return nil
}

func (w *World) ExistingNeighbors(p Point) []Point {
	ps := []Point{}
	for _, v := range SixDirs {
		n := p.add(v)
		if _, ok := w.Board[n]; ok {
			ps = append(ps, n)
		}
	}
	return ps
}

func (w *World) MoveEnemies() {
	for _, e := range w.Enemies {
		p, err := e.Move(w)
		if err == nil {
			w.MoveEntityTo(e, p)
		}
	}
}

func (w *World) BuildScent() {
	w.Scent = map[Point]int{w.Player.At(): 0}
	for q, c := []Point{w.Player.At()}, 1; len(q) > 0; c++ {
		r := []Point{}
		for _, p := range q {
			for _, n := range w.ExistingNeighbors(p) {
				if _, ok := w.Scent[n]; !ok {
					w.Scent[n] = c
					r = append(r, n)
				}
			}
		}
		q = r
	}
}
