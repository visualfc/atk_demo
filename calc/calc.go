// Go语言编写的科学计算器，支持四则运算，支持括号，支持自定义函数。

package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/visualfc/atk/tk"
)

func main() {
	tk.MainLoop(func() {
		w := NewWindow()
		w.SetTitle("Go 计算器 Copyright(C) 2018 QJH")
		w.Center()
		w.ShowNormal()
	})
}

type MainWindow struct {
	*tk.Window
	edit    *tk.Entry // 计算输入
	eval    *MathEval // 数学计算
	cflag   int       // 0:输入状态 1:完成计算 2:切换算式/结果
	express string    // 计算输入存储
	result  string    // 计算结果存储
}

//计算结果并显示计算式和得数
//如果完成计算继续按回车则只显示得数
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
	express = strings.TrimSpace(express)
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

func (w *MainWindow) inputSymbol(s string) {
	w.edit.SetFocus()
	switch s {
	case "C":
		w.edit.SetText("")
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
	mw.Window = tk.MainWindow()
	font := tk.NewUserFont("", 16).SetBold(true)
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

	mw.eval = NewMathEval()
	mw.eval.RegistrFunc1("pow2", func(v float64) float64 {
		return v * v
	})
	mw.eval.RegistrFunc1("pow3", func(v float64) float64 {
		return v * v * v
	})
	mw.eval.RegistrFunc1("sqrt", func(v float64) float64 {
		return math.Sqrt(v)
	})
	mw.eval.RegistrFunc1("cbrt", func(v float64) float64 {
		return math.Cbrt(v)
	})
	return mw
}

//数学计算库
type MathEval struct {
}

//注册一个参数的计算函数, 如 pow2 => pow2(10)
func (m *MathEval) RegistrFunc1(name string, fn func(float64) float64) {
	tk.MainInterp().CreateCommand("tcl::mathfunc::"+name, func(args []string) (string, error) {
		if len(args) != 1 {
			return "", errors.New("Invalid param")
		}
		v, err := strconv.ParseFloat(args[0], 0)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%v", fn(v)), nil
	})
}

func (m *MathEval) Eval(express string) (string, error) {
	re := regexp.MustCompile("([\\d\\.]+)([a-zπ]+)")
	express = strings.ToLower(express)
	express = re.ReplaceAllStringFunc(express, func(s string) string {
		ar := re.FindStringSubmatch(s)
		if ar[2] == "e" {
			return s
		}
		return ar[1] + "*" + ar[2]
	})
	express = strings.NewReplacer("π", "3.14159265", "pi", "3.14159265", "%", "/100", "×", "*", "x", "*", "÷", "/").Replace(express)
	r, err := tk.MainInterp().EvalAsString(fmt.Sprintf("expr [string map {/ *1.0/} %v]", express))
	v, err := strconv.ParseFloat(r, 0)
	if err != nil {
		return "无效", err
	}
	if v > 1e-15 {
		v, _ = strconv.ParseFloat(fmt.Sprintf("%.15f", v), 0)
	}
	return fmt.Sprintf("%v", v), nil
}

func NewMathEval() *MathEval {
	return &MathEval{}
}
