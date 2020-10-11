package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

// 欧拉公式
func euler() {
	//fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
	//fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b = 3, 4
	fmt.Println(calTriangle(a, b))
}

func calTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	//const filename string = "abc.txt"
	//const a, b int = 3, 4
	// const 特殊在于，可以指代任意类型，这里不指定int，调用 math.Sqrt 就不需要特意转 float64
	//const a, b = 3, 4
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)

	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	//const (
	//	cpp = 0
	//	java = 1
	//	python = 2
	//	golang = 3
	//)
	// 普通枚举类型
	const (
		cpp = iota
		java
		_
		python
		golang
		javascript
	)
	fmt.Println(cpp, java, python, golang, javascript)

	// 自增枚举类型
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)

}

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
	euler()
	triangle()
	consts()
	enums()
}
