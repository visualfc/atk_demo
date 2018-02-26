package main

import (
	"github.com/visualfc/atk/tk"
)

type Window struct {
	*tk.Window
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
	})
}
