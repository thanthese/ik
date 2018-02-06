package color

import "github.com/gdamore/tcell"

var (

	// solarized color scheme http://ethanschoonover.com/solarized
	// (comments for dark theme)

	Base03  = tcell.NewHexColor(0x002b36) // background
	Base02  = tcell.NewHexColor(0x073642) // background highlights
	Base01  = tcell.NewHexColor(0x586e75) // secondary content
	Base00  = tcell.NewHexColor(0x657b83)
	Base0   = tcell.NewHexColor(0x839496) // body text
	Base1   = tcell.NewHexColor(0x93a1a1) // (optional, emphasized content)
	Base2   = tcell.NewHexColor(0xeee8d5)
	Base3   = tcell.NewHexColor(0xfdf6e3)
	Yellow  = tcell.NewHexColor(0xb58900)
	Orange  = tcell.NewHexColor(0xcb4b16)
	Red     = tcell.NewHexColor(0xdc322f)
	Magenta = tcell.NewHexColor(0xd33682)
	Violet  = tcell.NewHexColor(0x6c71c4)
	Blue    = tcell.NewHexColor(0x268bd2)
	Cyan    = tcell.NewHexColor(0x2aa198)
	Green   = tcell.NewHexColor(0x859900)

	// our usage

	Background = Base00
	Black      = Base03
	White      = Base2

	WhiteStyle = tcell.StyleDefault.
			Foreground(White).
			Background(Background)
	BlackStyle = tcell.StyleDefault.
			Foreground(Black).
			Background(Background)
)
