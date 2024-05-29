package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {

	str := []string{}
	hex := "0123456789abcdef"

	for _, val := range arr {
		res := ""
		for val > 0 {
			res = string(hex[int(val%16)]) + res
			val /= 16
		}
		if res != "" {
			str = append(str, res)
		}
	}

	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str[i]); j++ {
			z01.PrintRune(rune(str[i][j]))
		}
		if i != 3 && i != len(str)-1 {
			z01.PrintRune(' ')
		}
		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			z01.PrintRune('0')
			z01.PrintRune('0')
			if i != len(arr)-1 {
				z01.PrintRune(' ')
			}
		}

	}
	z01.PrintRune('\n')

	for i := 0; i < len(arr); i++ {
		if arr[i] > 31 && arr[i] < 127 {
			z01.PrintRune(rune(arr[i]))
		} else {
			z01.PrintRune('.')
		}

	}
	z01.PrintRune('\n')
}
