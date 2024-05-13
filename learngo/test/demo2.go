package main

import (
	"fmt"
)

func add(a, b int) (sum int, err error) {
	sum = a + b
	return
}

func main() {
	// 在内存里面创建了一个地址  地址里面存了一个值就是1
	a := 1
	b := 2
	//fmt.Println(add(a, b)) // 只是把值传过去了，没有传内存地址  想改变一个值必须拿到他的内存地址才可以改变

	c, err := add(a, b)
	if err == nil {
		fmt.Println(c)
	}
}
