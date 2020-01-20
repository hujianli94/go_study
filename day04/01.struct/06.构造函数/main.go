package main

import "fmt"

// 构造函数
type person struct {
	name string
	city string
	age  int8
}

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func main() {
	hu9 := newPerson("胡建力", "北京", 19)
	fmt.Printf("%#v\n", hu9) //&main.person{name:"胡建力", city:"北京", age:19}
}
