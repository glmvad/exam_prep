package main

import "os" 

func main() {
	if len(os.Args) != 4 {
		return
	}
	a, err1 := Atoi(os.Args[1]) 
		if !err1 {
			return
		}
	oper := os.Args[2]
	b, err2 := Atoi(os.Args[3])
		if !err2 {
			return
		}
	if a > 0 && b > 0 && a+b < 0 {
		return 
	} 
	if a < 0 && b < a && a-b < 0 {
		return 
	}
	c := 0
	switch oper {
	case "+":
		c = a+b 
	case "-":
		c = a-b 
	case "*":
		c = a*b 
	case "/":
		if b == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		c = a/b 
	case "%":
		if b == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		c = a%b
	default:
		return
	}
	res := Itoa(c)
	os.Stdout.WriteString(res+"\n")
}

func Atoi(s string) (int, bool) {
	res := 0 
	temp := 0
	power := 1
	neg := 1
	for i := len(s)-1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			temp = int(s[i]-48)
			res += temp*power  
			power *= 10
		} else if s[i] == '-' && i == 0 {
			neg = -1 
		} else {
			return 0, false
		}
	}	
	return res*neg, true
}

func Itoa(n int) string{
	var res, neg string
	temp := 0
	if n < 0 {
		n = -n
		neg = "-"
	}
	if n == 0 {
		return "0"
	}
	for n >= 0 {
		temp = n%10
		res = string(temp+48) + res 
		n /= 10
	}
	return neg+res 
}