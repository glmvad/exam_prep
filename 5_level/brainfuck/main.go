package main

import "os"

const SIZE int = 2048

func main() {
	code := os.Args[1]

	if len(code) >= 4096 {
		return
	}

	array := make([]int, SIZE)

	idx := 0
	ptr := 0 

	for idx < len(code) {
		switch string(code[idx]) {
		case ">":
			ptr = (ptr + SIZE + 1)%SIZE 
		case "<":
			ptr = (ptr + SIZE - 1)%SIZE 
		case "+":
			array[ptr] = (array[ptr] + 255 + 1)%255
		case "-":
			array[ptr] = (array[ptr] + 255 - 1)%255
		case ".":
			os.Stdout.WriteString(string(array[ptr]))
		case "[":
			if array[ptr] == 0 {
				for string(code[idx]) != "]" {
					idx++
				}
			}
		case "]":
			if array[ptr] != 0 {
				for string(code[idx]) != "[" {
					idx--
				}
			}
		}
		idx++
	}

}




