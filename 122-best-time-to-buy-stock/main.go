package main

/*
  Since we can buy and sell each and every day, it doesn't really matter if we find when we can buy low and sell hi
    All that matters is that any time stock goes up, we profit
*/

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}
