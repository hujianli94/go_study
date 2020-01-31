package main

import "fmt"

/*
定义一个求两数之和的函数
*/

func intSum(x int, y int) int {
	return x + y
}

// 函数的参数中如果相邻变量的类型相同，则可以省略类型
func addSum(a, b int) int {
	return a + b
}

//打印函数
func sayhello() {
	fmt.Println("hello 沙河")
}

// 函数
func main() {
	sum := intSum(5, 6)
	fmt.Println(sum) //11
	sayhello()       //hello 沙河
}
