package json

import (
	"bytes"
	"encoding/json"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
)

func format(input *widget.Entry) {
	if len(input.Text) == 0 {
		return
	}

	var str bytes.Buffer
	_ = json.Indent(&str, []byte(input.Text), "", "    ")
	input.SetText(str.String())
}

func compact(input *widget.Entry) {
	if len(input.Text) == 0 {
		return
	}

	var str bytes.Buffer
	_ = json.Compact(&str, []byte(input.Text))
	input.SetText(str.String())
}

func Canvas() *container.TabItem {
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("please input json...")

	textArea := container.NewMax(input)

	buttons := container.NewHBox(
		widget.NewButton("format", func() {
			format(input)
		}),
		widget.NewButton("compact", func() {
			compact(input)
		}),
		widget.NewButton("copy", func() {
			_ = clipboard.WriteAll(input.Text)
		}),
	)

	tool := container.NewTabItem("JSON", container.NewVBox(
		buttons,
		textArea,
	))
	return tool
}
