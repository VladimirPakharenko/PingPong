package main

import (
	"bytes"
	"embed"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Screen1 struct {
	backgroundColor color.Color
	direction       int
	gameOver        bool
}

func NewScreen1(b color.Color, d int, over bool) *Screen1 {
	return &Screen1{
		backgroundColor: b,
		direction:       d,
		gameOver:        over,
	}
}

func (s *Screen1) img1(screen *ebiten.Image) {
	screen.Fill(s.backgroundColor)
}

func (s *Screen1) img2(screen *ebiten.Image) {
	vector.FillCircle(
		screen,
		float32(screenWidth/2),
		float32(screenHeight/2),
		float32(250),
		color.RGBA{255, 255, 255, 255},
		false,
	)
	vector.FillCircle(
		screen,
		float32(screenWidth/2),
		float32(screenHeight/2),
		float32(240),
		s.backgroundColor,
		false,
	)
	vector.FillCircle(
		screen,
		float32(screenWidth/2),
		float32(screenHeight/2),
		float32(30),
		color.RGBA{255, 255, 255, 255},
		false,
	)
	for i := 0; i <= 16; i++ {
		if i%2 == 0 {
			vector.FillRect(
				screen,
				float32(screenWidth/2-5),
				float32(0+55*i),
				float32(10),
				float32(50),
				color.RGBA{255, 255, 255, 255},
				false,
			)
		}
	}
}

func (s *Screen1) img3(screen *ebiten.Image) {
	vector.FillRect(
		screen,
		float32(80),
		float32(0),
		float32(20),
		float32(screenHeight),
		color.RGBA{255, 255, 255, 255},
		false,
	)
	vector.FillRect(
		screen,
		float32(1500),
		float32(0),
		float32(20),
		float32(screenHeight),
		color.RGBA{255, 255, 255, 255},
		false,
	)
}

//go:embed Y224.otf
var fontFile embed.FS

func font() *text.GoTextFace {
	fontData, err := fontFile.ReadFile("Y224.otf")
	if err != nil {
		log.Fatal(err)
	}

	tt, err := text.NewGoTextFaceSource(bytes.NewBuffer(fontData))
	if err != nil {
		log.Fatal(err)
	}

	face := &text.GoTextFace{
		Source: tt,
		Size:   16,
	}
	return face
}

func (s *Screen1) debug1(screen *ebiten.Image) {
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("FPS: %0.2f\n", ebiten.ActualFPS()),
	)
}

func (s *Screen1) debug2(screen *ebiten.Image, opt *text.DrawOptions, score int) {
	text.Draw(
		screen,
		fmt.Sprintf("first player score: %d", score),
		font(),
		opt,
	)
}

func (s *Screen1) debug3(screen *ebiten.Image, opt *text.DrawOptions, bX, bY int) {
	text.Draw(
		screen,
		fmt.Sprintf("coord ball x:y %d:%d", bX, bY),
		font(),
		opt,
	)
}
