func longestConsecutive(nums []int) int {
	setOfNumbers := make(map[int]int)
	longestSequence := 0
	for i := 0; i<len(nums); i++ {
		if _, ok := setOfNumbers[nums[i]]; !ok {
			setOfNumbers[nums[i]] = 1
		}
	}
	for key, _ := range setOfNumbers {
		num := key
		sequence := 1
		if _, ok := setOfNumbers[num - 1]; !ok {
			for {
				if  _, ok := setOfNumbers[num + 1]; !ok{
					break;
				}
				sequence += 1
				num += 1
			}
		}

		longestSequence = max(longestSequence, sequence)
	}

	return longestSequence
}
