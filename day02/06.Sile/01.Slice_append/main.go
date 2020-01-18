package main

import "fmt"

// 切片的追加
func main() {
	var citys []string
	// 追加1个元素
	citys = append(citys, "北京") //[北京]
	fmt.Println(citys)
	// 追加多个元素
	citys = append(citys, "上海", "深圳", "广州") //[北京 上海 深圳 广州]
	fmt.Println(citys)
	//追加一个切片
	a := []string{"成都", "重庆"} //[北京 上海 深圳 广州 成都 重庆]
	citys = append(citys, a...)
	fmt.Println(citys)
}


