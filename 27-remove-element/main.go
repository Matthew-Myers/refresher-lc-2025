package removeelement

import "slices"

/*
Quick impl using the question's range of numbers of 1<=n<=50 to make sure that our numbers end up on the back.
O(n log(n)) runtime complexity because of the slices.Sort() quicksort

Could make this O(n) by replacing and swapping elements as we go with two pointers
*/
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

/* O(n) imple w/ O(1) space
Just overwrite the element when need be, don't need to worry about appending the element to the end
*/

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
