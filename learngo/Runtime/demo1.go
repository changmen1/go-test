// runtime.Gosched()
// 让出CPU时间片，重新等待安排任务(大概意思就是本来计划的好好的周末出去烧烤，
// 但是你妈让你去相亲,两种情况第一就是你相亲速度非常快，见面就黄不耽误你继续烧烤，
// 第二种情况就是你相亲速度特别慢，见面就是你侬我侬的，耽误了烧烤，但是还馋就是耽误了烧烤你还得去烧烤)

//Go语言中的操作系统线程和goroutine的关系：
//
//1.一个操作系统线程对应用户态多个goroutine。
//2.go程序可以同时使用多个操作系统线程。
//3.goroutine和OS线程是多对多的关系，即m:n。

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	time.Sleep(10 * time.Second)
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}
}
