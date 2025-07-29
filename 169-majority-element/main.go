package majorityelement

/* sorting will give us O(nlog(n)), not linear, but O(1) space
Hashmap will give us O(n) but O(n) space

Using Boyer-moore will work O(n) and O(1)

should work out since the candidate will be len(arr) + 1 occurances, meaning that no matter the order,
it will be the candidate at the end with at least a count of 1
*/

func majorityElement(nums []int) int {
	count := 0
	var candidate int

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if candidate == num {
			count++
		} else {
			count--
		}
	}
	return candidate

}
