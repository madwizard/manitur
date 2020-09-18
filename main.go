package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	"github.com/faiface/beep/speaker"
	"image/color"
	"log"
	"strconv"
	"time"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"os"
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
				if int(tmp) == 60 {
					timeLeft.Text =  "1 minutes."
				} else if int(tmp) == 0 {
					playChime()
				} else {
					timeLeft.Text = strconv.Itoa(int(tmp/60)) + " minutes."
				}
			}
			timeLeft.Refresh()

		}
	})

	myWindow.SetContent(widget.NewHBox(widget.NewVBox(hello, progress),
		widget.NewVBox(timeLeft, start)))
	myWindow.ShowAndRun()
}

func playChime() {
	f, err := os.Open("sound/Temple Bell-SoundBible.com-756181215.mp3")
	if err != nil {
		log.Println("Couldn't open chime file! %s ", err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Print(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

}