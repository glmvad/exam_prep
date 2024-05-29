package main

import (
	"fmt"
)

func main() {
	fmt.Println(SwapLast([]int{1, 2, 3, 4}))
	fmt.Println(SwapLast([]int{3, 4, 5}))
	fmt.Println(SwapLast([]int{1}))
}

func SwapLast(slice []int) []int {
	if len(slice) < 2 {
		return slice 
	}

	for i := 0; i < len(slice); i++ {
		if i == len(slice)-1 {
			slice[i-1], slice[i] = slice[i], slice[i-1]
		}
	}
	return slice
}