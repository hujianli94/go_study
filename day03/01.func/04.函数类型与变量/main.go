package main

import (
	"errors"
	"fmt"
)

// 定义函数类型
type calculation func(int, int) int

// 凡是满足这个条件的函数都是calculation类型的函数
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// 高阶函数

// 函数作为参数,传入3个参数，2个int类型的，一个func，利用func去对两个int进行操作
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 函数作为返回值,返回int类型的函数和一个error错误
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

func main() {
	var a calculation
	//add和sub都能赋值给calculation类型的变量。
	a = add
	fmt.Printf("type:%T\n", a) //type:main.calculation
	fmt.Println(a(10, 20))     //30
	a1 := sub
	fmt.Printf("type:%T\n", a1) //type:func(int, int) int
	fmt.Println(a1(30, 20))     //10

	result := calc(2, 3, add)
	fmt.Println(result) //5

}
