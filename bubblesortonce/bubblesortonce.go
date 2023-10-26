package bubblesortonce

func BubbleSortOnce(nums []int) []int {
	numsCopy := nums[:]
	length := len(numsCopy)
	for index, value := range numsCopy {
		if (index + 1 < length && (value > numsCopy[index + 1])) {
			numsCopy[index] = numsCopy[index + 1]
			numsCopy[index + 1] = value
		} else {
			break
		}
	}
	return numsCopy
}