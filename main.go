package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/flopp/go-findfont"
	"github.com/ghouleztt/jakiro/base64"
	"github.com/ghouleztt/jakiro/datetime"
	"github.com/ghouleztt/jakiro/guid"
	"github.com/ghouleztt/jakiro/json"
	"os"
	"strings"
)

const (
	// 黑体
	simhei = "simhei.ttf"
	// 楷体
	simkai = "simkai.ttf"
	// 宋体
	simsun = "simsun.ttc"
	// 华文细黑
	stxihei  = "STXIHEI.TTF"
	fyneFont = "FYNE_FONT"
)

func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		if strings.Contains(path, stxihei) {
			_ = os.Setenv(fyneFont, path)
			break
		}
	}
}
func main() {
	a := app.New()

	//设置图标
	icon, _ := fyne.LoadResourceFromPath("jakiro.bmp")
	a.SetIcon(icon)

	w := a.NewWindow("Jakiro")
	w.Resize(fyne.NewSize(800, 500))
	w.CenterOnScreen()

	w.SetContent(container.NewAppTabs(
		json.MakeUI(),
		//sql.MakeUI(),
		datetime.MakeUI(),
		base64.MakeUI(),
		guid.MakeUI(),
	))
	w.ShowAndRun()
	_ = os.Unsetenv(fyneFont)
}
