package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Screen2 struct {
	backgroundColor color.Color
}

func NewScreen2(b color.Color) *Screen2 {
	return &Screen2{
		backgroundColor: b,
	}
}

func (s *Screen2) img1(screen *ebiten.Image) {
	screen.Fill(s.backgroundColor)
}
