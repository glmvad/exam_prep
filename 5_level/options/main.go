package main

import (
	"fmt"
	"os"
)

func main() {

	str := "options: abcdefghijklmnopqrstuvwxyz"

	if len(os.Args) == 1 {
		fmt.Println(str)
		return
	}

	args := os.Args[1:]

	for _, arg := range args {
		if len(arg) <= 1 {
			fmt.Println("Invalid Option")
			return
		}
		
		for i, char := range arg {
			if i == 0 && char != '-' {
				return 
			} else if i == 0 && char == '-' && arg[i+1] == 'h' {
				fmt.Println(str)
				return 
			}
			if (char < 'a' || char > 'z') && char != '-' {
				fmt.Println("Invalid Option")
				return
			} 
			
		}
		
	}

	slice := [32]int{}
	
	for _, arg := range args {
		for i := 1; i < len(arg); i++ {
			slice[len(slice)-int(arg[i]-96)] = 1 
			
		}
	}

	for i, v := range slice {
		if i%8 == 0 && i != 0{
			fmt.Print(" ")
		}
		fmt.Print(v)


	}
	fmt.Println()
	



}