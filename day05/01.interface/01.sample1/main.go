package main

import "fmt"

// 定义一个能叫的类型
type speaker interface {
	speak() // 只要实现了speak方法的变量都是speaker类型，方法签名
}

type cat struct{}
type dog struct{}
type person struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (d dog) speak() {
	fmt.Println("汪汪汪")
}

func (p person) speak() {
	fmt.Println("哈哈哈哈")
}

func da(x speaker) {
	// 传入一个参数，自动调用speak方法，
	// da了就开始叫
	x.speak()
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person

	da(c1)
	da(d1)
	da(p1)
}
