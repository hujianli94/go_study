package main

import "fmt"

//定义类型
type Newint int

//定义别名
type MyInt = int

func main() {
	var a Newint
	var b MyInt

	fmt.Printf("type of a:%T\n", a) //type of a:main.Newint
	fmt.Printf("type of b:%T\n", b) //type of b:int
}
