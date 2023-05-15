package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"strconv"
	"time"
)

var speedPong float32 = 1
var points = 0

var green = color.NRGBA{G: 0xae, A: 0xff}

const windowSize = 250

func main() {

	pongApp := app.New()
	window := pongApp.NewWindow("Go-Pong!")

	gradient := gradient()
	ball := createBall(color.NRGBA{R: 0xee, G: 0xfe, A: 0xff})
	shaddows := make([]fyne.CanvasObject, 0, 4)
	for n := 1; n <= 5; n++ {
		shaddows = append(shaddows, createBall(color.NRGBA{R: 0xfe, G: 0xff, B: 0x22 + -uint8(2*n), A: 0xff - uint8(80*n)}))
	}
	player := player()
	counter := counter()

	objs := make([]fyne.CanvasObject, 0, 4)

	objs = append(objs, gradient)
	for n := 1; n <= 10; n++ {
		objs = append(objs, horizontalLine(windowSize/10*float32(n)))
	}
	objs = append(objs, player)
	var positions []fyne.Position
	for n := len(shaddows) - 1; n >= 0; n-- {
		objs = append(objs, shaddows[n])
		positions = append(positions, ball.Position())
	}
	objs = append(objs, ball)
	objs = append(objs, counter)

	window.SetContent(container.NewWithoutLayout(objs...))
	window.Resize(fyne.NewSize(windowSize, windowSize))
	window.SetPadded(false)

	var dirX float32 = 3
	var dirY float32 = 2
	var br = false
	go func() {
		for range time.Tick(time.Millisecond * 40) {
			if ball.Position().X > 220 || ball.Position().X < 0 {
				dirX = -1 * dirX
			}

			if ball.Position().Y < 0 {
				dirY = -1 * dirY
			}
			if ball.Position().Y > 200 && !br {
				if ball.Position().X+15 > player.Position().X && ball.Position().X+15 < player.Position().X+80 {
					dirY = -1 * dirY
					speedPong *= 1.1
					points += 1
					counter.Text = strconv.Itoa(points)
					counter.Refresh()
					br = true
				}
			} else {
				// ball is out of 200 px zone
				br = ball.Position().Y < 200
			}

			positions = append(positions, ball.Position())[1:]

			for ix := range shaddows {
				shaddows[ix].Move(positions[ix])
			}

			newPos := fyne.NewPos(ball.Position().X+(dirX*speedPong), ball.Position().Y+(dirY*speedPong))
			ball.Move(newPos)

			if ball.Position().Y > windowSize {
				window.SetContent(container.NewWithoutLayout(gradient, gameOver()))
			}
		}
	}()

	window.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case fyne.KeyLeft:
			if player.Position().X-4 > 0 {
				player.Move(fyne.NewPos(player.Position().X-4, player.Position().Y))
			}
		case fyne.KeyRight:
			if player.Position().X+80 < windowSize {
				player.Move(fyne.NewPos(player.Position().X+4, player.Position().Y))
			}

		}
	})

	window.ShowAndRun()
}

func gradient() *canvas.LinearGradient {
	gradient := canvas.NewVerticalGradient(color.Black, color.White)
	gradient.Resize(fyne.NewSize(windowSize, windowSize))
	return gradient
}

func createBall(clr color.Color) *canvas.Circle {
	ball := canvas.NewCircle(clr)
	ball.Resize(fyne.NewSize(30, 30))
	return ball
}

func player() *canvas.Rectangle {
	playerA := canvas.NewRectangle(color.NRGBA{B: 0xff, A: 0xff})
	playerA.Resize(fyne.NewSize(80, 10))
	playerA.Move(fyne.NewPos(100, 230))
	return playerA
}

func counter() *canvas.Text {
	counter := canvas.NewText("0", color.White)
	counter.Alignment = fyne.TextAlignTrailing
	counter.TextStyle = fyne.TextStyle{Italic: true, Bold: true}
	counter.Move(fyne.NewPos(12, 10))
	return counter
}

func gameOver() *canvas.Text {
	counter := canvas.NewText("* GAME OVER *", color.White)
	counter.TextStyle = fyne.TextStyle{Bold: true}
	counter.TextSize = 30
	counter.Move(fyne.NewPos(10, 100))
	return counter
}

func horizontalLine(xPos float32) *canvas.Line {
	line := canvas.NewLine(green)
	line.Resize(fyne.NewSize(2, 250))
	line.Move(fyne.NewPos(xPos, 0))
	return line
}
