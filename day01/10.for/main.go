package main

import "fmt"

//for循环
func main() {
	////基本格式
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}

	////变种1
	//var i  = 5
	//for ;i<10 ;i++  {
	//	fmt.Println(i)
	//}

	//// 变种2
	//var i = 5
	//for i < 10 {
	//	fmt.Println(i)
	//	i++
	//}

	//// 无限循环
	//for  {
	//	fmt.Println("Infinite loop....")
	//}

	//for range循环
	s := "hello beijing"
	for i, i2 := range s {
		fmt.Printf("index:%d value:%c\n", i, i2)
	}

}
