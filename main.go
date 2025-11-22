package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/exp/rand"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	screenHeight = 800
	screenWeight = 1600
	padh         = 160
	padw         = 20
)

type Paddle struct {
	x, y          int
	color         color.Color
	width, height int
	score         int
}

type Ball struct {
	x, y   int
	color  color.Color
	radius int
}

type Game struct {
	face            font.Face
	backgroundColor color.Color
	gameOver        bool
	paddle1         *Paddle
	paddle2         *Paddle
	ball            *Ball
	direction       int
	ballX           float64
	ballY           float64
}

// Движение ракетки
func (g *Game) movePaddle() {
	//Первая ракетка
	if ebiten.IsKeyPressed(ebiten.KeyW) && g.paddle1.y > 10 {
		g.paddle1.y -= 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) && g.paddle1.y < screenHeight-g.paddle1.height-10 {
		g.paddle1.y += 10
	}

	//Вторая ракетка
	if ebiten.IsKeyPressed(ebiten.KeyUp) && g.paddle2.y > 10 {
		g.paddle2.y -= 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) && g.paddle2.y < screenHeight-g.paddle2.height-10 {
		g.paddle2.y += 10
	}
}

// Движение мяча
func (g *Game) moveBall() {
	if g.direction > 0 && g.direction <= 5 {
		g.ball.x += int(g.ballX)
		g.ball.y += int(g.ballY)
	} else if g.direction > 5 && g.direction <= 10 {
		g.ball.x += int(g.ballX)
		g.ball.y -= int(g.ballY)
	} else if g.direction > 10 && g.direction <= 15 {
		g.ball.x += int(g.ballX)
		g.ball.y += int(g.ballY)
	} else if g.direction > 15 && g.direction <= 20 {
		g.ball.x += int(g.ballX)
		g.ball.y -= int(g.ballY)
	}
}

// Коллизит мяча с полем
func (g *Game) colissionWithField() {
	if g.ball.y < g.ball.radius || g.ball.y > screenHeight-g.ball.radius {
		g.ballY = -g.ballY
	}
}

func (g *Game) colissionWithGates() {
	if g.ball.x < 0+g.ball.radius {
		g.paddle2.score++
		g.gameOver = true
	} else if g.ball.x > screenWeight-g.ball.radius {
		g.paddle1.score++
		g.gameOver = true
	}
}

// Рестарт игры
func (g *Game) resetGame() {
	//Если изменять хотябы один элемент объекта то нужно переинициализировать весь объект, иначе все обнулится
	g.paddle1 = &Paddle{
		x:      20,
		y:      screenHeight/2 - padh/2,
		color:  color.RGBA{31, 0, 204, 255},
		width:  padw,
		height: padh,
		score:  g.paddle1.score,
	}
	g.paddle2 = &Paddle{
		x:      screenWeight - (20 + padw),
		y:      screenHeight/2 - padh/2,
		color:  color.RGBA{204, 0, 31, 255},
		width:  padw,
		height: padh,
		score:  g.paddle2.score,
	}
	g.ball = &Ball{
		x:      screenWeight / 2,
		y:      screenHeight / 2,
		color:  color.RGBA{255, 208, 0, 255},
		radius: 20,
	}

	g.ballX = 10
	g.ballY = 10
	g.gameOver = false
	ran := rand.Intn(20) + 1
	if ran > 5 && ran <= 15 {
		g.ballX = -10
	} else {
		g.ballX = 10
	}
	g.direction = ran
}

