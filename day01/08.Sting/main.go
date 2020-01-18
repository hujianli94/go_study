package main

import (
	"fmt"
	"strings"
)

// 字符串

func main() {
	// 字符串
	s := "hello beijing"
	// 字符，单独的字母、汉字、符号表示一个字符
	c1 := 'h'
	c2 := 'j'
	c3 := 'j'
	c4 := '胡'
	fmt.Println(s)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)

	path := "D:\\21天学编程语言系列\\后台开发技术"
	fmt.Println(path)

	s2 := "I' am ok"
	fmt.Println(s2)
	s3 := `
		世情薄
		人易恶
		雨送黄昏花易落
		`
	fmt.Println(s3)

	// 反引号原样输出
	path2 := `D:\21天学编程语言系列\后台开发技术`
	fmt.Println(path2)

	// 字符串相关操作
	fmt.Println(len(s3))

	// 字符串拼接
	name := "理想"
	world := "大帅比"
	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s %s", name, world)
	fmt.Println(ss1)

	// 分割
	ret :=strings.Split(ss1,"\\")
	fmt.Println(ret)
	// 包含
	fmt.Println(strings.Contains(ss1,"理想"))
	fmt.Println(strings.Contains(ss1,"理想1"))

	// 前缀
	fmt.Println(strings.HasPrefix("理想",ss1))
	// 后缀
	//......

}
