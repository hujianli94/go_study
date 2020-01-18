package main

import "fmt"

// if条件判断
//noinspection ALL
func main() {
	age := 6
	//if age > 18 {
	//	fmt.Println("你该学习了...")
	//} else {
	//	fmt.Println("澳门首家赌场开业了........")
	//}
	if age > 35 && age < 50 {
		fmt.Println("三十而立了")
	} else if age > 18 && age < 35 {
		fmt.Println("你成年了...")
	} else if age > 50 {
		fmt.Println("步入老年了")
	}else {
		fmt.Println("你是个儿童...")
	}

	//作用域，age只在if语句中生效。结束就消失
	if age:=19;age >18{
		fmt.Println("欢迎来到澳门赌场，你成年了，可以进入")
	} else {
		fmt.Println("你还未成年")
	}
}
