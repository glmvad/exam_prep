package main

import (
	"fmt"
)

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}

func DigitLen(n1, n2 int) int {
	if n2 < 2 || n2 > 36 {
		return -1 
	}

	if n1 < 0 {
		n1 = -n1
	}

	count := 0
	for n1 > 0 {
		n1 /= n2 
		count++
	}
	return count


}