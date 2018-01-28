package main

import (
	"fmt"
	"log"
	"time"

	"github.com/visualfc/atk/event"
	"github.com/visualfc/atk/tk"
)

type Event struct {
	Widget tk.Widget
	Time   time.Time
}

type MouseEvent struct {
	Event
	Button int
	PosX   int
	PosY   int
}

type WhellEvent struct {
	MouseEvent
	Whell int
}

type KeyEvent struct {
	Event
	Key  int
	Text string
}

func main() {
	if tk.Init() != nil {
		log.Fatalln("load tk false")
	}
	tk.MainLoop(func() {
		mw := tk.MainWindow()
		info := tk.NewLabel(mw, "keyboard")
		info.SetFont(tk.NewUserFont("", 36))
		pos := tk.NewLabel(mw, "pos:")
		vpk := tk.NewVPackLayout(mw)
		vpk.AddWidget(info, tk.PackAttrFillY(), tk.PackAttrExpand(true), tk.PackAttrAnchor(tk.AnchorWest))
		vpk.AddWidget(pos, tk.PackAttrAnchor(tk.AnchorSouthEast))

		mw.BindEvent(event.KeyPress, func(e *tk.Event) {
			info.SetText(fmt.Sprintf("Scan:\t0x%x\nSym:\t%v\nText:\t%v", e.KeyScanCode, e.KeySym, e.KeyText))
		})
		mw.BindEvent(event.Motion, func(e *tk.Event) {
			pos.SetText(fmt.Sprintf("global: %v-%v\tpos: %3d-%3d", e.GlobalPosX, e.GlobalPosY, e.PosX, e.PosY))
		})
		mw.BindEvent(event.ButtonPress, func(e *tk.Event) {
			pos.SetText(fmt.Sprintf("mouse click: %v\tpos: %v-%v", e.MouseButton, e.PosX, e.PosY))
		})
		mw.BindEvent("<Double-ButtonPress>", func(e *tk.Event) {
			pos.SetText(fmt.Sprintf("mouse double click: %v\tpos: %v-%v", e.MouseButton, e.PosX, e.PosY))
		})

		mw.SetSizeN(400, 300)
		mw.Center()
		mw.ShowNormal()
	})
}
