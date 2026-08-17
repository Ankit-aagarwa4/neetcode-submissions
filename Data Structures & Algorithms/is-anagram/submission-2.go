func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var anag [26]int
	for i:=0; i<len(s); i++ {
		anag[s[i] - 'a']++
		anag[t[i] - 'a']--
	}

	for i:=0; i < 26; i++ {
		if anag[i] != 0 {
			return false
		}
	}

	return true
 }
