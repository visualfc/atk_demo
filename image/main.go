package main

import (
	"log"

	"github.com/visualfc/atk/tk"
)

func main() {
	if err := tk.Init(); err != nil {
		log.Fatalln(err)
	}

	tk.MainLoop(func() {
		mw := tk.RootWindow()
		mw.SetTitle("Go Tk")
		mw.ShowNormal()

		img, err := tk.LoadImage("./liteide400.png")
		if err != nil {
			log.Println(err)
		} else {
			lbl := tk.NewLabel(mw, "Image")
			lbl.SetText("Image")
			btn := tk.NewButton(mw, "Btn")
			btn.OnCommand(func() {
				lbl.SetImage(img)
			})
			tk.NewVPackLayout(mw).AddWidgetList([]tk.Widget{btn, lbl})
			mw.SetSize(img.Size())
		}
	})
}
