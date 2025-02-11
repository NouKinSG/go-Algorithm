package one

import "math"

func maxProfit(prices []int) int {

	minprice := math.MaxInt
	maxprofit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minprice {
			// 找最小价格
			minprice = prices[i]
		} else if prices[i]-minprice > maxprofit {
			// 如果今天卖的利润比上一次的大，记录最大利润
			maxprofit = prices[i] - minprice
		}
	}
	return maxprofit
}
