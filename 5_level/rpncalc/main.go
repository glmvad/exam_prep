package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1]

	if len(args) == 0 {
		fmt.Println("Error")
		return
	}
	strs := strings.Split(args, " ")

	oper := []string{}
	numbers := []int{}
	for _, val := range strs {
		if val >= "0" && val <= "9" {
			nb, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println("Error")
				return
			} 
			numbers = append(numbers, nb)
		} else if val == "-" || val == "+" || val == "*" || val == "/" || val == "%" {
			oper = append(oper, val)
		} else {
			fmt.Println("Error")
				return
		}
	}

	fmt.Println(numbers, oper)

	if len(numbers) - len(oper) != 1 {
		fmt.Println("Error")
		return
	}
	res := numbers[0]
	for i := 0; i < len(oper); i++ {
		switch oper[i] {
		case "+":
			res += numbers[i+1]
		case "-":
			res -= numbers[i+1]
		case "*":
			res *= numbers[i+1]
		case "/":
			res /= numbers[i+1]
		case "%":
			res %= numbers[i+1]
		}
		
	}

	fmt.Println(res)

}
