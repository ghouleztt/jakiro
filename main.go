package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/ghouleztt/jakiro/base64"
	"github.com/ghouleztt/jakiro/datetime"
	"github.com/ghouleztt/jakiro/json"
	"github.com/ghouleztt/jakiro/sql"
)

func main() {
	a := app.New()
	w := a.NewWindow("Jakiro")
	w.Resize(fyne.NewSize(900, 700))

	w.SetContent(container.NewVBox(
		container.NewAppTabs(
			json.Canvas(),
			sql.Canvas(),
			datetime.Canvas(),
			base64.Canvas(),
		),
	))
	w.ShowAndRun()
}
