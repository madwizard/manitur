package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	"image/color"
	"strconv"
	"time"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Timer")

	text1 := canvas.NewText("Hello. Time to next exercise: ", color.White)
	text2 := canvas.NewText("30 minutes.", color.White)
 	progress := widget.NewProgressBar()
 	progress.Max = 30.0

	button := widget.NewButton("Start", func() {
		for i := 0.0; i <= 3.0; i += 0.1 {
			time.Sleep(time.Second)
			progress.SetValue(i)
			tmp := 30.0 - i
			text2.Text = strconv.Itoa(int(tmp))
			text2.Refresh()

		}
	})

	myWindow.SetContent(widget.NewHBox(widget.NewVBox(text1, progress),
		widget.NewVBox(text2, button)))
	myWindow.ShowAndRun()
}
