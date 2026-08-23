func groupAnagrams(strs []string) [][]string {
	var mapGroups = make(map[[26]byte][]string)
	for _, str := range strs {
		var count [26]byte
		for i := 0; i<len(str); i++ {
			count[str[i] - 'a'] ++
		}

		mapGroups[count] = append(mapGroups[count], str)
	}

	var result = make([][]string, 0, len(mapGroups))
	for _, val := range mapGroups {
		result = append(result, val)
	}

	return result
}
