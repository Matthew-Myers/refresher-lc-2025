package main

/*
simple reassign logic.

Attempted to run this with slices and append, but in Go, when you pass a slice to a function you're actually passing a slice header
Meaning that you can use append to create a new slice and reassign.

Manually updating the elements here.

The key insight is: Never reassign the slice variable inside the function if you want to modify the original array!
*/
func rotateBrute(nums []int, k int) {
	if len(nums) == 0 {
		return
	}

	k = k % len(nums)

	for i := 0; i < k; i++ {
		last := nums[len(nums)-1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = last
	}
}

func rotateRev(nums []int, k int) {
	if len(nums) == 0 {
		return
	}

	k = k % len(nums)

	// Step 1: Reverse the entire array
	reverse(nums, 0, len(nums)-1)

	// Step 2: Reverse the first k elements
	reverse(nums, 0, k-1)

	// Step 3: Reverse the remaining elements
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
