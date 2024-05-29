package main

import (
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	nb1, _ := strconv.Atoi(os.Args[1])
	res := ""
	for i := '1'; i <= '9'; i++ {
		nb2, _ := strconv.Atoi(string(i))
		res = string(i) + " " + " x" + " " + os.Args[1] + " " + "=" + " " + Itoa(nb1*nb2)
		os.Stdout.WriteString(res+"\n")
	}
}

func Itoa(nb int) string {

	neg := ""
	res := ""
	temp := 0
	for nb < 0 {
		neg = "-"
		nb = -nb
	}
	for nb > 0 {
		temp = nb % 10
		res = string(temp+'0') + res
		nb /= 10
	}
	return neg + res
}