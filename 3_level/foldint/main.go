package main

import "fmt"

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

func Add(a, b int) int {
	return a+b
}

func Mul(a, b int) int {
	return a*b
}

func Sub(a, b int) int {
	return a-b
}

func FoldInt(f func(int, int) int, a []int, n int) {
	res := 0
	temp := 0

	for i:=0; i < len(a); i++ {
		temp = f(n, a[i])
		res = temp
		n = temp
		temp = 0
	}
	fmt.Println(res)
}