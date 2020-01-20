package main

import "fmt"

// 指针类型结构体
type person struct {
	name string
	city string
	age  int8
}

func main() {
	p4 := &person{}
	fmt.Printf("%T\n", p4)     //*main.person
	fmt.Printf("pe=%#v\n", p4) //pe=&main.person{name:"", city:"", age:0}
	p4.name = "三毛"
	p4.age = 19
	p4.city = "沙哈拉"
	fmt.Printf("p4=%#v\n", p4) //p4=&main.person{name:"三毛", city:"沙哈拉", age:19}

	var p5 person
	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"", city:"", age:0}

	//使用键值对初始化
	hu1 := person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	//%#v	值的Go语法表示
	fmt.Printf("hu1=%#v\n", hu1) //hu1=main.person{name:"小王子", city:"北京", age:18}

	hu2 := &person{
		name: "小王子2",
		city: "北京",
		age:  19,
	}
	fmt.Printf("hu2=%#v\n", hu2) //hu2=&main.person{name:"小王子2", city:"北京", age:19}

	//当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。
	hu3 := &person{
		city: "武汉",
	}
	fmt.Printf("hu3=%#v\n", hu3) //hu3=&main.person{name:"", city:"武汉", age:0}

	//使用值的列表初始化
	hu4 := &person{
		"hujianli",
		"孝感",
		20,
	}
	fmt.Printf("hu4=%#v\n", hu4) //hu4=&main.person{name:"hujianli", city:"孝感", age:20}

	/*
		使用这种格式初始化时，需要注意：
			1.必须初始化结构体的所有字段。
			2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
			3.该方式不能和键值初始化方式混用。
	*/

	// 结构体内存布局
	//结构体占用一块连续的内存。
	type test struct {
		a int8
		b int8
		c int8
		d int8
	}

	n := test{
		a: 1,
		b: 2,
		c: 3,
		d: 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
	/*
		n.a 0xc00000a1a0
		n.b 0xc00000a1a1
		n.c 0xc00000a1a2
		n.d 0xc00000a1a3
	*/
}
