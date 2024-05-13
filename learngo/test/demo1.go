package main

import (
	"container/list"
	"fmt"
)

func main() {
	var myList list.List

	myList.PushBack("go")
	myList.PushBack("grpc")
	myList.PushBack("gin")
	fmt.Println(myList)

}
