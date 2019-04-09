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
		btn5 := tk.NewButton(mw, "button5")
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
