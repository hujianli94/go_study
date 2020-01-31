package main

import "fmt"

// 指针取值

func main() {
	a := 10
	b := &a                          // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)  //type of b:*int
	c := *b                          // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)  //type of c:int
	fmt.Printf("value of c:%v\n", c) //value of c:10
}
