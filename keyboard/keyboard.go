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
	w := tk.MainWindow()
	key := tk.NewLabel(w, "keyboard")
	pos := tk.NewLabel(w, "pos")
	key.SetFont(tk.NewUserFont("", 24))

	vpk := tk.NewVPackLayout(w)
	vpk.AddWidget(key, tk.PackAttrFillBoth(), tk.PackAttrExpand(true))
	vpk.AddWidget(pos, tk.PackAttrAnchor(tk.AnchorSouthEast))

	win := &Window{w, key, pos}
	win.BindKeyEvent(win.OnKeyEvent)
	win.BindEvent(event.Motion, win.OnMotion)
	win.BindEvent(event.ButtonPress, win.OnButtonClick)
	win.BindEvent("<Double-ButtonPress>", win.OnButtonDbclick)

	return win
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
