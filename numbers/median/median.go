package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(median([]float64{3, 1, 2}))
	fmt.Println(median([]float64{3, 1, 4, 2}))
}

func median(nums []float64) float64 {
	// creating a copy of the slice
	vals := make([]float64, len(nums))
	copy(vals, nums)

	// sorting the values
	sort.Float64s(vals)

	// checking if the length is even or odd
	i := len(vals) / 2
	// odd number of values
	if len(vals)%2 == 1 {
		return vals[i]
	}

	return (vals[i-1] + vals[i]) / 2
}
