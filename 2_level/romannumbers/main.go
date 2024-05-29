package main

import (
	"os"
	"fmt"
	
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	nb, err := Atoi(os.Args[1])
	if !err || nb >= 4000 {
		fmt.Printf("ERROR: cannot convert to roman digit\n")
		return
	}

	arr := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	r_arr := []string{"M","CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	c_arr := []string{"M+", "(M-C)+", "D+", "(D-C)+", "C+", "(C-X)+", "L+", "(L-X)+", "X+", "(X-I)+", "V+", "(V-I)+", "I+"}

	roman := ""
	calc := ""
	
	for i := 0; i < len(arr); i++ {
		for nb > arr[i] {
			roman += r_arr[i]
			calc += c_arr[i]
			nb -= arr[i]
		}
	}


	calc = calc[:len(calc)-1]

	fmt.Printf("%v\n%v\n", calc, roman)

	
}

func Atoi(s string) (int, bool) {
	res := 0
	temp := 0
	power := 1
	neg := 1
	
	for i := len(s)-1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			temp = int(s[i]-48)
			res = temp*power + res 
			power *= 10
		} else if s[i] == '-' && i == 0 {
			neg = -1 
		} else {
			return 0 , false
		}
	}
	res = neg * res 
	return res, true
}