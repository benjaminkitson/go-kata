package removeduplicates

func reverse(s []int) []int {
	result := []int{}
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}
	return result
}

func contains(s []int, n int) bool {
	for _, value := range s {
		if value == n  {
			return true
		}
	}
	return false
}

func RemoveDuplicates(nums []int) []int {
	numsReverse := reverse(nums)	

	result := []int{}

	for _, value := range numsReverse {
		if !contains(result, value) {
			result = append(result, value)
		}
	}

	return reverse(result)
}