package removeelement

import "slices"

func removeElement(nums []int, val int) int {
	placeholder := 55
	replaceCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = placeholder
			replaceCount++
		}
	}
	slices.Sort(nums)
	return len(nums) - replaceCount
}
