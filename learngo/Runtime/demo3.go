// runtime.GOMAXPROCS
// Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码。默认值是机器上的CPU核心数。
// 例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上（GOMAXPROCS是m:n调度中的n）。
// Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数。
// Go1.5版本之前，默认使用的是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数。
package main

import (
	"fmt"
	"runtime"
	"time"
)

//
//import (
//	"fmt"
//	"runtime"
//	"time"
//)
//
//func a() {
//	for i := 1; i < 10; i++ {
//		fmt.Println("A:", i)
//	}
//}
//
//func b() {
//	for i := 1; i < 10; i++ {
//		fmt.Println("B:", i)
//	}
//}
//
//func main() {
//	runtime.GOMAXPROCS(1)
//	go a()
//	go b()
//	time.Sleep(time.Second)
//}

// 两个任务只有一个逻辑核心，此时是做完一个任务再做另一个任务。
//将逻辑核心数设为2，此时两个任务并行执行，代码如下。

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
