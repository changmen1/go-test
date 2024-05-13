package main

// 互斥锁是一种常用的控制共享资源访问的方法，
//它能够保证同时只有一个goroutine可以访问共享资源。
//Go语言中使用sync包的Mutex类型来实现互斥锁。 使用互斥锁来修复上面代码的问题：
// 使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁；
//当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时，唤醒的策略是随机的。
import (
	"fmt"
	"sync"
)

var x2 int64
var wg2 sync.WaitGroup
var lock sync.Mutex

func add2(a int) {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x2 = x2 + 1
		fmt.Println(a)
		lock.Unlock() // 解锁
	}
	wg2.Done()
}
func main() {
	wg2.Add(2)
	go add2(1)
	go add2(2)
	wg2.Wait()
	fmt.Println(x2)
}
