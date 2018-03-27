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
	mw.Window = tk.MainWindow()
	info := tk.NewProgressBar(mw, tk.Horizontal)
	//info.SetLength(200)
	vpk := tk.NewVPackLayout(mw)
	btnStart := tk.NewButton(mw, "Start")
	btnStop := tk.NewButton(mw, "stop")
	btnStart.OnCommand(func() {
		info.StartEx(100)
	})
	info.SetMaximum(100)
	btnStop.OnCommand(func() {
		//info.Pause()
		info.Stop()
	})
	info.SetStep(0.01)
	info.SetDeterminateMode(true)
	fmt.Println(info.IsDeterminateMode())
	//	go func() {
	//		tick := time.Tick(1e6)
	//		last := info.Phase()
	//		for {
	//			select {
	//			case <-tick:
	//				tk.Async(func() {
	//					if last != info.Phase() {
	//						last = info.Phase()
	//						fmt.Println(last)
	//					}
	//				})
	//			}
	//		}
	//	}()
	vpk.AddWidget(info, tk.PackAttrFillX())
	vpk.AddWidget(btnStart)
	vpk.AddWidget(btnStop)
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
