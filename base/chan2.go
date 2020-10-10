package main

import (
	"fmt"
	"sync"
)

func main() {
	num := 2000

	// 结果数组
	res := [2001]int{}
	// 通道数据
	ch := make(chan int, num)

	// 往数据通道写入数据
	for i := 1; i <= num; i++ {
		ch <- i
	}
	// 写入完毕后关闭通道
	close(ch)

	// 启动8个协程
	wg := sync.WaitGroup{}
	wg.Add(8)

	for j := 0; j < 8; j++ {
		go func() {
			for {
				val, _ := <-ch
				if val != 0 {
					res[val] = sum(val)
				} else {
					// 该协程完成工作
					wg.Done()
					break
				}
			}
		}()
	}

	// 等待所有协程完成工作
	wg.Wait()

	// 循环打印
	for i, v := range res {
		fmt.Printf("res[%d] = %d\n", i, v)
	}
}

// 计算
func sum(num int) int {
	s := 0
	for i := 1; i <= num; i++ {
		s += i
	}
	return s
}
