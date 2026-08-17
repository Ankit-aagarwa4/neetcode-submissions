func twoSum(nums []int, target int) []int {
	twoSum := make(map[int]int)
	for i, val := range nums {
		comple := target - val
		if prev, exist := twoSum[comple]; exist {
			return []int {prev, i}
		}

		twoSum[val] = i;
	}

	return []int{}
}
