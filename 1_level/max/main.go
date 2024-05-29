package main

import "fmt"

func main() {
	a := []int{23, 123, 1, 11, 455, 93}
	max := Max(a)
	fmt.Println(max)
}


func Max(a []int) int {
	for i:=0; i < len(a)-1; i++ {
		for j:=0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
			}
		}
		return a[len(a)-1]
	}
