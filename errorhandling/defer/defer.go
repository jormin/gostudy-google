package main

import (
	"bufio"
	"fmt"
	"go-study/functional/fib"
	"os"
)

func tryDefer() {
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(3)
	//panic(111)
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("too many defer")
		}
	}
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathErr, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathErr.Op, pathErr.Path, pathErr.Err)
		}
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		_, _ = fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()

	writeFile("abc.txt")
}
