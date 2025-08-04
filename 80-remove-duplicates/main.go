package removeduplicates

/*
Yet another two pointers approach.

pointer 1 is going to be keeping track of where we're reading from in the nums arr
pointer 2, k, is going to be keeping track of where we are in the rewrite

since we're allowing at most 2 of the same number, we can check the right neighbor of our right neighbot

this would ensure only the last 2 digits of a repeated digit set would be written out
*/
func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	k := 2

	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
