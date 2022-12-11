package json

import (
	"bytes"
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
)

func format(textArea *widget.Entry) {
	if len(textArea.Text) == 0 {
		return
	}

	var str bytes.Buffer
	_ = json.Indent(&str, []byte(textArea.Text), "", "    ")
	textArea.SetText(str.String())
}

func compact(textArea *widget.Entry) {
	if len(textArea.Text) == 0 {
		return
	}

	var str bytes.Buffer
	_ = json.Compact(&str, []byte(textArea.Text))
	textArea.SetText(str.String())
}

func MakeUI() *container.TabItem {
	textArea := widget.NewMultiLineEntry()
	textArea.SetPlaceHolder("please input json...")
	textArea.Resize(fyne.NewSize(800, 300))
	textArea.Move(fyne.NewPos(10, 60))

	formatBtn := widget.NewButton("format", func() {
		format(textArea)
	})
	formatBtn.Resize(fyne.NewSize(80, 40))
	formatBtn.Move(fyne.NewPos(10, 10))

	compactBtn := widget.NewButton("compact", func() {
		compact(textArea)
	})
	compactBtn.Resize(fyne.NewSize(80, 40))
	compactBtn.Move(fyne.NewPos(100, 10))

	copyBtn := widget.NewButton("copy", func() {
		_ = clipboard.WriteAll(textArea.Text)
	})
	copyBtn.Resize(fyne.NewSize(80, 40))
	copyBtn.Move(fyne.NewPos(190, 10))

	tool := container.NewTabItem("JSON", container.NewWithoutLayout(
		formatBtn, compactBtn, copyBtn,
		textArea,
	))
	return tool
}
