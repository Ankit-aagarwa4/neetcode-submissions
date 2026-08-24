import (
	"slices"
)
func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, val := range nums {
		countMap[val]++
	}

	keys := make([]int, 0, len(countMap))
	for key, _ := range countMap {
		keys = append(keys, key)
	}

	slices.SortFunc(keys, func(a, b int) int {
		return countMap[b] - countMap[a]
	})


	return keys[0:k]
}
