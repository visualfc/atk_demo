package main

import (
	"fmt"

	"github.com/visualfc/atk/tk"
)

type Window struct {
	*tk.Window
}

type Base struct {
	id string
}

func (b *Base) Id() string {
	return b.id
}

type M1 struct {
	Base
}

type M2 struct {
	Base
}

func (m *M2) Id() string {
	return m.id
}

type My struct {
	*M1
	*M2
}

func NewWindow() *Window {

	m := &My{&M1{Base{"m1"}}, &M2{Base{"m2"}}}
	fmt.Println(m.Id())

	mw := tk.MainWindow()

	//	paned := tk.NewPaned(mw, tk.Vertical)
	//, tk.WidgetAttrInitUseTheme(false))
	lst := tk.NewListBoxEx(mw)
	//lst.ShowXScrollBar(false)

	var items []string
	for i := 0; i < 100; i++ {
		items = append(items, fmt.Sprintf("%v-123456789-123456789-123456789", i))
	}
	lst.SetItems(items)

	vbox := tk.NewVPackLayout(mw)
	vbox.AddWidget(tk.NewLabel(mw, "ListBox Demo"))
	vbox.AddWidget(lst)
	//split.SetHeight(200).SetWidth(200)
	//	vbox.AddWidget(split)
	//	paned.AddWidget(lst, 0)
	//	paned.AddWidget(tk.NewLabel(paned, "Demo"), 0)
	//	paned.InsertWidget(0, tk.NewLabel(paned, "Test"), 0)
	//	paned.RemovePane(2)
	//fmt.Println(tk.MainInterp().EvalAsString(fmt.Sprintf("%v sashpos 1 0", split.Id())))
	//split.AddWidget(tk.NewLabel(split, "Demo"), 0)
	//vbox.AddWidget(paned, tk.PackAttrFillBoth(), tk.PackAttrExpand(true))
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
