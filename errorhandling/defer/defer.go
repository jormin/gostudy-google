package main

import "fmt"

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic(111)
}

func main() {
	tryDefer()
}
