package main

import (
	"fmt"

	"github.com/visualfc/atk/tk"
)

type M struct {
	X, Y int
}

func main() {
	tk.MainLoop(func() {
		mw := tk.RootWindow()
		grid := tk.NewGridLayout(mw)
		vbox := tk.NewVPackLayout(mw)
		//		vpk := tk.NewVPackLayout(mw)
		//		hpk1 := tk.NewHPackLayout(mw)
		//		hpk1.AddWidget(tk.NewLabel(mw, "Label"))
		btn := tk.NewButton(mw, "Button")
		btn2 := tk.NewButton(mw, "button2")
		btn3 := tk.NewButton(mw, "button3")
		btn4 := tk.NewButton(mw, "button4")
		btn5 := tk.NewButton(mw, "button5dsfsf")
		grid.AddWidgets(btn, btn2, btn3)
		grid.AddWidgets(btn4, nil, btn5)
		//grid.AddWidgetEx(btn, 1, 1, 1, 1, tk.StickyN|tk.StickyE)
		grid.SetColumnAttr(3, 0, 1, "")
		btn6 := tk.NewButton(mw, "button6")
		btn7 := tk.NewButton(mw, "button7")
		vbox.AddWidget(btn6)
		vbox.AddWidget(btn7, tk.PackAttrFillBoth(), tk.PackAttrExpand(true))
		grid.AddWidget(vbox, tk.GridAttrRow(0), tk.GridAttrColumn(3), tk.GridAttrRowSpan(4), tk.GridAttrSticky(tk.StickyAll))
		grid.AddWidget(tk.NewButton(nil, "btn7"))
		vbox.AddWidget(grid)
		vbox.SetBorderWidth(10)
		grid.SetBorderWidth(20)
		//grid.RemoveLayout(vbox)
		//vbox.Repack()
		//vbox.RemoveLayout(grid)
		//grid.Repack()
		//		tk.Grid(btn, tk.GridAttrRowSpan(2), tk.GridAttrSticky(tk.StickyNS)) //, tk.GridAttrColumnSpan(3), tk.GridAttrSticky(tk.StickyE|tk.StickyW))
		//		tk.Grid(btn2, tk.GridAttrRow(0), tk.GridAttrColumn(1))
		//		tk.Grid(btn3, tk.GridAttrRow(1), tk.GridAttrColumn(1))
		//		tk.Grid(btn4)
		//		tk.Grid(btn5, tk.GridAttrColumn(2))
		//		//tk.GridRowIndex(mw, btn, tk.GridIndexAttrWeight(1))
		//		tk.GridColumnIndex(mw, 0, tk.GridIndexAttrWeight(1), tk.GridIndexUniform("a"))
		//		tk.GridColumnIndex(mw, 2, tk.GridIndexAttrWeight(1), tk.GridIndexUniform("a"))
		//btn.SetWidth(30)
		//		hpk1.AddWidget(btn) //, tk.PackAttrFillX(), tk.PackAttrExpand(true))
		//		vpk.AddLayout(hpk1)
		//		vpk.SetPaddingN(10, 10)
		//vpk.AddWidget(tk.NewFrame(mw), tk.PackAttrFillBoth(), tk.PackAttrExpand(true))

		//create menu
		m := tk.NewMenu(mw)
		mw.SetMenu(m)

		sub := tk.NewMenu(m)
		m.AddSubMenu("Theme", sub)
		ag := tk.NewActionGroup()
		for _, id := range tk.TtkTheme.ThemeIdList() {
			act := ag.AddNewRadioAction(id)
			if id == tk.TtkTheme.ThemeId() {
				act.SetChecked(true)
			}
		}
		ag.OnCommand(func() {
			act := ag.CheckedAction()
			tk.TtkTheme.SetThemeId(act.Label())
		})
		sub.AddSeparator()
		sub.AddActions(ag.Actions())

		mw.Center()
		mw.SetSizeN(400, 300)
		mw.ShowNormal()

		fmt.Println(tk.DumpWidget(mw))
	})
}
