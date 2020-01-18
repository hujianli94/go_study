package main

import "fmt"

func main() {
	// 求数组所有元素的和
	a1 := [...]int{1, 2, 3, 4, 5, 6}
	sum := 0
	for _, i2 := range a1 {
		sum += i2
	}
	fmt.Println(sum) //21

	// 找出数组中和为指定值的两个元素之和等于8的下标
	for i := 0; i < len(a1); i++ {
		for j := i + 1; j < len(a1); j++ {
			if a1[i]+a1[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}
