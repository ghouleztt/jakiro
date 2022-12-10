package base64

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
)

func Canvas() *container.TabItem {
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("please input sql...")
	tool := container.NewTabItem("Base64", container.NewVBox(
		widget.NewButton("compact", func() {
			input.SetText(input.Text + "\n########\n" + input.Text)
		}),
		widget.NewButton("copy", func() {
			_ = clipboard.WriteAll(input.Text)
		}),
		input,
	))
	return tool
}
