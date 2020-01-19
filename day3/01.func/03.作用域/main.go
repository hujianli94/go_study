package main

import "fmt"

// 全局变量

// 定义全局变量
var name int64 = 10

func testGlobaVar() {
	fmt.Printf("num=%d\n", name) //函数中访问全局变量name
}

//局部变量
func testLocalVar() {
	//定义一个局部变量x，该变量x仅在函数内生效
	var x int64 = 100
	fmt.Printf("x=%d\n", x) //x=100
}

//局部变量与全局变量重名的话，函数中优先访问局部变量
func testLocalVar2() {
	//定义一个局部变量x，该变量x仅在函数内生效
	name := 100
	fmt.Printf("x=%d\n", name) //x=100
}
func main() {
	testGlobaVar() //num=10
	testLocalVar()
	//fmt.Println(x)  //报错，函数内的局部变量无法在外部使用
	testLocalVar2()   //x=100
	fmt.Println(name) // 函数外仍然使用全局的变量	10
}
