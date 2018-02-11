package main

import (
	"fmt"

	"github.com/visualfc/atk/tk"
)

type Window struct {
	*tk.Window
}

func NewWindow() *Window {
	mw := tk.MainWindow()

	split := tk.NewSplitter(mw, tk.Vertical) //, tk.WidgetAttrInitUseTheme(false))
	lst := tk.NewListBoxEx(mw)
	//lst.ShowXScrollBar(false)

	var items []string
	for i := 0; i < 100; i++ {
		items = append(items, fmt.Sprintf("%v-123456789-123456789-123456789", i))
	}
	lst.SetItems(items)

	vbox := tk.NewVPackLayout(mw)
	//	vbox.AddWidget(tk.NewLabel(mw, "ListBox Demo"))
	//	vbox.AddWidget(lst)
	//split.SetHeight(200).SetWidth(200)
	//	vbox.AddWidget(split)
	split.AddWidget(lst, 2)
	split.AddWidget(tk.NewLabel(mw, "Demo"), 0)
	//split.AddWidget(tk.NewLabel(split, "Demo"), 0)
	vbox.AddWidget(split, tk.PackAttrFillBoth(), tk.PackAttrExpand(true))
	lst.OnSelectionChanged(func() {
		fmt.Println(lst.SelectedItems())
	})
	return &Window{mw}
}

func main() {
	tk.MainLoop(func() {
		w := NewWindow()
		w.Center()
		w.ResizeN(300, 300)
		w.ShowNormal()
	})
}
