package main

import (
	"fmt"
)

// 生成2, 3, 4, ... 到 channel 'ch'中.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// 从管道复制值 'in' 到 channel 'out',
// 移除可整除的数 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // 接收值 'in'.
		if i%prime != 0 {
			out <- i // 传入 'i' 到 'out'.
		}
	}
}
func main() {
	ch := make(chan int) // Create a newchannel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < 90000; i++ {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}
