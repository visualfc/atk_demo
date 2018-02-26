package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/visualfc/atk/tk"
)

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
