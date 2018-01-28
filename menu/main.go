package main

import (
	"fmt"
	"log"

	"github.com/visualfc/atk/tk"
)

func main() {
	if tk.Init() != nil {
		log.Fatalln("load tk false")
	}
	tk.MainLoop(func() {
		mw := tk.MainWindow()

		mbar := tk.NewMenu(mw)
		mw.SetMenu(mbar)

		file := mbar.AddNewSubMenu("File")
		file.AddAction(tk.NewAction("Open", func() {
			fmt.Println("Open")
		}))
		file.AddSeparator()
		file.AddAction(tk.NewAction("Quit", func() {
			tk.Quit()
		}))

		copyAct := tk.NewAction("Copy", func() {
			fmt.Println("copy")
		})
		pasteAct := tk.NewAction("Paste", func() {
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

		boldAct := tk.NewCheckAction("Bold", nil)
		boldAct.OnCommand(func() {
			fmt.Println("bold", boldAct.IsChecked())
		})
		italicAct := tk.NewCheckAction("Italic", nil)
		fmt.Println(italicAct.IsChecked())
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
file.AddAction(tk.NewAction("Open", func() {
	fmt.Println("Open")
}))

//action
tk.NewAction
tk.NewCheckAction
tk.NewActionGroup //radio action group

`)
		btn := tk.NewButton(mw, "Quit")
		btn.OnCommand(func() {
			//tk.Quit()
			tk.PopupMenu(file, mw.Pos().X, mw.Pos().Y)
		})
		vpk := tk.NewVPackLayout(mw)
		vpk.SetBorderWidth(10)
		vpk.AddWidget(tk.NewLayoutSpacer(mw, 0, true))
		vpk.AddWidget(btn)
		vpk.InsertWidget(0, info)

		ma := tk.NewMenuRole(mbar, "system")
		mbar.AddSubMenu("apple", ma)
		ma.AddAction(tk.NewAction("OK", nil))

		mw.SetTitle("Menu Demo")
		mw.Center()
		mw.ShowNormal()
		fmt.Println(tk.DumpWidget(mw))
	})
}
