package main

import (
	"fmt"
	"io/ioutil"
)

func bounded(v int) int {
	if v > 100 {
		return 100
	} else if v < 0 {
		return 0
	} else {
		return v
	}
}

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}
	return result
}

func grande(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d\n", score))
	case score < 60:
		g = "F"
	case score < 70:
		g = "D"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
	default:
	}
	return g
}

func main() {
	fmt.Println(bounded(10))

	const filename = "abc.txt"
	//contents, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", contents)
	//}

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	// 条件里声明赋值，出了条件语句范围就不能使用
	//fmt.Println(contents)

	fmt.Println(eval(1, 2, "+"))
	//fmt.Println(eval(1, 2, "&"))

	fmt.Println(grande(0), grande(59), grande(60), grande(71), grande(82), grande(99), grande(100))

}
