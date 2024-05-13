package main

import "fmt"

// interface 定义一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节
// interface 是一种类型
type Cat struct{}

func (c Cat) Say() string { return "喵喵喵" }

type Dog struct{}

func (d Dog) Say() string { return "汪汪汪" }

func main() {
	c := Cat{}
	fmt.Println("猫:", c.Say())
	d := Dog{}
	fmt.Println("狗:", d.Say())
}
