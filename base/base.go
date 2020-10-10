package main

import (
	"fmt"
	"math"
)

func main() {
	a := 1
	b := a
	a++
	fmt.Println(a, b)

	var t complex128
	t = 2.1 + 3.14i
	fmt.Println(t)
	fmt.Println(real(t))
	fmt.Println(imag(t))

	nan := math.NaN()
	fmt.Println(nan, nan == nan, nan < nan, nan > nan)

	b1 := 'a'
	b2 := 'b'
	fmt.Println(b1, b2)
	fmt.Printf("%c\n", b1)
	fmt.Printf("%c\n", b2)

	b3 := '周'
	fmt.Println(b3)
	fmt.Printf("%c\n", b3)
	b4 := '北'
	fmt.Println(b4)
	fmt.Printf("%T\n", b4)
	fmt.Printf("b4 = %c\n", b4)

	//var c3 byte = '北'
	//fmt.Printf("%T\n", c3)
	//fmt.Printf("c3=%c", c3)

	//panic("111111")

	//fmt.Println("222222")
}
