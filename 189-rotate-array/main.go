package main

/*
simple reassign logic.

Attempted to run this with slices and append, but in Go, when you pass a slice to a function you're actually passing a slice header
Meaning that you can use append to create a new slice and reassign.

Manually updating the elements here.

The key insight is: Never reassign the slice variable inside the function if you want to modify the original array!
*/
func rotate(nums []int, k int) {
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
