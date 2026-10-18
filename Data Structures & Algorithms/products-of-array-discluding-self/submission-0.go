func productExceptSelf(nums []int) []int {
	productToLeft := getProductToLeft(nums)
	productToRight := getProductToRight(nums)
	result := make([]int, 0)
	for i:=0; i<len(nums); i++ {
		result = append(result, productToLeft[i] * productToRight[i])
	}

	return result
}

func getProductToLeft(num []int) []int {
	result := make([]int, 0)
	for i:=0; i<len(num); i++ {
		if i == 0{
			result = append(result, 1)
		} else {
			result = append(result, result[i-1]*num[i-1])
		}
	}

	return result;
}

func getProductToRight(num []int) []int {
	result := make([]int, len(num))
	for i:=len(num) - 1; i>=0; i-- {
		if i == len(num) - 1{
			result[i] = 1
		} else {
			result[i] = result[i+1]*num[i+1]
		}
	}

	return result;
}
