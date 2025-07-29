package removeduplicates

/*
another 2 pointer solution

Checking to see if the element is already in the list, since it's sorted, we should be able to just step through
O(N) time complexity, O(1) space complexity
*/
func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
