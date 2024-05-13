//单向通道
//有的时候我们会将通道作为参数在多个任务函数间传递，
//很多时候我们在不同的任务函数中使用通道都会对其进行限制，
//比如限制通道在函数中只能发送或只能接收。
//1.chan<- int是一个只能发送的通道，可以发送但是不能接收；
//2.<-chan int是一个只能接收的通道，可以接收但是不能发送。
//在函数传参及任何赋值操作中将双向通道转换为单向通道是可以的，但反过来是不可以的。

package main

import "fmt"

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
