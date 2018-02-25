package main

import (
	"errors"
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
		//		w.ResizeN(300, 200)
		w.ShowNormal()
	})
}

type MainWindow struct {
	*tk.Window
	edit  *tk.Entry
	cflag bool
	eval  *MathEval
}

func (w *MainWindow) Equals() {
	express := w.edit.Text()
	pos := strings.Index(express, "=")
	if pos != -1 {
		express = express[:pos]
	}
	express = strings.TrimSpace(express)
	result, err := w.eval.Eval(express)
	if err != nil {
		w.edit.SetForeground("red")
	} else {
		w.edit.SetForeground("black")
		w.edit.SetText(fmt.Sprintf("%v = %v", express, result))
	}
	w.cflag = true
}

func (w *MainWindow) inputSymbol(s string) {
	w.edit.SetFocus()
	switch s {
	case "C":
		w.edit.SetText("")
		w.cflag = true
	case "=":
		w.Equals()
	default:
		if w.cflag {
			w.edit.SetText("")
			w.cflag = false
		}
		switch s {
		case "π":
			s = "Pi"
		case "平方":
			s = "pow2("
		case "立方":
			s = "pow3("
		case "平方根":
			s = "sqrt("
		case "立方根":
			s = "cbrt("
		}
		w.edit.Append(s)
	}
}

var (
	symbolsList = []string{
		"7 8 9 + - 平方",
		"4 5 6 * / 立方",
		"1 2 3 ( ) 平方根",
		"C 0 . = π 立方根",
	}
)

func NewWindow() *MainWindow {
	mw := &MainWindow{}
	mw.Window = tk.MainWindow()
	font := tk.NewUserFont("", 16).SetBold(true)
	mw.edit = tk.NewEntry(mw)
	mw.edit.SetFont(font)
	mw.edit.SetJustify(tk.JustifyRight)
	mw.edit.OnEditReturn(mw.Equals)
	mw.edit.SetFocus()

	grid := tk.NewGridLayout(mw)
	grid.AddWidgetEx(mw.edit, 0, 0, 1, 6, tk.StickyAll)
	for _, syms := range symbolsList {
		var buttons []tk.Widget
		for _, sym := range strings.Split(syms, " ") {
			btn := tk.NewButton(mw, sym)
			btn.SetWidth(5)
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

func (m *MathEval) Eval(express string) (float64, error) {
	express = strings.ToLower(express)
	express = strings.Replace(express, "pi", "3.14159265", -1)
	express = strings.Replace(express, "%", "/100", -1)
	r, err := tk.MainInterp().EvalAsString(fmt.Sprintf("expr [string map {/ *1.0/} %v]", express))
	v, err := strconv.ParseFloat(r, 0)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(fmt.Sprintf("%.15f", v), 0)
}

func NewMathEval() *MathEval {
	return &MathEval{}
}
