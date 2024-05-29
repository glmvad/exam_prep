package main

import (
	"fmt"
)

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}

func ConcatAlternate(s1, s2 []int) []int {

		res := []int{}

		var i, j int

		if len(s1) == len(s2) {
		for i < len(s1) && j < len(s2) {
			res = append(res, s1[i])
			res = append(res, s2[j])
			i++
			j++
		}
	}
	
		if len(s2) > len(s1) {
			for i < len(s1) && j < len(s2) {
				res = append(res, s2[j])
				res = append(res, s1[i])
				i++
				j++
				for i == len(s1) && j < len(s2){
					res = append(res, s2[j])
					j++
				}
			}
		}

		if len(s1) > len(s2) {
			for i < len(s1) && j < len(s2) {
				res = append(res, s1[i])
				res = append(res, s2[j])
				i++
				j++
				for j == len(s2) && i < len(s1){
					res = append(res, s1[i])
					i++
				}
			}
		}

		if len(s2) == 0 {
			for i < len(s1) {
				res = append(res, s1[i])
				i++
			}
		} else if len(s1) == 0 {
			for j < len(s2) {
				res = append(res, s2[j])
				j++
			}
		}

		return res
}

