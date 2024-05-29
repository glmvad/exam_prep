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
// 	n := len(base)
// 	res := 0
// 	power := 1


// 	for i, v := range base {
// 		if i == 0 && v == '-' {
// 			return 0
// 		}
// 	}

// 	for i := len(s)-1; i >= 0; i-- {
// 		res += GetInt(s[i], base) * power 
// 		power *= n 
// 	}
// 	return res 
// }

// func GetInt(b byte, base string) int {
// 	for i:=0; i < len(base); i++ {
// 		if base[i] == b {
// 			return i
// 		}
// 	}
// 	return 0 





n := len(base)
res := 0 
raz := len(s)-1
for _, val := range s {
	res += Index(val, base) * Power(n, raz)
	raz--
}
return res



}

func Index(r rune, base string) int {
	for i, ch := range base {
		if r == ch {
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