package main

import "fmt"

//指针传值示例
/*
变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：

	·对变量进行取地址（&）操作，可以获得这个变量的指针变量。
	·指针变量的值是指针地址。
	·对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
*/

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func main() {
	a := 10
	modify1(a)
	fmt.Println(a) //10
	modify2(&a)
	fmt.Println(a) //100
}
