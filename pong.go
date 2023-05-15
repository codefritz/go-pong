package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"os"
	"time"
)

var speedPong float32 = 1

func main() {

	a := app.New()
	window := a.NewWindow("Go-Pong!")

	ball := canvas.NewCircle(color.NRGBA{R: 0xee, G: 0xfe, A: 0xff})
	ball.Resize(fyne.NewSize(30, 30))

	playerA := canvas.NewRectangle(color.NRGBA{B: 0xff, A: 0xff})
	playerA.Resize(fyne.NewSize(80, 10))
	playerA.Move(fyne.NewPos(100, 240))

	window.SetContent(container.NewWithoutLayout(ball, playerA))

	window.Resize(fyne.NewSize(250, 250))
	window.SetPadded(false)

	var dirX float32 = 3
	var dirY float32 = 2

	go func() {
		for range time.Tick(time.Millisecond * 40) {
			if ball.Position().X > 220 || ball.Position().X < 0 {
				dirX = -1 * dirX
			}
			if ball.Position().Y < 0 {
				dirY = -1 * dirY
			}
			if ball.Position().Y > 210 {
				if ball.Position().X > playerA.Position().X && ball.Position().X < playerA.Position().X+80 {
					dirY = -1 * dirY
					speedPong *= 1.1
				}
			}
			ball.Move(fyne.NewPos(ball.Position().X+(dirX*speedPong), ball.Position().Y+(dirY*speedPong)))
			//fmt.Println(ball.Position().X, ball.Position().Y, playerA.Position().X, playerA.Position().Y)
			if ball.Position().Y > 250 {
				fmt.Println("You loose, speed: ", speedPong)
				os.Exit(8)
			}
		}
	}()

	window.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case fyne.KeyLeft:
			playerA.Move(fyne.NewPos(playerA.Position().X-4, playerA.Position().Y))
		case fyne.KeyRight:
			playerA.Move(fyne.NewPos(playerA.Position().X+4, playerA.Position().Y))

		}
	})

	window.ShowAndRun()

}
