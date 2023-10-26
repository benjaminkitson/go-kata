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

func TestAnotherSlice(t *testing.T) {
    input := []int{1,4,2}
    want := []int{1,2,4}
    output := BubbleSortOnce(input)
    if !utils.CheckIntSliceEquality(want, output) {
        t.Fatalf("Le fail")
    }
}

func TestOneMoreSlice(t *testing.T) {
    input := []int{3,1,8,5}
    want := []int{1,3,5,8}
    output := BubbleSortOnce(input)
    if !utils.CheckIntSliceEquality(want, output) {
        t.Fatalf("Le fail")
    }
}

func TestNoModify(t *testing.T) {
    input := []int{1,4,2}
    want := []int{1,4,2}
    BubbleSortOnce(input)
    if !utils.CheckIntSliceEquality(input, want) {
        t.Fatalf("Le fail")
    }
}



