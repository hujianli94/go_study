package main

import "fmt"

//var name1 string
//var age2 int
//var isOk3 bool

// 批量声明
var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "胡建力"
	age = 18
	isOk = true
	fmt.Printf("name:%s age:%d isOk:%d\n", name, age, isOk)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isOk)

	// 声明变量同时赋值
	var s1 string = "hujianli"
	fmt.Println(s1)

	// 类型推导（根据值判断是什么类型）
	var s2  = "20"
	fmt.Println(s2)

	// 简短变量声明,只能在函数里面使用
	s3 := "呵呵呵呵"
	fmt.Println(s3)
	// 同一个作用域中不能重复声明相同名的变量
	//s1 := 100
	// 匿名变量是一个特殊的变量

}
