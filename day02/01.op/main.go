package main

import "fmt"

// 运算符
func main() {
	var (
		a = 5
		b = 10
	)
	fmt.Println(a + b) //15
	fmt.Println(a - b) //-5
	fmt.Println(a * b) //50
	fmt.Println(a / b) //0
	fmt.Println(a % b) //5

	// 关系运算符
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a > b)
	fmt.Println(a < b)

	//逻辑运算符
	// 如果年龄大于18岁，并且年龄小于60
	age := 22
	if age > 18 && age < 60 {
		fmt.Println("上班吧")
	} else if age > 60 || age < 18 {
		fmt.Println("退休吧")
	}

}
