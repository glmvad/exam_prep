package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2, 3}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

func ReduceInt(a []int, f func(int, int) int) {
	n1 := a[0]
	
	for i := 1; i < len(a); i++ {
		n1 = f(n1, a[i])
	}
	res := strconv.Itoa(n1)
	for _, v := range res {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}