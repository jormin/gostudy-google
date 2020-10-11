package main

import (
	"bufio"
	"fmt"
	"go-study/functional/fib"
	"io"
)

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fib.Fibonacci()
	//for i := 0; i < 20; i++ {
	//	fmt.Printf("%d ", f())
	//}

	printFileContents(f)
}
