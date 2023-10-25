package main

import (
	"fmt"
	"removeduplicates"
)

func main() {
	sliceBoi := []int{1,2,1,3,1}
	result := removeduplicates.RemoveDuplicates(sliceBoi)
	fmt.Println(result)
}