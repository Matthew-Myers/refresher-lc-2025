package removeduplicates

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	// Use two pointers approach
	// k points to the position where we'll place the next valid element
	// i iterates through the array
	k := 2

	for i := 2; i < len(nums); i++ {
		// If current element is different from the element at position k-2,
		// it means we can include this element (allowing at most 2 occurrences)
		if nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
