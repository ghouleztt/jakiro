package datetime

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
	"image/color"
	"strconv"
	"time"
)

func updateTime(clock *widget.Label) {
	clock.SetText(time.Now().Format(datetimeFormat))
}

func updateTimeStamps(clock *widget.Label) {
	clock.SetText(strconv.FormatInt(time.Now().Unix(), 10))
}

func updateTimeStampms(clock *widget.Label) {
	clock.SetText(strconv.FormatInt(time.Now().UnixMilli(), 10))
}

const datetimeFormat string = "2006-01-02 15:04:05.999"

func MakeUI() *container.TabItem {
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	TimeLbl := canvas.NewText("Current Time: ", green)
	timeEntry := widget.NewLabel("")
	timeEntry.SetText(time.Now().Format(datetimeFormat))

	TimestampsLbl := canvas.NewText("Current Timestamp(s): ", green)
	timestampsEntry := widget.NewLabel("")
	timestampsEntry.SetText(strconv.FormatInt(time.Now().Unix(), 10))

	TimestampmsLbl := canvas.NewText("Current Timestamp(ms): ", green)
	timestampmsEntry := widget.NewLabel("")
	timestampmsEntry.SetText(strconv.FormatInt(time.Now().UnixMilli(), 10))

	go func() {
		for range time.Tick(time.Second) {
			updateTime(timeEntry)
			updateTimeStamps(timestampsEntry)
			updateTimeStampms(timestampmsEntry)
		}
	}()

	tsInput := widget.NewEntry()
	tsTypeA := widget.NewSelect([]string{"s", "ms"}, func(value string) {})
	tsTypeA.SetSelectedIndex(1)
	bjOutput := widget.NewEntry()

	bjInput := widget.NewEntry()
	tsTypeB := widget.NewSelect([]string{"s", "ms"}, func(value string) {})
	tsTypeB.SetSelectedIndex(1)
	tsOutput := widget.NewEntry()

	tool := container.NewTabItem("DateTime", container.NewVBox(
		// 当前时间
		container.NewGridWithColumns(2,
			container.NewGridWithColumns(2,
				TimeLbl, timeEntry),
			container.NewGridWithColumns(3,
				widget.NewButton("copy", func() {
					_ = clipboard.WriteAll(timeEntry.Text)
				}))),
		// 当前时间戳(s)
		container.NewGridWithColumns(2,
			container.NewGridWithColumns(2,
				TimestampsLbl, timestampsEntry),
			container.NewGridWithColumns(3,
				widget.NewButton("copy", func() {
					_ = clipboard.WriteAll(timestampsEntry.Text)
				}))),
		// 当前时间戳(ms)
		container.NewGridWithColumns(2,
			container.NewGridWithColumns(2,
				TimestampmsLbl, timestampmsEntry),
			container.NewGridWithColumns(3,
				widget.NewButton("copy", func() {
					_ = clipboard.WriteAll(timestampmsEntry.Text)
				}))),
		// 转换类型提示：时间戳转北京时间
		canvas.NewText("timestamp 2 beijingTime", green),
		// 时间戳输入
		container.NewGridWithColumns(3,
			tsInput,
			container.NewGridWithColumns(2,
				tsTypeA,
				widget.NewButton("convert", func() {
					if len(tsInput.Text) == 0 {
						return
					}
					if tsTypeA.Selected == "s" {
						ts, _ := strconv.ParseInt(tsInput.Text, 10, 64)
						tm := time.Unix(ts, 0)
						bjOutput.SetText(tm.Format(datetimeFormat))
					} else {
						ts, _ := strconv.ParseInt(tsInput.Text, 10, 64)
						tm := time.UnixMilli(ts)
						bjOutput.SetText(tm.Format(datetimeFormat))
					}

				})),
		),
		// 北京时间输出
		container.NewGridWithColumns(2,
			bjOutput,
			container.NewGridWithColumns(3,
				widget.NewButton("copy", func() {
					_ = clipboard.WriteAll(bjOutput.Text)
				}),
			),
		),
		// 转换类型提示：时间戳转北京时间
		canvas.NewText("beijingTime 2 timestamp", green),
		// 北京时间输入
		container.NewGridWithColumns(2,
			bjInput,
			container.NewGridWithColumns(3,
				widget.NewButton("convert", func() {
					if len(bjInput.Text) == 0 {
						return
					}
					loc, _ := time.LoadLocation("Asia/Shanghai")
					if tsTypeB.Selected == "s" {
						tm, _ := time.ParseInLocation(datetimeFormat, bjInput.Text, loc)
						tsOutput.SetText(strconv.FormatInt(tm.Unix(), 10))
					} else {
						tm, _ := time.ParseInLocation(datetimeFormat, bjInput.Text, loc)
						tsOutput.SetText(strconv.FormatInt(tm.UnixMilli(), 10))
					}
				}),
			),
		),
		// 时间戳输出
		container.NewGridWithColumns(3,
			tsOutput,
			container.NewGridWithColumns(2,
				tsTypeB,
				widget.NewButton("copy", func() {
					_ = clipboard.WriteAll(tsOutput.Text)
				})),
		),
	))
	return tool
}
