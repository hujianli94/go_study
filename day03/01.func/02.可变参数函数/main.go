package main

import "fmt"

// 可变参数是指函数的参数是不固定的。Go中的可变参数通过在参数名后加....来标识

func intSum2(x ...int) int {
	fmt.Println(x) // x是一个切片
	sum := 0
	for _, i2 := range x {
		sum = sum + i2
	}
	return sum
}

//固定参数搭配可变参数使用时，可变参数要放在固定参数的后面
func intSum3(x int, y ...int) int {
	fmt.Println(x, y) // x是一个切片
	sum := 0
	for _, i2 := range y {
		sum = sum + i2
	}
	return sum
}

// 多返回值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub

}

// 返回值命名
//函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回。
func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	ret1 := intSum2()
	ret2 := intSum2(10)
	ret3 := intSum2(10, 20)
	ret4 := intSum2(10, 20, 30)
	fmt.Println(ret1, ret2, ret3, ret4)

	//本质上，函数的可变参数是通过切片来实现的。
	ret5 := intSum3(100)
	ret6 := intSum3(100, 10)
	ret7 := intSum3(100, 10, 20)
	ret8 := intSum3(100, 10, 20, 30)
	fmt.Println(ret5, ret6, ret7, ret8)

	resultsum, resultsub := calc(10, 5)
	fmt.Println(resultsum, resultsub) //15 5

	ressum, ressub := calc2(20, 10)
	fmt.Println(ressum, ressub) //30 10
}

/*
[]
[10]
[10 20]
[10 20 30]
0 10 30 60

100 []
100 [10]
100 [10 20]
100 [10 20 30]
0 10 30 60

*/
