package main

import "fmt"

// 指针类型结构体
type person struct {
	name string
	city string
	age  int8
}

func main() {
	var p2 = new(person)
	fmt.Printf("%T\n", p2)     //*main.person
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}

	var p3 = new(person)
	p3.name = "小王子"
	p3.age = 28
	p3.city = "上海"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"小王子", city:"上海", age:28}
}
