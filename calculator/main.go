package main

import (
	"fmt"

	"github.com/visualfc/atk/tk"
)

func main() {
	tk.MainLoop(func() {
		w := NewWindow()
		w.Center()
		w.ShowNormal()
	})
}

type MainWindow struct {
	*tk.Window
	edit   *tk.Entry
	number string
}

func (w *MainWindow) inputNumber(s string) {
	w.number += s
	w.edit.SetText(w.number)
}

func (w *MainWindow) inputCmd(s string) {

}

func NewWindow() *MainWindow {
	mw := &MainWindow{}
	mw.Window = tk.MainWindow()

	mw.edit = tk.NewEntry(mw)
	mw.edit.SetAlignment(tk.AlignmentRight)

	var number [10]*tk.Button
	for i := 0; i < 10; i++ {
		text := fmt.Sprintf("%v", i)
		number[i] = tk.NewButton(mw, text)
		number[i].SetWidth(2)
		number[i].OnCommand(func() {
			mw.inputNumber(text)
		})
	}
	number[0].SetWidth(8)

	var decimal *tk.Button
	decimal = tk.NewButton(mw, ".")
	decimal.SetWidth(2)
	decimal.OnCommand(func() {
		mw.inputNumber(".")
	})

	var clear *tk.Button
	clear = tk.NewButton(mw, "C")
	clear.SetWidth(2)

	var sig *tk.Button
	sig = tk.NewButton(mw, "±")
	sig.SetWidth(2)

	var precent *tk.Button
	precent = tk.NewButton(mw, "%")
	precent.SetWidth(2)

	// + - * /
	var plus *tk.Button
	var minus *tk.Button
	var multiple *tk.Button
	var division *tk.Button
	var equal *tk.Button
	plus = tk.NewButton(mw, "+")
	plus.SetWidth(2)
	minus = tk.NewButton(mw, "-")
	minus.SetWidth(2)
	multiple = tk.NewButton(mw, "×")
	multiple.SetWidth(2)
	division = tk.NewButton(mw, "÷")
	division.SetWidth(2)
	equal = tk.NewButton(mw, "=")
	equal.SetWidth(2)

	grid := tk.NewGridLayout(mw)
	grid.AddWidgets(clear, sig, precent, division)
	grid.AddWidgets(number[7], number[8], number[9], multiple)
	grid.AddWidgets(number[4], number[5], number[6], minus)
	grid.AddWidgets(number[1], number[2], number[3], plus)
	grid.AddWidget(number[0], tk.GridAttrColumnSpan(2))
	grid.AddWidget(decimal, tk.GridAttrRow(4), tk.GridAttrColumn(2))
	grid.AddWidget(equal, tk.GridAttrRow(4), tk.GridAttrColumn(3))

	hbox := tk.NewHPackLayout(mw)
	hbox.AddWidget(tk.NewLayoutSpacer(mw, 0, true), tk.PackAttrFillX(), tk.PackAttrExpand(true))
	hbox.AddLayout(grid, tk.PackAttrExpand(false), tk.PackAttrAnchor(tk.AnchorWest))

	vbox := tk.NewVPackLayout(mw)
	vbox.AddWidget(mw.edit, tk.PackAttrFillX(), tk.PackAttrExpand(true), tk.PackAttrAnchor(tk.AnchorNorth))
	vbox.AddLayout(hbox, tk.PackAttrFillX(), tk.PackAttrExpand(true), tk.PackAttrAnchor(tk.AnchorNorth))
	vbox.AddWidget(tk.NewLayoutSpacer(mw, 0, true), tk.PackAttrFillY(), tk.PackAttrExpand(true))

	return mw
}
