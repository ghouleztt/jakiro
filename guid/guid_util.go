package guid

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
	"github.com/google/uuid"
	"image/color"
	"strconv"
	"strings"
)

func MakeUI() *container.TabItem {
	// quantity input
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	quantityLbl := canvas.NewText("quantity: ", green)
	quantityLbl.Resize(fyne.NewSize(80, 40))
	quantityLbl.Move(fyne.NewPos(10, 10))

	quantity := widget.NewEntry()
	quantity.Resize(fyne.NewSize(80, 40))
	quantity.Move(fyne.NewPos(100, 10))

	// generate option: wrap with quot
	quotOpt := widget.NewCheck("add quot", func(b bool) {
	})
	quotOpt.Resize(fyne.NewSize(80, 40))
	quotOpt.Move(fyne.NewPos(210, 10))

	// generate option: include dash
	dashOpt := widget.NewCheck("add dash", func(b bool) {
	})
	dashOpt.Resize(fyne.NewSize(80, 40))
	dashOpt.Move(fyne.NewPos(320, 10))

	// generate option: add comma at the end of each line
	commaOpt := widget.NewCheck("add comma", func(b bool) {
	})
	commaOpt.Resize(fyne.NewSize(80, 40))
	commaOpt.Move(fyne.NewPos(430, 10))

	// output text area
	output := widget.NewMultiLineEntry()
	output.Resize(fyne.NewSize(800, 300))
	output.Move(fyne.NewPos(10, 110))

	// generate button
	generateBtn := widget.NewButton("generate", func() {
		output.SetText("")
		q, _ := strconv.Atoi(quantity.Text)

		result := ""
		for i := 0; i < q; i++ {
			curGuid := uuid.New().String()
			if quotOpt.Checked {
				curGuid = "\"" + curGuid + "\""
			}
			if !dashOpt.Checked {
				curGuid = strings.Replace(curGuid, "-", "", -1)
			}
			if commaOpt.Checked {
				curGuid = curGuid + ","
			}
			result = result + curGuid + "\n"
		}

		output.SetText(result)
	})
	generateBtn.Resize(fyne.NewSize(80, 40))
	generateBtn.Move(fyne.NewPos(10, 60))

	// copy button
	copyBtn := widget.NewButton("copy", func() {
		_ = clipboard.WriteAll(output.Text)
	})
	copyBtn.Resize(fyne.NewSize(80, 40))
	copyBtn.Move(fyne.NewPos(100, 60))

	tool := container.NewTabItem("GUID", container.NewWithoutLayout(
		quantityLbl,
		quantity,
		quotOpt,
		dashOpt,
		commaOpt,
		generateBtn,
		copyBtn,
		output,
	))
	return tool
}
