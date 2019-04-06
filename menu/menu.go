package main

import (
	"fmt"

	"github.com/visualfc/atk/tk"
)

func main() {
	tk.MainLoop(func() {
		mw := tk.RootWindow()

		mbar := tk.NewMenu(mw)
		mw.SetMenu(mbar)

		file := mbar.AddNewSubMenu("File")
		file.AddAction(tk.NewActionEx("Open", func() {
			fmt.Println("Open")
		}))
		file.AddSeparator()
		file.AddAction(tk.NewActionEx("Quit", func() {
			tk.Quit()
		}))

		copyAct := tk.NewActionEx("Copy", func() {
			fmt.Println("copy")
		})
		pasteAct := tk.NewActionEx("Paste", func() {
			fmt.Println("paste")
		})
		group := tk.NewActionGroup()
		group.AddNewRadioAction("Left")
		group.AddNewRadioAction("Center")
		group.AddNewRadioAction("Right")
		group.SetCheckedIndex(1)

		group.OnCommand(func() {
			fmt.Println(group.CheckedActionIndex())
		})

		boldAct := tk.NewCheckAction("Bold")
		boldAct.OnCommand(func() {
			fmt.Println("bold", boldAct.IsChecked())
		})
		italicAct := tk.NewCheckAction("Italic")
		italicAct.OnCommand(func() {
			fmt.Println("italic", italicAct.IsChecked())
		})

		edit := mbar.AddNewSubMenu("Edit")
		edit.AddAction(copyAct)
		edit.AddAction(pasteAct)
		edit.AddSeparator()
		edit.AddActions(group.Actions())
		edit.AddSeparator()
		edit.AddAction(boldAct)
		edit.AddAction(italicAct)

		info := tk.NewLabel(mw, "")
		info.SetText(`Menu Demo

//menu
mbar := tk.NewMenu(mw)
mw.SetMenu(mbar)

//sub menu
file := mbar.AddNewSubMenu("File")
file.AddAction(tk.NewActionEx("Open", func() {
	fmt.Println("Open")
}))

//action
tk.NewActionEx
tk.NewCheckAction
tk.NewActionGroup //radio action group

`)
		btn := tk.NewButton(mw, "Quit")
		btn.OnCommand(func() {
			tk.Quit()
		})
		vpk := tk.NewVPackLayout(mw)
		vpk.SetBorderWidth(10)
		vpk.AddWidget(tk.NewLayoutSpacer(mw, 0, true))
		vpk.AddWidget(btn)
		vpk.InsertWidget(0, info)

		mw.SetTitle("Menu Demo")
		mw.Center()
		mw.ShowNormal()
	})
}
