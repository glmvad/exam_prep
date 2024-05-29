package main

import (
	"fmt"
)

func main() {
	fmt.Println(RectPerimeter(10, 2))
	fmt.Println(RectPerimeter(434343, 898989))
	fmt.Println(RectPerimeter(10, -2))
}

func RectPerimeter(w, h int) int {
	if h < 0 {
		return -1
	}
	return (w+h)*2
}