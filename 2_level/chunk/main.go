package main

import "fmt"

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {
	if len(slice) == 0 {
		fmt.Println(slice)
		return
	}
	if size == 0 {
		fmt.Println()
		return
	}
	res := [][]int{}
	temp := []int{}

	for len(slice) > size {
		temp = slice[:size]
		res = append(res, temp)
		slice = slice[size:]
	}
	if len(slice) < size || len(slice) == size{ 
		res = append(res, slice)
	}
	fmt.Println(res)
}
