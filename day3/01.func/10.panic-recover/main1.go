package main

import "fmt"

func funcE() {
	fmt.Println("func E")
}

//程序运行期间funcF中引发了panic导致程序崩溃，异常退出了。这个时候我们就可以通过recover将程序恢复回来，继续往后执行。
func funcF() {
	defer func() {
		err := recover()
		//如果程序出现了panic错误，可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in F")
}

func funcG() {
	fmt.Println("func G")
}

func main() {
	funcE() //func E
	funcF() //recover in B
	funcG() //func G
}

/*
注意：
	1.recover()必须搭配defer使用。
	2.defer一定要在可能引发panic的语句之前定义。
*/
