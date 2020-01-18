package main

import "fmt"

//copy

func main() {
	//a1 := []int{1, 3, 5}
	//a2 := a1 //赋值
	//var a3 = make([]int, 3, 3)
	//copy(a3, a1) 				//copy
	//fmt.Println(a1, a2, a3)		//[1 3 5] [1 3 5] [1 3 5]
	//a1[0] = 100
	//fmt.Println(a1, a2, a3)		//[100 3 5] [100 3 5] [1 3 5]
	//
	////从切片中删除元素
	////将索引为1的3删掉
	//aa1 := append(a1[:1],a1[2:]...)
	//fmt.Println(aa1)
	//fmt.Println(cap(aa1))		//查看切片的容量为3

	x1 := [...]int{1, 2, 3} //数组
	s1 := x1                //从数组拿到切片
	// 1.切片不保存具体的值
	// 2.切片对应一个底层数组
	// 3.底层数组都是占用一块连续的内存
	fmt.Printf("value :%v len(s1): %d cap(s1):%d\n", s1, len(s1), cap(s1))
	ss1 := append(s1[:1],s1[2:]...)
	fmt.Printf("value :%v len(ss1): %d cap(ss1):%d\n", ss1, len(ss1), cap(ss1))
}
