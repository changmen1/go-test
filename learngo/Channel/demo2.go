package main

import "fmt"

// channel操作
// 通道有发送（send）、接收(receive）和关闭（close）三种操作。
//发送和接收都使用<-符号。

func main() {
	// 定义 有缓冲通道 1表示通道允许存储的数量 为0的话就是无缓冲通道只能使用goroutine去接受
	ch1 := make(chan int, 1)
	//发送一个值到通道
	ch1 <- 10
	//接收 从一个通道中接收一个值
	x := <-ch1
	fmt.Println(x)
	fmt.Println(<-ch1)

	// 可以通过内置的close()函数关闭channel（如果你的管道不往里存值或者取值的时候一定记得关闭管道）
	close(ch1)
}
