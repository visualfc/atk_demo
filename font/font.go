package main

import (
	"strconv"

	"github.com/visualfc/atk/tk"
)

var (
	sizeList = []int{9, 10, 11, 12, 13, 14, 16, 18, 24, 36, 48, 72, 96, 144}
)

type Window struct {
	*tk.Window
	font *tk.UserFont
	info *tk.Label
	vpk  *tk.PackLayout
}

func NewWindow() *Window {
	mw := tk.RootWindow()

	font := tk.NewUserFont("", 36)
	info := tk.NewLabel(mw, "Hello, 中国")
	info.SetFont(font)

	vpk := tk.NewVPackLayout(mw)
	w := &Window{mw, font, info, vpk}
	w.Init()
	return w
}

func (w *Window) Init() {
	cmbFamily := tk.NewComboBox(w)
	cmbFamily.SetValues(tk.FontFamilieList())

	upfamily := func() {
		w.font.SetFamily(cmbFamily.CurrentText())
		w.Resize()
	}
	cmbFamily.OnSelected(upfamily)
	cmbFamily.OnEditReturn(upfamily)
	cmbFamily.SetCurrentText(w.font.Family())

	cmbSize := tk.NewComboBox(w)
	var values []string
	for _, v := range sizeList {
		values = append(values, strconv.Itoa(v))
	}
	cmbSize.SetValues(values)
	cmbSize.SetCurrentText("36")
	upsize := func() {
		v, err := strconv.Atoi(cmbSize.CurrentText())
		if err == nil {
			w.font.SetSize(v)
			w.Resize()
		}
	}
	cmbSize.OnEditReturn(upsize)
	cmbSize.OnSelected(upsize)

	chkBold := tk.NewCheckButton(w, "Bold")
	chkItalic := tk.NewCheckButton(w, "Italic")
	chkUnderline := tk.NewCheckButton(w, "Underline")
	chkOverstrike := tk.NewCheckButton(w, "Overstrike")

	chkBold.OnCommand(func() {
		w.font.SetBold(chkBold.IsChecked())
	})
	chkItalic.OnCommand(func() {
		w.font.SetItalic(chkItalic.IsChecked())
	})
	chkUnderline.OnCommand(func() {
		w.font.SetUnderline(chkUnderline.IsChecked())
	})
	chkOverstrike.OnCommand(func() {
		w.font.SetOverstrike(chkOverstrike.IsChecked())
	})

	hpk := tk.NewHPackLayout(w)
	hpk.AddWidget(tk.NewLabel(w, "Family"))
	hpk.AddWidget(cmbFamily)
	hpk.AddWidget(tk.NewLabel(w, "Size"))
	hpk.AddWidget(cmbSize)

	hpk2 := tk.NewHPackLayout(w)
	hpk2.AddWidget(chkBold)
	hpk2.AddWidget(chkItalic)
	hpk2.AddWidget(chkUnderline)
	hpk2.AddWidget(chkOverstrike)

	w.vpk.SetBorderWidth(10)
	w.vpk.AddWidgetEx(w.info, tk.FillY, true, tk.AnchorCenter)
	w.vpk.AddWidget(tk.NewLayoutSpacer(w, 0, true))
	w.vpk.AddWidgetEx(hpk, tk.FillNone, false, tk.AnchorWest)
	w.vpk.AddWidgetEx(hpk2, tk.FillNone, false, tk.AnchorWest)
}

func (w *Window) Resize() {
	width := w.font.MeasureTextWidth(w.info.Text()) + w.vpk.BorderWidth()*2
	if width > w.Width() {
		w.SetWidth(width)
	} else if width < 400 {
		w.SetWidth(400)
	}
}

func main() {
	tk.MainLoop(func() {
		mw := NewWindow()
		mw.SetTitle("ATK Font Demo")
		mw.ResizeN(400, 300)
		mw.Center()
		mw.ShowNormal()
	})
}
