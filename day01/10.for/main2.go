package main

import "fmt"

//switch语句,简化大量的判断

func main() {
	//var n = 5
	//switch n {
	//case 1:
	//	fmt.Println("111")
	//case 2:
	//	fmt.Println("2222")
	//case 3:
	//	fmt.Println("333333")
	//case 4:
	//	fmt.Println("444444")
	//default:
	//	fmt.Println("not found...")
	//
	//}
	//// 变种1
	//switch n := 5; n {
	//case 1:
	//	fmt.Println("111")
	//case 2:
	//	fmt.Println("2222")
	//case 3:
	//	fmt.Println("333333")
	//case 4:
	//	fmt.Println("444444")
	//default:
	//	fmt.Println("not found...")
	//
	//}

	//// 变种2
	//switch n := 5; n {
	//case 1, 3, 5, 7, 9:
	//	fmt.Println("111")
	//case 2, 4, 6, 8, 10:
	//	fmt.Println("2222")
	//default:
	//	fmt.Println("not found...")
	//
	//}

	////// 变种3
	//age := 17
	//switch {
	//case age > 18 && age < 30:
	//	fmt.Println("aaaaaaaaaaaaa")
	//case age < 40 && age > 30:
	//	fmt.Println("bbbbbbbbbb")
	//default:
	//	fmt.Println("not found...")
	//
	//}

	////// 变种3,兼容C语言
	age := 19
	switch {
	case age > 18 && age < 30:
		fmt.Println("aaaaaaaaaaaaa")
		fallthrough
	case age < 40 && age > 30:
		fmt.Println("bbbbbbbbbb")
	default:
		fmt.Println("not found...")

	}


}
