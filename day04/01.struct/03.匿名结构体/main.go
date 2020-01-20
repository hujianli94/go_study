package main

import "fmt"

// 匿名结构体
func main() {
	var user struct {
		name string
		age  int
	}

	user.name = "小王子"
	user.age = 18
	fmt.Printf("%#v\n", user) //struct { name string; age int }{name:"小王子", age:18}
}
