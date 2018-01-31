package main

import (
	"strings"

	"github.com/visualfc/atk/tk"
)

type Window struct {
	*tk.Window
}

func NewWindow() *Window {
	w := tk.MainWindow()
	//button
	btn1 := tk.NewButton(w, "Button1")
	btn2 := tk.NewButton(w, "Button2")
	btn3 := tk.NewButton(w, "Button3")
	btnInfo := tk.NewLabel(w, "Info")
	fninfo := func(btn *tk.Button) {
		btnInfo.SetText(btn.Text())
	}
	btn1.OnCommand(func() { fninfo(btn1) })
	btn2.OnCommand(func() { fninfo(btn2) })
	btn3.OnCommand(func() { fninfo(btn3) })

	hbox1 := tk.NewHPackLayout(w)
	hbox1.AddWidgets(btn1, btn2, btn3, tk.NewLayoutSpacer(w, 0, true), btnInfo)

	//check
	chk1 := tk.NewCheckButton(w, "Check1")
	chk2 := tk.NewCheckButton(w, "Check2")
	chk3 := tk.NewCheckButton(w, "Check3")
	chkInfo := tk.NewLabel(w, "Info")
	fnchk := func() {
		var info []string
		if chk1.IsChecked() {
			info = append(info, chk1.Text())
		}
		if chk2.IsChecked() {
			info = append(info, chk2.Text())
		}
		if chk3.IsChecked() {
			info = append(info, chk3.Text())
		}
		chkInfo.SetText(strings.Join(info, "&"))
	}
	chk1.OnCommand(fnchk)
	chk2.OnCommand(fnchk)
	chk3.OnCommand(fnchk)
	hbox2 := tk.NewHPackLayout(w)
	hbox2.AddWidgets(chk1, chk2, chk3, tk.NewLayoutSpacer(w, 0, true), chkInfo)

	//radio
	radio1 := tk.NewRadioButton(w, "Radio1")
	radio2 := tk.NewRadioButton(w, "Radio2")
	radio3 := tk.NewRadioButton(w, "Radio3")

	radioInfo := tk.NewLabel(w, "Info")
	group1 := tk.NewRadioGroup()
	group1.OnRadioChanged(func() {
		radioInfo.SetText(group1.CheckedRadio().Text())
	})
	group1.AddRadios(radio1, radio2, radio3)
	hbox3 := tk.NewHPackLayout(w)
	hbox3.AddWidgetList(group1.WidgetList())
	hbox3.AddWidgets(tk.NewLayoutSpacer(w, 0, true), radioInfo)
	radio2.Invoke()

	//menubutton & menu
	mbtn := tk.NewMenuButton(w, "MenuButton")
	mbtnDir := mbtn.Direction()
	mgroup := tk.NewActionGroup()
	var dirs []tk.Direction = []tk.Direction{tk.DirectionBelow, tk.DirectionAbove, tk.DirectionLeft, tk.DirectionRight}
	for _, v := range dirs {
		act := mgroup.AddNewRadioAction("Direction" + strings.Title(v.String()))
		act.SetData(v)
		if mbtnDir == v {
			act.SetChecked(true)
			mbtn.SetText(act.Label())
		}
	}
	mmenu := tk.NewMenu(w)
	mmenu.AddActions(mgroup.Actions())
	mbtn.SetMenu(mmenu)
	mgroup.OnCommand(func() {
		act := mgroup.CheckedAction()
		if act != nil {
			mbtn.SetDirection(act.Data().(tk.Direction))
			mbtn.SetText(act.Label())
		}
	})

	hbox4 := tk.NewHPackLayout(w)
	hbox4.AddWidget(mbtn)

	vbox := tk.NewVPackLayout(w)
	vbox.AddLayout(hbox1)
	vbox.AddLayout(hbox2)
	vbox.AddLayout(hbox3)
	vbox.AddLayout(hbox4)

	vbox.SetBorderWidth(10)

	return &Window{w}
}

func main() {
	tk.MainLoop(func() {
		w := NewWindow()
		w.Center()
		w.ResizeN(400, 300)
		w.ShowNormal()
	})
}