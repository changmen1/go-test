package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	var peo People = &Student{} // 创建 Student 实例时需要使用取地址符号 &
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
