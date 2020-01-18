package main

import "fmt"

func main() {
	// make函数创建切片
	s1 := make([]int, 5, 10)		//长度是5,容量是10
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))	//s1=[0 0 0 0 0] len(s1)=5 cap(s1)=10

	s2 := make([]int, 0, 10)		//长度是0,容量是10
	fmt.Printf("s1=%v len(s2)=%d cap(s2)=%d\n", s2, len(s2), cap(s2))	//s1=[0 0 0 0 0] len(s1)=5 cap(s1)=10

	//切片的赋值
	ss3 :=[]int{1,2,3}
	ss4 :=ss3				//s3和s4都指向同一个底层数组
	fmt.Println(ss3,ss4)	//[1 2 3] [1 2 3]
	ss3[0] = 1000
	fmt.Println(ss3,ss4)	//[1000 2 3] [1000 2 3]

	// 切片的遍历
	for i:=0;i<len(ss3) ;i++  {
		fmt.Println(ss3[i])
	}

	for _, i2 := range ss3 {
		fmt.Println(i2)
	}
}
