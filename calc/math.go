package main

import (
	"errors"
	"fmt"
	"math"
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

//对数学表达式求值
func (m *MathEval) Eval(express string) (string, error) {
	re := regexp.MustCompile("([\\d\\.]+)([a-zπ]+)")
	express = strings.ToLower(express)
	express = strings.Replace(express, " ", "", -1)
	express = re.ReplaceAllStringFunc(express, func(s string) string {
		ar := re.FindStringSubmatch(s)
		if ar[2] == "e" {
			return s
		}
		return ar[1] + "*" + ar[2]
	})
	express = strings.NewReplacer("π", "3.141592653589793", "pi", "3.141592653589793", "%", "/100", "×", "*", "x", "*", "÷", "/").Replace(express)
	r, err := tk.MainInterp().EvalAsString(fmt.Sprintf("expr [string map {/ *1.0/} %v]", express))
	v, err := strconv.ParseFloat(r, 0)
	if err != nil {
		return "无效", err
	}
	if Abs(v) > 1e-9 {
		v, _ = strconv.ParseFloat(fmt.Sprintf("%.9f", v), 0)
	}
	return fmt.Sprintf("%v", v), nil
}

/*
用牛顿法实现平方根函数。
计算机通常使用循环来计算 x 的平方根。从某个猜测的值 z 开始，
我们可以根据 z² 与 x 的近似度来调整 z，产生一个更好的猜测：
z -= (z*z - x) / (2*z)
重复调整的过程，猜测的结果会越来越精确，得到的答案也会尽可能接近实际的平方根。
*/

//绝对值
func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

//求平方根
func Sqrt(x float64) float64 {
	switch {
	case x == 0 || math.IsNaN(x) || math.IsInf(x, 1):
		return x
	case x < 0:
		return math.NaN()
	}
	z := 1.0
	for {
		d := (z*z - x) / (2 * z)
		z -= d
		if Abs(d) < 1e-9 {
			break
		}
	}
	return z
}

//求立方根
func Cbrt(x float64) float64 {
	switch {
	case x == 0 || math.IsNaN(x) || math.IsInf(x, 0):
		return x
	}
	z := 1.0
	for {
		d := (z*z*z - x) / (3 * z * z)
		z -= d
		if Abs(d) < 1e-9 {
			break
		}
	}
	return z
}

func NewMathEval() *MathEval {
	return &MathEval{}
}
