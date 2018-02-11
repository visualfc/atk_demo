package main

import (
	"fmt"

	"github.com/visualfc/atk/tk"
)

type Window struct {
	*tk.Window
}

func (w *Window) LayoutId() string {
	return w.Id()
}

func NewWindow() *Window {
	mw := &Window{}
	mw.Window = tk.MainWindow()
	info := tk.NewLabel(mw, "Hello ATL")
	vpk := tk.NewVPackLayout(mw)
	vpk.AddWidget(info)
	return mw
}

func main() {
	tk.MainLoop(func() {
		w := NewWindow()
		w.SetTitle("ATL Sample")
		w.Center()
		w.ResizeN(300, 200)
		w.ShowNormal()
		fmt.Println(tk.DumpWidget(w))
	})
}
