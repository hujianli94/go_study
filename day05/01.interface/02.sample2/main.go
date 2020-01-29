package main

import "fmt"

// 接口的实现
type animal interface {
	move()
	eat(string)
}

type cat struct {
	// 名字
	name string
	// 脚
	feet int8
}

type chicken struct {
	feet int8
}

func (c cat) move() {
	fmt.Println("走猫步....")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s！\n", food)
}

func (c chicken) move() {
	fmt.Println("鸡动！")
}

func (c chicken) eat(food string) {
	fmt.Printf("鸡吃%s..！\n", food)
}

func main() {
	// 定义一个接口类型变量
	var a1 animal
	bc := cat{
		name: "淘气",
		feet: 4,
	}
	a1 = bc
	fmt.Println(a1)
	a1.move()
	a1.eat("小黄鱼！")
	fmt.Printf("%T\n", a1)

	cc := chicken{feet: 2}
	a1 = cc
	a1.move()
	a1.eat("饲料！")
	fmt.Printf("%T\n", a1)

}
