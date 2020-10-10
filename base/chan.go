package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan []byte, 10)
	go func() {
		for {
			fmt.Printf("[go] [12] ch 内存地址： %p，长度： %d，容量： %d\n", ch, len(ch), cap(ch))
			select {
			case data := <-ch:
				fmt.Printf("[go] [15] data： %s，内存地址：%p\n", string(data), &data)
			}
		}
	}()
	data := make([]byte, 0, 32)
	data = append(data, []byte("bbbbbb")...)
	fmt.Printf("[main] [21] data： %s，内存地址：%p\n", string(data), &data)
	ch <- data
	fmt.Printf("[main] [23] 内存地址： %p，长度： %d 容量： %d\n", ch, len(ch), cap(ch))

	data = data[:0]

	data = append(data, []byte("aaa")...)
	fmt.Printf("[main] [28] data： %s，内存地址：%p\n", string(data), &data)
	ch <- data
	fmt.Printf("[main] [30] ch 内存地址： %p，长度： %d 容量： %d\n", ch, len(ch), cap(ch))

	time.Sleep(time.Second * 5)
}
