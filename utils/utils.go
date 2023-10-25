package utils

func CheckIntSliceEquality(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, n := range a {
		if b[i] != n {
			return false
		}
	}
	return true
}