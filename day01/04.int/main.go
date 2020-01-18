package main

import "fmt"

//整型
func main() {
	// 十进制
	var i1 = 101
	fmt.Printf("%d\n", i1)
	// 转换为二进制
	fmt.Printf("%b\n",i1)
	// 转换八进制
	fmt.Printf("%o\n",i1)
	//转换为十六进制
	fmt.Printf("%x\n",i1)

	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	// 查看变量的类型
	fmt.Printf("%T",i2)

	// 十六进制
	i3 :=0x1234567
	fmt.Printf("%d\n",i3)
}
