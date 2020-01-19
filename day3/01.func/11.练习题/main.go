package main

import (
	"fmt"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Println("分配后还剩下：", left) //分配后还剩下： 30 枚金币
}

func dispatchCoin() int {

	// 初始化字典
	distribution["Matthew"] = 0
	distribution["Sarah"] = 0
	distribution["Augustus"] = 0
	distribution["Heidi"] = 0
	distribution["Emilie"] = 0
	distribution["Peter"] = 0
	distribution["Giana"] = 0
	distribution["Adriano"] = 0
	distribution["Aaron"] = 0
	distribution["Elizabeth"] = 0
	sum_gold := 0
	for i, conut := range distribution {
		for _, str := range i {
			if string(str) == "e" || string(str) == "E" {
				conut += 1
			}
			if string(str) == "i" || string(str) == "I" {
				conut += 2
			}
			if string(str) == "o" || string(str) == "O" {
				conut += 3
			}
			if string(str) == "u" || string(str) == "U" {
				conut += 4
			}
		}
		fmt.Printf("name:%s count:%d\n", i, conut)
		// 总共查找匹配到的金币数量
		sum_gold += conut

	}
	left_gold := coins - sum_gold
	return left_gold
}
