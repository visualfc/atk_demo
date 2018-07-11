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
		mw := tk.MainWindow()
		mw.SetTitle("Go Tk")
		mw.ShowNormal()

		img, err := tk.LoadImage("./liteide400.png")
		if err != nil {
			log.Println(err)
		} else {
			lbl := tk.NewLabel(mw, "Image")
			lbl.SetImage(img)
			tk.NewVPackLayout(mw).AddWidget(lbl)
			mw.SetSize(img.Size())
		}
	})
}
