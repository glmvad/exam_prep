package main

import (
	"fmt"
)

func main() {
	fmt.Println(HalfSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(HalfSlice([]int{1, 2, 3}))
	fmt.Println(HalfSlice([]int{1, 2, 3, 4, 5, 6, 7}))
}

func HalfSlice(slice []int) []int {
	sl := []int{}
	if len(slice)%2 == 0 {
	for i := 0; i < len(slice)/2; i++ {
		sl = append(sl, slice[i])		
	}
	} else {
		for i := 0; i < (len(slice)/2)+1 ; i++ {
			sl = append(sl, slice[i])
		}
	}
		
	return sl
}