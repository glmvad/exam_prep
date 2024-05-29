package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		os.Stdout.WriteString("0\n")
		return
	}

	nb, err := Atoi(os.Args[1])
	if !err {
		os.Stdout.WriteString("0\n")
		return
	}
	sum := 0
	for i := 2; i <= nb; i++ {
		if !IsPrime(i) {
			sum += i 
		}
	} 
	
	res := Itoa(sum)

	for _, v := range res {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')

}

func IsPrime(nb int) bool {
	if nb < 2 {
		return false 
	}

	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return true 
		}
	}
	return false
}

func Atoi(s string) (int, bool) {
	res := 0 
	power := 1
	for i := len(s)-1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			res = power*(int(s[i]-48)) + res 
			power *= 10
		} else if i == 0 && s[i] == '-' {
			return 0, false 
		} else {
			return 0, false 
		}
	}
	return res, true
}

func Itoa(nb int) string {
	res := ""
	temp := 0

	for nb > 0 {
		temp = nb%10
		res = string(temp + 48) + res
		nb /= 10 
	}
	return res
}