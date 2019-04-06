package main

import (
	"fmt"

	"github.com/visualfc/atk/tk"
)

func main() {
	tk.MainLoop(func() {
		mw := tk.RootWindow()
		hbox := tk.NewHPackLayout(mw)
		vbox := tk.NewVPackLayout(mw)

		lbl := tk.NewLabel(mw, "Label")
		btn := tk.NewButton(mw, "Button")
		frm := tk.NewFrame(mw)
		frm.SetReliefStyle(tk.ReliefStyleFlat)
		frm.SetBorderWidth(5)
		lbl2 := tk.NewLabel(mw, "lable2")

		hbox.AddWidget(lbl)
		hbox.AddWidget(btn)

		vbox.AddWidget(hbox, tk.PackAttrExpand(false))
		vbox.AddWidget(frm) //, tk.PackAttrFillBoth(), tk.PackAttrExpand(true))
		vbox.AddWidget(lbl2)

		hbox3 := tk.NewHPackLayout(frm)
		hbox3.AddWidget(tk.NewButton(frm, "OK"))
		hbox3.AddWidget(tk.NewButton(frm, "OK2"))
		hbox3.SetPaddingN(10, 10)

		//vbox.Master().(*tk.LayoutFrame).SetPaddingN(10, 10)
		vbox.SetBorderWidth(10)
		vbox.SetPaddingN(10, 10)
		vbox.Repack()
		mw.Center()
		mw.ShowNormal()
		fmt.Println(tk.DumpWidget(mw))
	})
}
