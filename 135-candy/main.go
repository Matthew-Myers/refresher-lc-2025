package main

func candy(ratings []int) int {
	candy := make([]int, len(ratings))
	for i := range candy {
		candy[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candy[i] = candy[i-1] + 1
		}
	}
	sum := candy[len(candy)-1]
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			rCandy := candy[i] + 1
			if candy[i-1] < rCandy {
				candy[i-1] = rCandy
			}
		}
		sum += candy[i-1]
	}
	return sum
}
