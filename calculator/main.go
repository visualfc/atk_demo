package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/visualfc/atk/tk"
)

func main() {
	tk.MainLoop(func() {
		tk.SetMainTheme(nil)
		w := NewWindow()
		w.SetTitle("Go 计算器")
		w.Center()
		w.ShowNormal()
		w.SetResizable(false, false)
	})
}

type Mode int

const (
	ModeWait Mode = iota
	ModeCmd
	ModeNext
)

type MainWindow struct {
	*tk.Window
	edit   *tk.Entry
	font   *tk.UserFont
	number string
}

func (w *MainWindow) inputNumber(s string) {
	if s == "." && strings.Contains(w.number, ".") {
		return
	}
	w.number += s
	w.edit.SetText(w.number)
}

func (w *MainWindow) inputCmd() {

}

type Button struct {
	*tk.Button
}

func (btn *Button) OnCommand(fn func()) {
	btn.Button.OnCommand(fn)
}

func (w *MainWindow) NewButton(text string) *Button {
	return w.NewButtonEx(text, 30, 30)
}

func (w *MainWindow) NewButtonEx(text string, width int, height int) *Button {
	blank := image.NewNRGBA(image.Rect(0, 0, width, height))
	img := tk.NewImage()
	img.SetImage(blank)
	btn := &Button{}
	btn.Button = tk.NewButton(w, text)
	btn.SetWidth(0)
	if !tk.HasTheme() {
		btn.SetWidth(width)
	}
	btn.SetImage(img)
	btn.SetCompound(tk.CompoundCenter)
	btn.SetAttributes(tk.WidgetAttrFont(w.font), tk.WidgetAttrBorderStyle(tk.BorderStyleGroove))
	return btn
}

func NewWindow() *MainWindow {
	mw := &MainWindow{}
	mw.Window = tk.MainWindow()
	mw.font = tk.LoadSysFont(tk.SysTextFont).Clone()
	mw.font.SetSize(20).SetBold(true)

	mw.edit = tk.NewEntry(mw)
	mw.edit.SetAlignment(tk.AlignmentRight)
	mw.edit.SetState(tk.StateReadOnly)
	mw.edit.SetFont(mw.font)
	mw.BindEvent("<Key>", func(e *tk.Event) {
		if e.KeyRune >= '0' && e.KeyRune <= '9' || e.KeyRune == '.' {
			mw.inputNumber(e.KeyText)
		}
	})

	var number [10]*Button
	for i := 0; i < 10; i++ {
		text := fmt.Sprintf("%v", i)
		number[i] = mw.NewButton(text)
		number[i].OnCommand(func() {
			mw.inputNumber(text)
		})
	}

	decimal := mw.NewButton(".")
	decimal.OnCommand(func() {
		mw.inputNumber(".")
	})

	clear := mw.NewButton("C")

	sig := mw.NewButton("+/-")

	precent := mw.NewButton("%")

	// + - * /
	plus := mw.NewButton("+")
	minus := mw.NewButton("-")
	multiple := mw.NewButton("×")
	division := mw.NewButton("÷")
	equal := mw.NewButton("=")

	grid := tk.NewGridLayout(mw)
	grid.AddWidgets(clear, nil, precent, division)
	grid.AddWidgets(number[7], number[8], number[9], multiple)
	grid.AddWidgets(number[4], number[5], number[6], minus)
	grid.AddWidgets(number[1], number[2], number[3], plus)
	grid.AddWidgets(number[0], sig, decimal, equal)

	hbox := tk.NewHPackLayout(mw)
	hbox.AddWidget(tk.NewLayoutSpacer(mw, 0, true), tk.PackAttrFillX(), tk.PackAttrExpand(true))
	hbox.AddLayout(grid, tk.PackAttrExpand(false), tk.PackAttrAnchor(tk.AnchorWest))

	vbox := tk.NewVPackLayout(mw)
	vbox.AddWidget(mw.edit, tk.PackAttrFillX(), tk.PackAttrExpand(false), tk.PackAttrAnchor(tk.AnchorNorth))
	vbox.AddLayout(hbox, tk.PackAttrFillX(), tk.PackAttrExpand(false), tk.PackAttrAnchor(tk.AnchorNorth))
	//vbox.AddWidget(tk.NewLayoutSpacer(mw, 0, true), tk.PackAttrFillY(), tk.PackAttrExpand(true))
	vbox.Repack()
	return mw
}
