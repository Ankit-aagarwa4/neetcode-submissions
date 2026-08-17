import(
	"slices"
)
func groupAnagrams(strs []string) [][]string {
	anagram := make(map[string][]string)
	var final [][]string
	for _, val := range strs {
		runes := []rune(val)
		slices.Sort(runes)
		str := string(runes)
		if _, exists := anagram[str]; !exists {
			anagram[str] = []string{val}
		} else {
			anagram[str] = append(anagram[str], val)
		}
	}

	for _, val := range anagram {
		final = append(final, val)
	}

	return final
}
