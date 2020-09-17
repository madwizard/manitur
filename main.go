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

	hello := canvas.NewText("Hello. Time to next exercise: ", color.White)
	timeLeft := canvas.NewText("30 minutes.", color.White)
 	progress := widget.NewProgressBar()
 	progress.Max = 180.0

	start := widget.NewButton("Start", func() {
		for i := 0.0; i <= 180.0; i += 0.1 {
			time.Sleep(time.Second)
			progress.SetValue(i)
			tmp := (180.0 - i)*10
			if(int(tmp) % 60) == 0 {
				timeLeft.Text = strconv.Itoa(int(tmp / 60))
			}
			timeLeft.Refresh()

		}
	})

	myWindow.SetContent(widget.NewHBox(widget.NewVBox(hello, progress),
		widget.NewVBox(timeLeft, start)))
	myWindow.ShowAndRun()
}
