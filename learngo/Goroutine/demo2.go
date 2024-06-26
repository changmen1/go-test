package main

// 启动多个goroutine

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello2(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello2(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}
