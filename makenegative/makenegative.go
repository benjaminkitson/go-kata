package makenegative

func MakeNegative(n int) int {
	if n < 0 {
		return n
	}
	return 0 - n
}