package main

import "fmt"

// pointer

func main() {
	// 1.&:取地址
	// 2.*: 根据地址取值

	/*
	make和new的区别
		new很少用，一般给基本的数据类型申请内存，string，int、float、bool.....
			返回的是对应类型的指针(*string、*int)

		make是用来给slice、map、chan申请内存的，make函数返回的是对应的这三个类型本身。
	 */

	// 1.&:取地址
	n := 18
	fmt.Println(&n)

	p := &n
	fmt.Println(p)
	fmt.Printf("p type :%T", p)

	//2. *根据地质取值
	m := *p
	fmt.Println(m)
	fmt.Printf("p value: %v\n", *p)


	var a1 *int				//nill pointer
	fmt.Println(a1)			//<nil>
	var a2 = new(int)		//new函数申请一个新的内存地址
	fmt.Println(a2)			//0xc00000a100
	fmt.Println(*a2)		//0
	*a2 = 100
	fmt.Println(*a2)		//100


}
