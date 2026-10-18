type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	result := ""
	for i:=0; i<len(strs); i++ {
		result = result + strconv.Itoa(len(strs[i])) + "#" + strs[i]
	}

	return result
}

func (s *Solution) Decode(encoded string) []string {
	str := ""
	strs := make([]string, 0)
	i := 0
	num := ""
	for i < len(encoded) {
		if encoded[i] != '#' {
			num = num + string(encoded[i])
			i++
		} else {
			a, _ := strconv.Atoi(num)
			str = encoded[i+1:i+1+a]
			strs = append(strs, str)
			i = i + 1 + a
			num = ""
		}		
	}

	return strs 
}
