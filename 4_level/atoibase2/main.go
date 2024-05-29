package main

import (
	"fmt"
)

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

func AtoiBase(s string, base string) int {
	n := len(base)
	if n < 2 {
		return 0
	}
	for i, v := range base {
		if i == 0 && v == '-' || i == 0 && v == '+' {
			return 0
		}
	}
	res := 0
	raz := len(s) - 1
	for _, v := range s {
		res += Index(v, base) * Power(n, raz)
		raz--
	}
	return res
}

func Index(v rune, base string) int {
	for i, ch := range base {
		if v == ch {
			return i
		}
	}
	return -1
}

func Power(n, raz int) int {
	res := 1
	for raz > 0 {
		res *= n
		raz--
	}
	return res
}
