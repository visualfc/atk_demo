package main

import (
	"fmt"
	"time"

	"github.com/visualfc/atk/tk"
)

type Window struct {
	*tk.Window
}

func NewWindow() *Window {
	mw := &Window{}
	mw.Window = tk.RootWindow()
	info := tk.NewProgressBar(mw, tk.Horizontal)
	//info.SetLength(200)
	vpk := tk.NewVPackLayout(mw)
	btnStart := tk.NewButton(mw, "Start")
	btnStop := tk.NewButton(mw, "stop")
	btnStart.OnCommand(func() {
		//info.StartEx(1e9)
		info.Start()
	})
	info.SetMaximum(100)
	info.SetValue(10)
	btnStop.OnCommand(func() {
		//info.Pause()
		info.Stop()
	})
	//info.SetStep(0.01)
	//info.SetDeterminateMode(false)
	tick := time.NewTicker(1e6)
	quit := make(chan bool)
	go func() {
		var last int
		for {
			select {
			case <-tick.C:
				tk.Async(func() {
					if last != info.Phase() {
						last = info.Phase()
						mw.SetTitle(fmt.Sprintf("%d-%v", last, info.Value()))
					}
				})
			case <-quit:
				return
			}
		}
	}()
	mw.OnClose(func() bool {
		quit <- true
		return true
	})
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
