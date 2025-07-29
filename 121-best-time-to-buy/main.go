package besttimetobuy

/*
Keeps track of the running minimum.
If there is ever a valley then peak, we store the mox profit if it's larger
don't need to worry about keeping track of the days,
the continuesd min price tracker will catch only valid trades thanks to the count up
*/
func maxProfit(prices []int) int {
	minprice := 10000
	maxprofit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minprice {
			minprice = prices[i]
		} else if prices[i]-minprice > maxprofit {
			maxprofit = prices[i] - minprice
		}
	}
	return maxprofit
}
