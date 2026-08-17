func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	valMap := make(map[string]int)
	for _, ch := range s {
		if _, exists := valMap[string(ch)]; exists {
			valMap[string(ch)]++
		} else {
			valMap[string(ch)] = 1;
		}
	}

	for _, ch := range t {
		if _, exists := valMap[string(ch)]; exists {
			valMap[string(ch)]--
			if valMap[string(ch)] < 0 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
