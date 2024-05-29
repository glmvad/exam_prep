package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(NotDecimal("0.1"))
	fmt.Println(NotDecimal("174.2"))
	fmt.Println(NotDecimal("0.1255"))
	fmt.Println(NotDecimal("1.20525856"))
	fmt.Println(NotDecimal("-0.0f00d00"))
	fmt.Println(NotDecimal(""))
	fmt.Println(NotDecimal("-19.525856"))
	fmt.Println(NotDecimal("1952"))
}

func NotDecimal(dec string) string  {
	nb := 0 
	temp := 0
	power := 1
	if dec == "" {
		return dec
	}
	for i := len(dec)-1; i >= 0; i-- {
		if dec[i] >= '0' && dec[i] <= '9' {
			temp = int(dec[i] - 48)
			nb = temp*power + nb 
			power *= 10
		
		} else if dec[i] == '-' {
			nb = -nb
		}
	}
	if nb == 0 {
		return dec
	}

	res := strconv.Itoa(nb)
	return res
}