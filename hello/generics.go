package main

import (
	"fmt"
)

var int32Slice = []int32{1, 2, 35, 6, 43, 2}

var float32Slice = []float32{23, 452, 636, 234}

func main() {
	fmt.Println(sumSlice[float32](float32Slice)) // You can omit the type inside the Square Bracket

	fmt.Println(sumSlice[int32](int32Slice))
}

func sumSlice[T int | float32 | rune](slice []T) T {
	// You can also mention the Type T as "any"
	// if you are sure the func has some generic operation
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}