// Коллизии мяча с ракетками
func (g *Game) collisionWithRacket() {
	LeftBall := g.ball.x - g.ball.radius
	RightBall := g.ball.x + g.ball.radius
	UpBall := g.ball.y - g.ball.radius
	DownBall := g.ball.y + g.ball.radius

	UpPad1 := g.paddle1.y
	DownPad1 := g.paddle1.y + g.paddle1.height
	RightPad1 := g.paddle1.x + g.paddle1.width

	UpPad2 := g.paddle2.y
	DownPad2 := g.paddle2.y + g.paddle2.height
	LeftPad2 := g.paddle2.x

	if LeftBall <= RightPad1 && (g.ball.y >= UpPad1 && g.ball.y <= DownPad1) {
		if g.ballX >= 15 {
			g.ballX = -(g.ballX + float64(rand.Intn(3)-2))
		} else if g.ballX <= -15 {
			g.ballX = -g.ballX + float64(rand.Intn(3)-2)
		} else if g.ballX > -15 && g.ballX < -5 {
			g.ballX = -g.ballX + float64(rand.Intn(5)-2)
		} else if g.ballX > 5 && g.ballX < 15 {
			g.ballX = -g.ballX + float64(rand.Intn(5)-2)
		} else if g.ballX >= 0 && g.ballX <= 5 {
			g.ballX = -(g.ballX + float64(rand.Intn(3)))
		} else if g.ballX < 0 && g.ballX >= -5 {
			g.ballX = -g.ballX + float64(rand.Intn(3))
		}
	}
	if RightBall >= LeftPad2 && (g.ball.y >= UpPad2 && g.ball.y <= DownPad2) {
		if g.ballX >= 15 {
			g.ballX = -(g.ballX + float64(rand.Intn(3)-2))
		} else if g.ballX <= -15 {
			g.ballX = -g.ballX + float64(rand.Intn(3)-2)
		} else if g.ballX > -15 && g.ballX < -5 {
			g.ballX = -g.ballX + float64(rand.Intn(5)-2)
		} else if g.ballX > 5 && g.ballX < 15 {
			g.ballX = -g.ballX + float64(rand.Intn(5)-2)
		} else if g.ballX >= 0 && g.ballX <= 5 {
			g.ballX = -(g.ballX + float64(rand.Intn(3)))
		} else if g.ballX < 0 && g.ballX >= -5 {
			g.ballX = -g.ballX + float64(rand.Intn(3))
		}
	}
	if UpBall <= DownPad1 && (g.ball.x >= 20 && g.ball.x <= 40) || UpBall <= DownPad2 && (g.ball.x >= 1560 && g.ball.x <= 1580) {
		g.ballY = -g.ballY
	}
	if DownBall <= UpPad1 && (g.ball.x >= 20 && g.ball.x <= 40) || DownBall <= UpPad2 && (g.ball.x >= 1560 && g.ball.x <= 1580) {
		g.ballY = -g.ballY
	}
}

// Движения
func (g *Game) Update() error {
	//полноэкранный режим по кнопке F
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	//Рестарт Игры
	if g.gameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.resetGame()
			return nil
		}
		return nil
	}
	//Рестарт Игры
	g.moveBall()
	g.movePaddle()
	g.colissionWithField()
	g.collisionWithRacket()
	g.colissionWithGates()
	return nil
}

