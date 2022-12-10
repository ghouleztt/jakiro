package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("03:04:05")
	clock.SetText(formatted)
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(600, 400))

	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text1 := canvas.NewText("Current Time: ", green)
	text1.Move(fyne.NewPos(20, 20))
	clock := widget.NewLabel("")
	formatted := time.Now().Format("03:04:05")
	clock.SetText(formatted)
	clock.Move(fyne.NewPos(120, 12))

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(

		container.NewAppTabs(container.NewTabItem("aaa", container.NewWithoutLayout(
			text1)), container.NewTabItem("bbb", container.NewWithoutLayout(
			clock))),
		container.NewWithoutLayout(
			text1,
			clock),
		hello,
		widget.NewButton("Hi!", func() {
			hello.TextStyle.Bold = true
			hello.SetText("Welcome, Logan:)")
		}),

		widget.NewButton("Change!", func() {
			hello.SetText("Changed!!!")
		}),

		widget.NewButton("creatWindow", func() {
			w2 := a.NewWindow("subWindow")
			w2.Resize(fyne.NewSize(200, 200))

			greeting := widget.NewLabel("Hello")
			in := widget.NewEntry()
			in.OnChanged = func(content string) {
				if len(content) > 0 {
					greeting.SetText("Hello," + content + "!")
				} else {
					greeting.SetText("Hello")
				}
			}
			w2.SetContent(container.NewVBox(
				greeting,
				in,
			))
			w2.Show()
		}),
	))

	w.ShowAndRun()
}
