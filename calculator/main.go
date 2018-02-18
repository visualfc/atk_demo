package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/visualfc/atk/tk"
)

func main() {
	tk.MainLoop(func() {
		w := NewWindow()
		w.SetTitle("Go 计算器")
		w.Center()
		w.ShowNormal()
		//	w.SetResizable(false, false)
	})
}

type MainWindow struct {
	*tk.Window
	edit       *tk.Entry
	status     *tk.Label
	waitNext   bool
	number     string //数字记录
	operator   string //运算符 + - × ÷
	prevNumber float64
}

func (w *MainWindow) inputNumber(s string) {
	if w.waitNext {
		w.waitNext = false
		w.number = "0"
	}
	if s == "." && strings.Contains(w.number, ".") {
		return
	}
	if w.number == "0" || w.number == "-0" {
		if s == "0" {
			return
		} else if s != "." {
			if w.number == "0" {
				w.number = ""
			} else {
				w.number = "-"
			}
		}
	}
	w.number += s
	w.updateText()
}

func (w *MainWindow) updateText() {
	w.edit.SetText(w.number)
}

func (w *MainWindow) currentNumber() float64 {
	r, _ := strconv.ParseFloat(w.number, 0)
	return r
}

func (w *MainWindow) calculate(operator string) {
	var divError bool
	if w.operator == "" {
		w.prevNumber = w.currentNumber()
	} else if !w.waitNext {
		switch w.operator {
		case "+":
			w.prevNumber += w.currentNumber()
		case "-":
			w.prevNumber -= w.currentNumber()
		case "×", "*":
			w.prevNumber *= w.currentNumber()
		case "÷", "/":
			if w.currentNumber() == 0 {
				divError = true
			} else {
				w.prevNumber /= w.currentNumber()
			}
		}
		w.number = fmt.Sprintf("%v", w.prevNumber)
	}
	w.operator = operator
	w.waitNext = true
	if divError {
		w.status.SetText("被除数不能为 0 !")
	} else if operator == "=" {
		w.status.SetText(fmt.Sprintf("%v", w.prevNumber))
	} else {
		w.status.SetText(fmt.Sprintf("%v %v", w.prevNumber, w.operator))
	}
	w.updateText()
}

func (w *MainWindow) inputSymbol(s string) {
	switch s {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ".":
		w.inputNumber(s)
	case "π":
		w.number = fmt.Sprintf("%v", math.Pi)
		w.waitNext = false
		w.updateText()
	case "+/-":
		if w.currentNumber() != 0 {
			w.number = fmt.Sprintf("%v", -1*w.currentNumber())
			w.waitNext = false
			w.updateText()
		}
	case "%":
		if w.currentNumber() != 0 {
			w.number = fmt.Sprintf("%v", w.currentNumber()/100)
			w.waitNext = false
			w.updateText()
		}
	case "C", "c": //清除
		w.number = "0"
		w.operator = ""
		w.waitNext = true
		w.updateText()
	case "+", "-", "×", "÷", "*", "/":
		w.calculate(s)
	case "=":
		w.calculate("=")
		w.operator = ""
		w.waitNext = true
	}
}

var (
	symbols = []string{
		"C", "+/-", "%", "÷",
		"7", "8", "9", "×",
		"4", "5", "6", "-",
		"1", "2", "3", "+",
		"0", "π", ".", "="}
)

func NewWindow() *MainWindow {
	mw := &MainWindow{}
	mw.Window = tk.MainWindow()

	font := tk.LoadSysFont(tk.SysTextFont).Clone()
	font.SetSize(24).SetBold(true)

	mw.edit = tk.NewEntry(mw)
	mw.edit.SetAlignment(tk.AlignmentRight)
	mw.edit.SetFont(font)
	mw.edit.SetState(tk.StateReadOnly)
	mw.edit.SetText("0")
	mw.status = tk.NewLabel(mw, "")
	mw.number = "0"

	mw.BindEvent("<Key>", func(e *tk.Event) {
		if e.KeySym == "Return" {
			mw.inputSymbol("=")
		} else {
			mw.inputSymbol(e.KeyText)
		}
	})

	grid := tk.NewGridLayout(mw)
	for n, sym := range symbols {
		btn := tk.NewButton(mw, sym)
		btn.SetWidth(5)
		btn.OnCommand(func() {
			mw.inputSymbol(btn.Text())
		})
		row := n / 4
		col := n % 4
		grid.AddWidgetEx(btn, row, col, 1, 1, tk.StickyAll)
	}
	grid.SetColumnAttr(-1, 0, 1, "")
	grid.SetRowAttr(-1, 0, 1, "")

	vbox := tk.NewVPackLayout(mw)
	vbox.SetBorderWidth(4)
	vbox.AddWidgetEx(mw.status, tk.FillX, false, tk.AnchorWest)
	vbox.AddWidgetEx(mw.edit, tk.FillX, false, 0)
	vbox.AddWidgetEx(grid, tk.FillBoth, true, 0)
	return mw
}
