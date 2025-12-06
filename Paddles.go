package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Paddle struct {
	x, y          int
	color         color.Color
	width, height int
	score         int
}

func NewPaddle(x, y, width, height, score int, color color.Color) *Paddle {
	return &Paddle{
		x:      x,
		y:      y,
		width:  width,
		height: height,
		color:  color,
		score:  score,
	}
}

func (p *Paddle) rectPaddle(screen *ebiten.Image) {
	vector.FillRect(
		screen,
		float32(p.x),
		float32(p.y),
		float32(p.width),
		float32(p.height),
		p.color,
		false,
	)
}

func (p *Paddle) movePaddle1() {
	if ebiten.IsKeyPressed(ebiten.KeyW) && p.y > 10 {
		p.y -= 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) && p.y < screenHeight-p.height-10 {
		p.y += 10
	}
}

func (p *Paddle) movePaddle2() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) && p.y > 10 {
		p.y -= 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) && p.y < screenHeight-p.height-10 {
		p.y += 10
	}
}