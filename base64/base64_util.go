package base64

import (
	"encoding/base64"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MakeUI() *container.TabItem {
	plain := widget.NewMultiLineEntry()
	plain.SetPlaceHolder("please input plain text...")
	plain.Resize(fyne.NewSize(800, 150))
	plain.Move(fyne.NewPos(10, 10))

	cipher := widget.NewMultiLineEntry()
	cipher.SetPlaceHolder("please input cipher text...")
	cipher.Resize(fyne.NewSize(800, 150))
	cipher.Move(fyne.NewPos(10, 220))

	base := widget.NewButton("base64", func() {
		cipher.SetText(base64.StdEncoding.EncodeToString([]byte(plain.Text)))
	})
	base.Resize(fyne.NewSize(80, 40))
	base.Move(fyne.NewPos(10, 170))

	debase := widget.NewButton("debase64", func() {
		tmp, _ := base64.StdEncoding.DecodeString(cipher.Text)
		plain.SetText(string(tmp))
	})
	debase.Resize(fyne.NewSize(80, 40))
	debase.Move(fyne.NewPos(100, 170))

	tool := container.NewTabItem("Base64", container.NewWithoutLayout(
		plain,
		base,
		debase,
		cipher,
	))
	return tool
}
