package main

import (
	"gitlab.wcxst.com/jormin/go-tools/log"
	"sync"
)

func init() {
	log.SetPrefix("goroutine")
}

func main() {
	//var a [10]int
	//for i := 0; i < 10; i++ {
	//	go func(i int) {
	//		for {
	//			a[i]++
	//		}
	//		//fmt.Printf("Hello from goroutine %d\n", i)
	//	}(i)
	//	//go func() { // race condition！
	//	//	for {
	//	//		a[i]++
	//	//	}
	//	//	//fmt.Printf("Hello from goroutine %d\n", i)
	//	//}()
	//}
	//time.Sleep(time.Millisecond)
	//fmt.Println(a)

	a := 0
	num := 1000000
	wg := sync.WaitGroup{}
	wg.Add(num)
	locker := sync.Mutex{}
	for i := 0; i < num; i++ {
		go func() {
			locker.Lock()
			a += 1
			locker.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	log.Info("a: %d", a)
}
