package main

import "fmt"

// 数组
// 存放元素的容器
// 必须指定存放的元素的类型和容量（长度）
func main() {
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	// 数组初始化，如果不初始化，默认都是零值，（布尔值：false，整型浮点型都是0，字符串是""）
	var b1 [3]int
	b1 = [3]int{1, 2, 3}

	b2 := [4]int{5, 6, 7}
	fmt.Println(b1)
	fmt.Println(b2)

	//根据初始值自动推断数组的长度
	hu1 := [...] string{"hu", "jian", "li"}
	fmt.Println(hu1)

	// 根据索引来初始化数组
	hu2 := [5]int{0: 1, 4: 2}
	fmt.Println(hu2) //[1 0 0 0 2]

	// 数组的遍历
	for _, i2 := range hu1 {
		fmt.Printf("%s ", i2)
	}

	citys := [...]string{"北京", "上海", "深圳", "广州"}
	// 根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Printf("index:%d, city:%s\n", i, citys[i])
	}

	// 使用for range遍历
	for i, i2 := range citys {
		fmt.Printf("index:%d, city:%s\n", i, i2)
	}

	// 多维数组
	//[[1 2] [3 4] [5 6]]
	var all [3][2]int //有3个元素，每个元素2个
	all = [3][2]int{
		[2]int{1,2},
		[2]int{3,4},
		[2]int{5,6},
	}
	fmt.Println(all)
	
	// 多维数组的遍历
	for _, i2 := range all {
		fmt.Println(i2)
		for _, i3 := range i2 {
			fmt.Println(i3)
		}
	}

	// 数组是值类型
	hh1 := [3]int{1,2,3}	// [1 2 3]
	hh2 := hh1				// [1 2 3] ctrl+C ctrl+v
	hh2[0] = 200			//[200 2 3]
	fmt.Println(hh1)		//[1 2 3]
	fmt.Println(hh2)		//[200 2 3]

}
