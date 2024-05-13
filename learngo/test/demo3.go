package main

import "fmt"

func callback(a int, f func(int, int)) {
	fmt.Println("第二步进入了callback方法")
	fmt.Println("第三步调用了 f方法 = 在调用 func1这个方法")
	f(a, 2)
}
func main() {
	func1 := func(a, b int) {
		fmt.Println("第四步调用了 func1被调用 func1(1,2)")
		fmt.Printf("sum: %d\r\n", a+b)
	}
	fmt.Println("第一步调用callback方法")
	callback(1, func1)
}
