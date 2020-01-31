package main

import (
	"fmt"
	"strings"
)

//闭包其实并不复杂,只要牢记闭包=函数+引用环境。

// 匿名函数

// 闭包,使用函数作为一个返回值
func adder() func(int) int {
	var x int // 声明局部变量x
	return func(y int) int {
		x += y
		return x // 返回x+传入int参数的值
	}
}

// 闭包进阶示例1,adder2函数必须传入1个参数，返回一个int的函数
func adder2(x int) func(int) int {
	return func(i int) int {
		x += i
		return x
	}
}

// 闭包进阶示例2
// 一个带1个参数的函数，返回值是一个带string返回值的函数
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 闭包进阶示例3,函数参数1个int，返回带有1个int值的2个函数
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x, y)
	}

	add(3, 5) //3 5

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y) //30
	}(10, 20)

	result1 := adder()
	result2 := result1(10)
	result3 := result1(20)
	fmt.Println(result2) //10
	fmt.Println(result3) //30

	var a = adder()
	fmt.Println(a(100)) //100
	fmt.Println(a(200)) //300

	// 闭包1进阶处理
	f2 := adder2(100)
	fmt.Println(f2(10)) //110
	fmt.Println(f2(20)) //130

	// 闭包2进阶处理
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("text")) //text.jpg
	fmt.Println(txtFunc("text")) //text.txt

	f1, f2 := calc(100)
	fmt.Println(f1(10), f2(20)) //110 90
	fmt.Println(f1(30), f2(40)) //120 80
	fmt.Println(f1(1), f2(2))   //81 79

}
