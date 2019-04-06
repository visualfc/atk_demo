package main

import (
	"fmt"

	"github.com/visualfc/atk/event"
	"github.com/visualfc/atk/tk"
)

type Window struct {
	*tk.Window
	keyLabel *tk.Label
	posLabel *tk.Label
}

func NewWindow() *Window {
	mw := &Window{}
	mw.Window = tk.RootWindow()
	mw.keyLabel = tk.NewLabel(mw, "keyboard")
	mw.posLabel = tk.NewLabel(mw, "pos")
	mw.keyLabel.SetFont(tk.NewUserFont("", 24))

	vpk := tk.NewVPackLayout(mw)
	vpk.AddWidget(mw.keyLabel, tk.PackAttrFillBoth(), tk.PackAttrExpand(true))
	vpk.AddWidget(mw.posLabel, tk.PackAttrAnchor(tk.AnchorSouthEast))

	mw.BindKeyEvent(mw.OnKeyEvent)
	mw.BindEvent(event.Motion, mw.OnMotion)
	mw.BindEvent(event.ButtonPress, mw.OnButtonClick)
	mw.BindEvent("<Double-ButtonPress>", mw.OnButtonDbclick)

	return mw
}

func (w *Window) OnKeyEvent(e *tk.KeyEvent) {
	w.keyLabel.SetText(fmt.Sprintf("Code:\t0x%x\nSym:\t%v\nText:\t%v\nRune:\t%x\nModifier:\t%v",
		e.KeyCode, e.KeySym, e.KeyText, e.KeyRune, e.KeyModifier))
}

func (w *Window) OnMotion(e *tk.Event) {
	w.posLabel.SetText(fmt.Sprintf("global: %v-%v\tpos: %3d-%3d", e.GlobalPosX, e.GlobalPosY, e.PosX, e.PosY))
}

func (w *Window) OnButtonClick(e *tk.Event) {
	w.posLabel.SetText(fmt.Sprintf("mouse click: %v\tpos: %v-%v", e.MouseButton, e.PosX, e.PosY))
}

func (w *Window) OnButtonDbclick(e *tk.Event) {
	w.posLabel.SetText(fmt.Sprintf("mouse double click: %v\tpos: %v-%v", e.MouseButton, e.PosX, e.PosY))
}

func main() {
	tk.MainLoop(func() {
		mw := NewWindow()
		mw.SetTitle("ATK Keyboard Demo")
		mw.SetSizeN(400, 300)
		mw.Center()
		mw.ShowNormal()
	})
}
