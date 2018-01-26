package main

import (
	"strconv"

	"github.com/visualfc/atk/tk"
)

var (
	sizeList = []int{9, 10, 11, 12, 13, 14, 16, 18, 24, 36, 48, 72, 96, 144}
)

func main() {
	err := tk.Init()
	if err != nil {
		return
	}

	tk.MainLoop(func() {
		mw := tk.MainWindow()

		font := tk.NewUserFont("", 36)
		info := tk.NewLabel(mw, "Hello, 中国")
		info.SetFont(font)

		var onResize func()

		cmbFamily := tk.NewComboBox(mw)
		cmbFamily.SetValues(tk.FontFamilieList())

		upfamily := func() {
			font.SetFamily(cmbFamily.CurrentText())
			if onResize != nil {
				onResize()
			}
		}
		cmbFamily.OnSelected(upfamily)
		cmbFamily.OnEditReturn(upfamily)
		cmbFamily.SetCurrentText(font.Family())

		cmbSize := tk.NewComboBox(mw)
		var values []string
		for _, v := range sizeList {
			values = append(values, strconv.Itoa(v))
		}
		cmbSize.SetValues(values)
		cmbSize.SetCurrentText("36")
		upsize := func() {
			v, err := strconv.Atoi(cmbSize.CurrentText())
			if err == nil {
				font.SetSize(v)
				if onResize != nil {
					onResize()
				}
			}
		}
		cmbSize.OnEditReturn(upsize)
		cmbSize.OnSelected(upsize)

		chkBold := tk.NewCheckButton(mw, "Bold")
		chkItalic := tk.NewCheckButton(mw, "Italic")
		chkUnderline := tk.NewCheckButton(mw, "Underline")
		chkOverstrike := tk.NewCheckButton(mw, "Overstrike")

		chkBold.OnCommand(func() {
			font.SetBold(chkBold.IsChecked())
		})
		chkItalic.OnCommand(func() {
			font.SetItalic(chkItalic.IsChecked())
		})
		chkUnderline.OnCommand(func() {
			font.SetUnderline(chkUnderline.IsChecked())
		})
		chkOverstrike.OnCommand(func() {
			font.SetOverstrike(chkOverstrike.IsChecked())
		})

		hpk := tk.NewHPackLayout(mw)
		hpk.AddWidget(tk.NewLabel(mw, "Font"))
		hpk.AddWidget(cmbFamily)
		hpk.AddWidget(tk.NewLabel(mw, "Size"))
		hpk.AddWidget(cmbSize)

		hpk2 := tk.NewHPackLayout(mw)
		hpk2.AddWidget(chkBold)
		hpk2.AddWidget(chkItalic)
		hpk2.AddWidget(chkUnderline)
		hpk2.AddWidget(chkOverstrike)

		vpk := tk.NewVPackLayout(mw)
		vpk.SetBorderWidth(10)
		vpk.AddWidgetEx(info, tk.FillY, true, tk.AnchorCenter)
		vpk.AddWidget(tk.NewLayoutSpacer(mw, 0, true))
		vpk.AddLayoutEx(hpk, tk.FillNone, false, tk.AnchorWest)
		vpk.AddLayoutEx(hpk2, tk.FillNone, false, tk.AnchorWest)

		onResize = func() {
			width := font.MeasureTextWidth(info.Text())
			if width > mw.Width() {
				mw.SetWidth(width + 20)
			} else if width < 400 {
				mw.SetWidth(400)
			}
		}

		mw.ResizeN(400, 300)
		mw.Center()
		mw.ShowNormal()
	})
}
