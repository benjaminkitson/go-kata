package removeduplicates

import (
	"testing"
	"utils"
)

func TestValidSlice(t *testing.T) {
    input := []int{1,2,3,3,1,2}
    want := []int{3,1,2}
    output := RemoveDuplicates(input)
    if !utils.CheckIntSliceEquality(want, output) {
        t.Fatalf("Le fail")
    }
}

func TestInvalidSlice(t *testing.T) {
    input := []int{1,2,3,3,1,2,3}
    want := []int{3,1,2}
    output := RemoveDuplicates(input)
    if !utils.CheckIntSliceEquality(want, output) {
        t.Fatalf("Le fail")
    }
}

func TestEmptySlices(t *testing.T) {
    input := []int{}
    want := []int{}
    output := RemoveDuplicates(input)
    if !utils.CheckIntSliceEquality(want, output) {
        t.Fatalf("Le fail")
    }
}
