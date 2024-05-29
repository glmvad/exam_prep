package main

import (
	"fmt"
)

func main() {
	fmt.Println(SwapFirst([]int{1, 2, 3, 4}))
	fmt.Println(SwapFirst([]int{3, 4, 6}))
	fmt.Println(SwapFirst([]int{1}))
}

func SwapFirst(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	for i := 0; i < len(slice); i++ {
		if i == 0 {
			slice[i], slice[i+1] = slice[i+1], slice[i]
		}
	}
	return slice
}