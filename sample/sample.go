package main

import (
	"fmt"

	"github.com/visualfc/atk/tk"
)

type Window struct {
	*tk.Window
}

func NewWindow() *Window {
	mw := &Window{}
	mw.Window = tk.RootWindow()
	info := tk.NewLabel(mw, "Hello ATL")
	vpk := tk.NewVPackLayout(mw)
	vpk.AddWidget(info)
	tk.MainInterp().Eval(`proc every {ms body} {eval $body; after $ms [info level 0]}        
 pack [label .clock -textvar time]                                
 every 1000 {set ::time [clock format [clock sec] -format %H:%M:%S]}`)
	//(".clock")
	//info2 := tk.FindWidgetInfo(".clock")
	//lbl := tk.NewLabel(mw, "ok")
	lbl := &tk.Label{}
	lbl.Attach(".clock")
	lbl.SetForground("red")
	fmt.Println(lbl)
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
