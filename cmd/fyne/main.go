package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

var green = color.NRGBA{R: 0, G: 180, B: 0, A: 255}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	myWindow.Resize(fyne.NewSize(450, 600))
	text := canvas.NewText("TEXT 1", green)
	text2 := canvas.NewText("tex to", green)
	text2.Move(fyne.NewPos(50, 50))
	content := container.NewWithoutLayout(text, text2)
	//myWindow.Resize(fyne.NewSize(100, 100))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
