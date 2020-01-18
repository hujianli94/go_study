package main

import "fmt"

// 字符串修改
func main() {
	s2 := "白萝卜"      //'白' '萝' '卜'
	s3 := []rune(s2) // 把字符串强制转换成一个rune切片
	s3[0] = '红'
	fmt.Println(string(s3))

	c1 := "红"
	c2 := '红'
	fmt.Printf("c1:%T c2:%T\n", c1, c2) //c1:string c2:int32

	c3 := "H"
	c4 := 'h'
	fmt.Printf("c3:%T c4:%T\n", c3, c4)
	fmt.Printf("%d\n", c4)

	// 类型转换
	n1 := 10 //int
	var f float64
	f = float64(n1) //将int强制转换为float64
	fmt.Println(f)
	fmt.Printf("%T\n", f)		//float64
}
