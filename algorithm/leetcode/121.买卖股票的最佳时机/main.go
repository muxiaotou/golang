package main

import "fmt"

//所谓的动态规划法，一个判断最大收益，一个查找最小买入
func maxprofit(price []int) int {
	if len(price) < 1 {
		return 0
	}
	min, maxprofit := price[0], 0
	for i := 1; i < len(price); i++ {
		if price[i]-min > maxprofit { //此处只是判断当前的价格减去最小买入的价格是不是大于当前记录的最大收益
			maxprofit = price[i] - min
		}
		if price[i] < min { //此处只是找最小的买入价格
			min = price[i]
		}
	}
	return maxprofit
}

func main() {
	price := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxprofit(price))
}
