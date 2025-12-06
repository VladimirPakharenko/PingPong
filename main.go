package main

import (
	"image/color"
	"log"

	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// Движения
func (g *Game) Update() error {
	g.fullscreen() //полноэкранный режим по кнопке F
	g.gameOver()   //Рестарт Игры
	if !g.screenFon.gameOver {
		g.ball.moveBall(g.screenFon.direction) //Движение мяча
		g.paddle1.movePaddle1()                //Движение первой ракетки
		g.paddle2.movePaddle2()                //Движение второй ракетки
		g.ball.colissionWithField()            //Коллизии мяча с границами поля
		g.collisionWithRacket()                //Коллизии мяча с ракетками
		g.colissionWithGates()                 //коллизии мяча с воротами
	}
	return nil
}

// Отрисовка
func (g *Game) Draw(screen *ebiten.Image) {
	// if flag != 0 {
	op1 := &text.DrawOptions{}
	op2 := &text.DrawOptions{}
	op3 := &text.DrawOptions{}
	op1.GeoM.Translate(float64(screenWidth)/2-600, 50)
	op2.GeoM.Translate(float64(screenWidth)/2+100, 50)
	op3.GeoM.Translate(float64(screenWidth)/2-200, 100)
	op1.ColorScale.ScaleWithColor(color.White)
	op2.ColorScale.ScaleWithColor(color.White)
	op3.ColorScale.ScaleWithColor(color.White)

	g.screenFon.img1(screen)                                              //Отрисовка поля
	g.screenFon.img2(screen)                                              //центральная отрисовка
	g.screenFon.img3(screen)                                              //Ворота
	g.paddle1.rectPaddle(screen)                                          //Ракетка первая
	g.paddle2.rectPaddle(screen)                                          //Ракетка вторая
	g.ball.rectBall(screen)                                               //Шарик
	g.screenFon.debug1(screen)                                            //Отрисовка FPS
	g.screenFon.debug2(screen, op1, g.paddle1.score)                      //Вывод счета первого игрока
	g.screenFon.debug2(screen, op2, g.paddle2.score)                      //Вывод счета второго игрока
	g.screenFon.debug3(screen, op3, int(g.ball.ballX), int(g.ball.ballY)) //Вывод счета второго игрока
}

// Экран
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{
		screenFon1: NewScreen2(color.RGBA{0, 0, 0, 255}),
		screenFon: NewScreen1(
			color.RGBA{99, 167, 16, 255},
			rand.IntN(20)+1,
			false,
		), //Инициализация экрана игра
		paddle1: NewPaddle(
			paddleX,
			screenHeight/2-paddleHeight/2,
			paddleWidth,
			paddleHeight,
			0,
			color.RGBA{31, 0, 204, 255},
		), //Инициализация Ракетки номер 1
		paddle2: NewPaddle(
			screenWidth-(paddleX+paddleWidth),
			screenHeight/2-paddleHeight/2,
			paddleWidth, paddleHeight, 0,
			color.RGBA{204, 0, 31, 255},
		), //Инициализация Ракетки номер 2
		ball: NewBall(
			screenWidth/2,
			screenHeight/2,
			ballRadius,
			color.RGBA{255, 208, 0, 255},
			float64(ballX),
			float64(ballY),
		), //Инициализация Мяча
		timer: timer,
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Paddle")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
