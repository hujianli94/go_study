package main

import "fmt"

//map
func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)        //还没有初始化，没有在内存中开辟空间	//true
	m1 = make(map[string]int, 10) // 要估算好该map容量，避免在程序运行时再动态扩容
	m1["理想"] = 18
	m1["hujianli"] = 19
	fmt.Println(m1)       //map[hujianli:35 理想:18]
	fmt.Println(m1["理想"]) //18
	// 约定俗成使用ok来接收，且ok是个bool值
	v, ok := m1["娜扎"]
	if !ok {
		fmt.Println("查无此key") //查无此key
	} else {
		fmt.Println(v)
	}

	// 遍历map
	for i, i2 := range m1 {
		fmt.Printf("%v %v \n", i, i2)
	}

}
