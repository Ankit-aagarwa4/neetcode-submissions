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
}
