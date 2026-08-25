func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, val := range nums {
		countMap[val]++
	}

	frequency := make([][]int, len(nums)+1)
	for key, val := range countMap {
		frequency[val] = append(frequency[val], key)
	}

	var result []int
	for i := len(nums); i>0; i-- {
		for _, res := range frequency[i]{
			result = append(result, res)
			if len(result) == k {
				return result
			}
		}		
	}


	return result
}
