package main

import "fmt"

//切片

func main() {
	// 切片的定义
	var s1 []int    //定义一个存放int类型的切片
	var s2 []string // 定义一个存放string类型的切片
	fmt.Println(s1, s2)

	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"沙河", "张江", "平山村"}
	fmt.Println(s1)
	fmt.Println(s2)
	//长度和容量
	fmt.Printf("len: (s1):%d cap(s1):%d\n", len(s1), cap(s1)) //len: (s1):3 cap(s1):3
	fmt.Printf("len: (s2):%d cap(s2):%d\n", len(s2), cap(s2)) //len: (s2):3 cap(s2):3

	//由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7}
	s3 := a1[0:4]   //左包含右不包含,基于数组的切片
	fmt.Println(s3) //[1 2 3 4]
	// 切片的容量是指底层数组的容量
	fmt.Printf("len: (s3):%d cap(s3):%d\n", len(s3), cap(s3))	//len: (s3):4 cap(s3):7
	s4 := a1[1:6]
	fmt.Println(s4) //[2 3 4 5 6]
	s5 := a1[:5]
	fmt.Println(s5) //[1 2 3 4 5]
	// 切片是指向一个底层的数组，切片的长度就是元素的个数，切片的容量是底层数组第一个元素到最后一个元素的数量
	s6 := a1[3:]
	fmt.Println(s6) //[4 5 6 7]
	fmt.Printf("len: (s6):%d cap(s6):%d\n", len(s6), cap(s6))	//len: (s6):4 cap(s6):4
	s7 := a1[:]
	fmt.Println(s7) //[1 2 3 4 5 6 7]

}
