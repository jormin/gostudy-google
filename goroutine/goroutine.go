package main

import (
	"fmt"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
			}
			//fmt.Printf("Hello from goroutine %d\n", i)
		}(i)
		//go func() { // race condition！
		//	for {
		//		a[i]++
		//	}
		//	//fmt.Printf("Hello from goroutine %d\n", i)
		//}()
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
