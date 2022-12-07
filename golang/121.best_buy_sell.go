package main

func maxProfit(prices []int) int {
	profit := 0
	for i, j := 0, 1; j < len(prices); j++ {
		if prices[j] < prices[i] {
			i = j
		}
		if prices[j]-prices[i] > profit {
			profit = prices[j] - prices[i]
		}
	}
	return profit
}
