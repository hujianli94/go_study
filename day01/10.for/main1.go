package main

import "fmt"

// 流程控制之跳出for循环
func main() {
	//for i:=0; i<10 ;i++  {
	//	if i == 5{
	//		break
	//	}
	//	fmt.Printf("%d ",i)
	//}

	// 流程控制之跳过for循环
	for i:=0; i<10 ;i++  {
		if i == 5{
			continue
		}
		fmt.Printf("%d ",i)
	}
}
