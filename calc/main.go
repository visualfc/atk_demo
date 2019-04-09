// Go语言编写的科学计算器，支持四则运算，支持括号，支持自定义函数。

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/visualfc/atk/tk"
)

func main() {
	tk.MainLoop(func() {
		w := NewWindow()
		w.SetTitle("Go 计算器 Copyright(C) 2018 visualfc")
		w.Center()
		w.ShowNormal()
	})
}

// 主窗口
type MainWindow struct {
	*tk.Window
	edit    *tk.Entry // 计算输入
	eval    *MathEval // 数学计算
	cflag   int       // 0:输入状态 1:完成计算 2:切换算式/结果
	express string    // 计算表达式存储
	result  string    // 计算结果存储
}

//计算，保存表达式和结果
func (w *MainWindow) Equals() {
	if w.cflag == 2 {
		w.edit.SetText(fmt.Sprintf("%v = %v", w.express, w.result))
		w.edit.SetCursorPosition(utf8.RuneCountInString(w.express))
		w.cflag = 1
		return
	}
	express := w.edit.Text()
	pos := strings.Index(express, "=")
	if pos != -1 {
		if w.cflag == 1 {
			w.edit.SetText(express[pos+2:])
			w.edit.SetCursorPosition(w.edit.TextLength())
			w.cflag = 2
			return
		}
		express = express[:pos]
	}
	express = strings.Replace(express, " ", "", -1)
	if len(express) == 0 {
		return
	}
	result, err := w.eval.Eval(express)
	w.express = express
	w.result = result
	w.edit.SetText(fmt.Sprintf("%v = %v", express, result))
	if err != nil {
		w.edit.SetForeground("red")
		w.cflag = 0
	} else {
		w.edit.SetForeground("black")
		w.cflag = 1
	}
}

// 输入
func (w *MainWindow) inputSymbol(s string) {
	w.edit.SetFocus()
	switch s {
	case "C":
		w.edit.SetText("")
		w.cflag = 2
	case "⇦":
		w.edit.Delete(w.edit.CursorPosition() - 1)
	case "=":
		w.Equals()
	default:
		switch s {
		case "π":
			s = "π"
		case "平方":
			s = "pow2("
		case "立方":
			s = "pow3("
		case "平方根":
			s = "sqrt("
		case "立方根":
			s = "cbrt("
		}
		w.edit.Insert(w.edit.CursorPosition(), s)
	}
}

var (
	symbolsList = []string{
		"C π ( ) ⇦",
		"平方 7 8 9 ÷",
		"立方 4 5 6 ×",
		"平方根 1 2 3 -",
		"立方根 0 . = +",
	}
)

func NewWindow() *MainWindow {
	mw := &MainWindow{}
	mw.Window = tk.RootWindow()
	font := tk.NewUserFont(bestFont(), 16).SetBold(true)
	mw.edit = tk.NewEntry(mw)
	mw.edit.SetFont(font)
	mw.edit.SetJustify(tk.JustifyRight)
	mw.edit.SetExportSelection(false)
	mw.BindEvent("<Key-Escape>", func(e *tk.Event) {
		mw.inputSymbol("C")
	})
	mw.BindEvent("<Key-Return>", func(e *tk.Event) {
		mw.Equals()
	})
	mw.edit.OnUpdate(func() {
		mw.cflag = 0
		text := mw.edit.Text()
		newText := strings.NewReplacer("*", "×", "/", "÷", "Pi", "π", "PI", "π", "pi", "π", "pI", "π").Replace(text)
		mw.edit.SetText(newText)
		count, newCount := utf8.RuneCountInString(text), utf8.RuneCountInString(newText)
		if count != newCount {
			mw.edit.SetCursorPosition(mw.edit.CursorPosition() + newCount - count)
		}
	})
	mw.edit.SetFocus()

	grid := tk.NewGridLayout(mw)
	grid.AddWidgetEx(mw.edit, 0, 0, 1, 5, tk.StickyAll)
	for _, syms := range symbolsList {
		var buttons []tk.Widget
		for _, sym := range strings.Split(syms, " ") {
			btn := tk.NewButton(mw, sym)
			btn.OnCommand(func() {
				mw.inputSymbol(btn.Text())
			})
			buttons = append(buttons, btn)
		}
		grid.AddWidgetList(buttons, tk.GridAttrSticky(tk.StickyAll))
	}
	grid.SetColumnAttr(-1, 0, 1, "")
	grid.SetRowAttr(-1, 0, 1, "")
	grid.SetBorderWidth(5)

	//注册自定义计算函数
	mw.eval = NewMathEval()
	mw.eval.RegistrFunc1("pow2", func(v float64) float64 {
		return v * v
	})
	mw.eval.RegistrFunc1("pow3", func(v float64) float64 {
		return v * v * v
	})
	mw.eval.RegistrFunc1("sqrt", func(v float64) float64 {
		return Sqrt(v)
	})
	mw.eval.RegistrFunc1("cbrt", func(v float64) float64 {
		return Cbrt(v)
	})
	return mw
}
