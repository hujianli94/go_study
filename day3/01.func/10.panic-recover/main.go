package main

import "fmt"

func funcA() {
	fmt.Println("func A")
}
func funcB() {
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}

func main() {
	funcA()
	funcB()
	funcC()
}

/*
panic: panic in B

goroutine 1 [running]:
main.funcB(...)
	D:/go_studay/go_path/src/github.com/go_study/day3/01.func/10.panic-recover/main.go:9
main.main()
	D:/go_studay/go_path/src/github.com/go_study/day3/01.func/10.panic-recover/main.go:18 +0x9d
*/