// Отрисовка
func (g *Game) Draw(screen *ebiten.Image) {
	//Отрисовка поля
	screen.Fill(g.backgroundColor)

	//Кольцо по центру
	vector.FillCircle(
		screen,
		float32(screenWeight/2),
		float32(screenHeight/2),
		float32(250),
		color.RGBA{255, 255, 255, 255},
		false,
	)
	vector.FillCircle(
		screen,
		float32(screenWeight/2),
		float32(screenHeight/2),
		float32(240),
		g.backgroundColor,
		false,
	)

	//круг по центру
	vector.FillCircle(
		screen,
		float32(screenWeight/2),
		float32(screenHeight/2),
		float32(30),
		color.RGBA{255, 255, 255, 255},
		false,
	)

	//поляна 1
	vector.FillRect(
		screen,
		float32(80),
		float32(0),
		float32(20),
		float32(screenHeight),
		color.RGBA{255, 255, 255, 255},
		false,
	)

	//поляна 2
	vector.FillRect(
		screen,
		float32(1500),
		float32(0),
		float32(20),
		float32(screenHeight),
		color.RGBA{255, 255, 255, 255},
		false,
	)

	//Ракетка первая
	vector.FillRect(
		screen,
		float32(g.paddle1.x),
		float32(g.paddle1.y),
		float32(g.paddle1.width),
		float32(g.paddle1.height),
		g.paddle1.color,
		false,
	)

	//линия по центру
	for i := 0; i <= 16; i++ {
		if i%2 == 0 {
			vector.FillRect(
				screen,
				float32(screenWeight/2-5),
				float32(0+55*i),
				float32(10),
				float32(50),
				color.RGBA{255, 255, 255, 255},
				false,
			)
		}
	}

	//Ракетка вторая
	vector.FillRect(
		screen,
		float32(g.paddle2.x),
		float32(g.paddle2.y),
		float32(g.paddle2.width),
		float32(g.paddle2.height),
		g.paddle2.color,
		false,
	)

	//Шарик
	vector.FillCircle(
		screen,
		float32(g.ball.x),
		float32(g.ball.y),
		float32(g.ball.radius),
		g.ball.color,
		false,
	)

	ebitenutil.DebugPrint( //Отрисовка FPS и Координат
		screen,
		fmt.Sprintf("FPS: %0.2f\n", ebiten.ActualFPS()),
	)
	//Вывод счета первого игрока
	text.Draw(screen, fmt.Sprintf("first player score: %d", g.paddle1.score), g.face, 150, 50, color.White)

	//Вывод счета второго игрока
	text.Draw(screen, fmt.Sprintf("second player's score: %d", g.paddle2.score), g.face, screenWeight-600, 50, color.White)

	//Вывод счета второго игрока
	text.Draw(screen, fmt.Sprintf("coord ball x:y %f:%f", g.ballX, g.ballY), g.face, screenWeight/2-300, 100, color.White)
}

// Экран
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWeight, screenHeight
}

func main() {

	//ШРифт
	fontData, err := os.ReadFile("text/Y224.otf")
	if err != nil {
		log.Fatal(err)
	}

	tt, err := opentype.Parse(fontData)
	if err != nil {
		log.Fatal(err)
	}

	face, err := opentype.NewFace(
		tt,
		&opentype.FaceOptions{
			Size:    16,
			DPI:     72,
			Hinting: font.HintingFull,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	//Шрифт

	//Сид рандома
	seed := time.Now().UnixNano()
	rand.Seed(uint64(seed))
	ran := rand.Intn(20) + 1
	var dir float64 = 0
	if ran > 5 && ran <= 15 {
		dir = -10
	} else {
		dir = 10
	}
	//Сид рандома

	game := &Game{
		face:            face,
		backgroundColor: color.RGBA{99, 167, 16, 255},
		paddle1: &Paddle{ //20:320
			x:      20,
			y:      screenHeight/2 - padh/2, //  800/2-160/2=400-80=320
			color:  color.RGBA{31, 0, 204, 255},
			width:  padw,
			height: padh,
			score:  0,
		},
		paddle2: &Paddle{ //1560:320
			x:      screenWeight - (20 + padw), // 1600-(20+20)=1600-40=1560
			y:      screenHeight/2 - padh/2,    //   800/2-160/2=400-80=320
			color:  color.RGBA{204, 0, 31, 255},
			width:  padw,
			height: padh,
			score:  0,
		},
		ball: &Ball{
			x:      screenWeight / 2,
			y:      screenHeight / 2,
			color:  color.RGBA{255, 208, 0, 255},
			radius: 20,
		},
		direction: ran,
		ballX:     dir,
		ballY:     10,
		gameOver:  false,
	}
	ebiten.SetWindowSize(screenWeight, screenHeight)
	ebiten.SetWindowTitle("Paddle")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
