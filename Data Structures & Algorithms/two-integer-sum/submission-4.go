func twoSum(nums []int, target int) []int {
    twosum := make(map[int]int)
	for i, val := range nums {
		compliment := target - val
		if prvIndx, exists := twosum[compliment]; exists {
			return []int{prvIndx, i}
		}

		twosum[val] = i
	}

	return []int{}

	// slices.Sort(nums)
	// left, right := 0, len(nums) - 1
	// for left < right {
	// 	sum := nums[left] + nums[right]
	// 	if sum == target {
	// 		return []int {left, right}
	// 	} else if sum > target {
	// 		right --
	// 	} else if sum < target {
	// 		left++;
	// 	}
	// }

	// return []int{}
}
