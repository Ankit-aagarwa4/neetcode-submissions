func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var alphabets [26]int
	lenStr := len(s)

	for i := 0; i<lenStr; i++ {
		alphabets[s[i] - 'a']++
		alphabets[t[i] - 'a']--
	}

	for i :=0; i < len(alphabets); i++ {
		if alphabets[i] != 0 {
			return false
		}
	}

	return true

	// valMap := make(map[string]int)
	// for _, ch := range s {
	// 	if _, exists := valMap[string(ch)]; exists {
	// 		valMap[string(ch)]++
	// 	} else {
	// 		valMap[string(ch)] = 1;
	// 	}
	// }

	// for _, ch := range t {
	// 	if _, exists := valMap[string(ch)]; exists {
	// 		valMap[string(ch)]--
	// 		if valMap[string(ch)] < 0 {
	// 			return false
	// 		}
	// 	} else {
	// 		return false
	// 	}
	// }

	// return true
}
