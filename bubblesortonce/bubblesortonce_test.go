package bubblesortonce

import (
	"testing"
	"utils"
)

func TestSlice(t *testing.T) {
    input := []int{9, 7, 5, 3, 1, 2, 4, 6, 8}
    want := []int{7, 5, 3, 1, 2, 4, 6, 8, 9}
    output := BubbleSortOnce(input)
    if !utils.CheckIntSliceEquality(want, output) {
        t.Fatalf("Le fail")
    }
}

