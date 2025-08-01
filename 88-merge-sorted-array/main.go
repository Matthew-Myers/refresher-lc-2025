package main

/*
 * Work backwards in the new array.
 * There will be enough space so we don't need to worry about overwriting
 *
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	// Start from the end of both arrays
	i := m - 1
	j := n - 1
	k := m + n - 1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}

}
