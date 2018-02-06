package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
	"github.com/thanthese/ik/color"
)

type UI struct {
	tcell.Screen
	viewWidth, viewHeight int
}

func NewUI(width, height int) *UI {
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e = s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	ui := &UI{viewWidth: width, viewHeight: height}
	ui.Screen = s
	ui.SetStyle(color.WhiteStyle)
	ui.Clear()
	return ui
}

func (ui *UI) Write(w *World) {
	ui.Clear()
	for p, e := range w.Board {
		cx := p.X*4 + p.Y*2
		cy := ui.viewHeight - p.Y
		if p == w.Player.Point {
			ui.SetContent(cx, cy, '@', nil, getStyle(w.Player.Polarity))
		} else if e == nil {
			ui.SetContent(cx, cy, '.', nil, color.WhiteStyle)
		} else {
			ui.SetContent(cx, cy, 'e', nil, getStyle(e.Polarity))
		}
	}
	ui.Show()
}

func getStyle(p Polarity) tcell.Style {
	if p == White {
		return color.WhiteStyle
	}
	return color.BlackStyle
}

func (ui *UI) GetInput() (v Vector, quit bool) {
	for {
		switch ev := ui.PollEvent().(type) {
		case *tcell.EventResize:
			ui.Sync()
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				return Vector{}, true
			case tcell.KeyRune:
				switch ev.Rune() {
				case 'q', 'Q', ' ':
					return Vector{}, true
				case 'h', 'H':
					return Left, false
				case 'l', 'L':
					return Right, false
				case 'y', 'Y':
					return UpperLeft, false
				case 'u', 'U':
					return UpperRight, false
				case 'b', 'B':
					return LowerLeft, false
				case 'n', 'N':
					return LowerRight, false
				}
			}
		}
	}
}
