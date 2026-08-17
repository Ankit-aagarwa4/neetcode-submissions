func hasDuplicate(nums []int) bool {
    valMap := make(map[int]int)
	for _, val := range nums {
		if _, exists := valMap[val]; exists {
			return true
		} else {
			valMap[val] = 1
		}
	}

	return false
}
