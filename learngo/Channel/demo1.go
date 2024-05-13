package main

import "fmt"

// 创建channel
func main() {
	var ch chan int
	fmt.Println(ch) // <nil>
	ch1 := make(chan int)
	fmt.Println(ch1)
}
