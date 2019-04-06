package main

import (
	"log"

	"github.com/visualfc/atk/tk"
)

func main() {
	tk.MainLoop(func() {
		mw := NewWindow()
		mw.Center()
		mw.ResizeN(400, 300)
		mw.ShowNormal()
	})
}

type Window struct {
	*tk.Window
	tab *tk.Notebook
}

func NewWindow() *Window {
	mw := &Window{}
	mw.Window = tk.RootWindow()
	mw.tab = tk.NewNotebook(mw)

	page1 := tk.NewButton(mw, "OK1")
	button2 := tk.NewButton(mw, "OK2")

	page2 := tk.NewVPackLayout(mw)
	page2.AddWidgetEx(button2, tk.FillNone, true, tk.AnchorSouthWest)

	page3 := tk.NewListBoxEx(mw)

	mw.tab.AddTab(page1, "page1")
	mw.tab.AddTab(page2, "page2")
	mw.tab.AddTab(page3, "page3")
	mw.tab.SetTab(page2, "page2-change")
	mw.tab.SetCurrentTab(page1)

	pane4 := tk.NewLabel(mw, "Image")
	img, err := tk.LoadImage(tk.TkLibrary() + "/images/pwrdLogo200.gif")
	if err != nil {
		log.Println(err)
	}
	pane4.SetImage(img)

	mw.tab.AddTab(pane4, "page4")

	vbox := tk.NewVPackLayout(mw)
	vbox.AddWidgetEx(mw.tab, tk.FillBoth, true, 0)
	return mw
}
