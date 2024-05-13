package main

import "fmt"

func test() {
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("1")
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
}

func main() {
	test()
}
