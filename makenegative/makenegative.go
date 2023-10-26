package makenegative

// https://www.codewars.com/kata/55685cd7ad70877c23000102 - 8

func MakeNegative(n int) int {
	if n < 0 {
		return n
	}
	return -n
}