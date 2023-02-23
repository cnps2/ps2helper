package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-vgo/robotgo"
	"golang.design/x/hotkey"
)

func main() {
	w := app.New().NewWindow("ps2helper")
	label := widget.NewLabel("Hello golang.design!")
	button := widget.NewButton("Hi!", func() {
		label.SetText("Welcome :)")
		println("Hi!")
	})
	w.SetContent(container.NewVBox(label, button))

	go func() {
		// Register a desired hotkey.
		hk := hotkey.New([]hotkey.Modifier{}, hotkey.KeyF10)
		if err := hk.Register(); err != nil {
			panic("hotkey registration failed")
		}
		// Start listen hotkey event whenever it is ready.
		for range hk.Keydown() {
			robotgo.Click()
			robotgo.MilliSleep(10)
			robotgo.Click()
			robotgo.MilliSleep(10)
			robotgo.Toggle()
			robotgo.MilliSleep(600)
			robotgo.Toggle("left", "up")
			robotgo.Move(80, 80)
			sx, sy := robotgo.GetScreenSize()
			println("get screen size: ", sx, sy)
			// button.Tapped(&fyne.PointEvent{})

		}
	}()

	w.ShowAndRun()
}
