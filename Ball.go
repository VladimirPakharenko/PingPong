package main

import (
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	x, y   int
	radius int
	color  color.Color
	ballX  float64
	ballY  float64
}

func NewBall(x, y, radius int, color color.Color, bX, bY float64) *Ball {
	return &Ball{
		x:      x,
		y:      y,
		radius: radius,
		color:  color,
		ballX:  bX,
		ballY:  bY,
	}
}

func (b *Ball) rectBall(screen *ebiten.Image) {
	vector.FillCircle(
		screen,
		float32(b.x),
		float32(b.y),
		float32(b.radius),
		b.color,
		false,
	)
}

func (b *Ball) moveBall(dir int) {
	if dir > 0 && dir <= 5 {
		b.x += int(b.ballX)
		b.y += int(b.ballY)
	} else if dir > 5 && dir <= 10 {
		b.x += int(b.ballX)
		b.y -= int(b.ballY)
	} else if dir > 10 && dir <= 15 {
		b.x -= int(b.ballX)
		b.y += int(b.ballY)
	} else if dir > 15 && dir <= 20 {
		b.x -= int(b.ballX)
		b.y -= int(b.ballY)
	}
}

func (b *Ball) ChangeX() {
	if b.ballX >= 15 {
		b.ballX = -(b.ballX + float64(rand.IntN(3)-2))
	} else if b.ballX <= -15 {
		b.ballX = -b.ballX + float64(rand.IntN(3)-2)
	} else if b.ballX > -15 && b.ballX < -5 {
		b.ballX = -b.ballX + float64(rand.IntN(5)-2)
	} else if b.ballX > 5 && b.ballX < 15 {
		b.ballX = -b.ballX + float64(rand.IntN(5)-2)
	} else if b.ballX >= 0 && b.ballX <= 5 {
		b.ballX = -(b.ballX + float64(rand.IntN(3)))
	} else if b.ballX < 0 && b.ballX >= -5 {
		b.ballX = -b.ballX + float64(rand.IntN(3))
	}
}

func (b *Ball) colissionWithField() {
	if b.y < b.radius || b.y > screenHeight-b.radius {
		b.ballY = -b.ballY
	}
}