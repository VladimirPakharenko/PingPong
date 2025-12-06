package main

import (
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	screenFon1 *Screen2
	screenFon  *Screen1
	paddle1    *Paddle
	paddle2    *Paddle
	ball       *Ball
	timer      int
}

func (g *Game) colissionWithGates() {
	if g.ball.x < 0+g.ball.radius {
		g.paddle2.score++
		g.screenFon.gameOver = true
	} else if g.ball.x > screenWidth-g.ball.radius {
		g.paddle1.score++
		g.screenFon.gameOver = true
	}
} //Попадание в ворота

func (g *Game) resetGame() {
	g.paddle1 = NewPaddle(
		paddleX,
		screenHeight/2-paddleHeight/2,
		paddleWidth, paddleHeight,
		g.paddle1.score,
		color.RGBA{31, 0, 204, 255},
	)
	g.paddle2 = NewPaddle(
		screenWidth-(paddleX+paddleWidth),
		screenHeight/2-paddleHeight/2,
		paddleWidth, paddleHeight,
		g.paddle2.score,
		color.RGBA{204, 0, 31, 255},
	)
	g.ball = NewBall(
		screenWidth/2,
		screenHeight/2,
		ballRadius,
		color.RGBA{255, 208, 0, 255},
		float64(ballX),
		float64(ballY),
	)
	g.ball.ballX = 10
	g.ball.ballY = 10
	g.screenFon.gameOver = false
	g.screenFon.direction = rand.IntN(20) + 1
} //Рестарт игры

func (g *Game) fullscreen() {
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
} //Полноэкранный режим

// func (g *Game) menuESC() {
// 	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
// 		flag += 1
// 	}
// }

func (g *Game) gameOver() error {
	if g.screenFon.gameOver {
		timer++
		if timer >= 30 {
			timer = 0
			g.resetGame()
			return nil
		}
		return nil
	}
	return nil
} //Конец игры

func (g *Game) collisionWithRacket() {
	LeftBall := g.ball.x - g.ball.radius
	RightBall := g.ball.x + g.ball.radius

	UpPad1 := g.paddle1.y
	DownPad1 := g.paddle1.y + g.paddle1.height
	RightPad1 := g.paddle1.x + g.paddle1.width

	UpPad2 := g.paddle2.y
	DownPad2 := g.paddle2.y + g.paddle2.height
	LeftPad2 := g.paddle2.x

	if LeftBall <= RightPad1 && (g.ball.y >= UpPad1 && g.ball.y <= DownPad1) {
		g.ball.x += 10
		g.ball.ChangeX()
	}
	if RightBall >= LeftPad2 && (g.ball.y >= UpPad2 && g.ball.y <= DownPad2) {
		g.ball.x -= 10
		g.ball.ChangeX()
	}
} // Коллизии мяча с ракетками
