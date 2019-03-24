package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 3
var ss = "kkk"

var (
	ac = 1
	bc = 2
	cc = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, s, b)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "efg"
	b = 5
	var f int
	fmt.Println(a, b, c, s, f)
}

func euler() {
	c := 3 + 4i
	fmt.Println(
		cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(
		cmplx.Exp(1i*math.Pi) + 1)
	fmt.Println(cmplx.Abs(c))
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt((float64(a*a + b*b))))
	fmt.Println(c)
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, ac, bc, cc)
	euler()
	triangle()
	consts()
}
