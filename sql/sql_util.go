package sql

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"regexp"
)

// MakeUI 当前还没有实现跳过单引号中的字符串不处理，
func MakeUI() *container.TabItem {
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("please input sql...")
	input.Resize(fyne.NewSize(800, 300))
	input.Move(fyne.NewPos(10, 60))

	btn := widget.NewButton("compact", func() {
		reg := regexp.MustCompile("\\s{2,}")
		input.SetText(reg.ReplaceAllString(input.Text, " "))
	})
	btn.Resize(fyne.NewSize(80, 40))
	btn.Move(fyne.NewPos(10, 10))

	tool := container.NewTabItem("SQL", container.NewWithoutLayout(
		btn,
		input,
	))
	return tool
}
