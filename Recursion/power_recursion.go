package Recursion

import "fmt"

func PowerRecursion() {
	m := 2
	n := 3
	fmt.Println(powFun(m, n))
}

func powFun(m, n int) int {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		return powFun(m*m, n/2)
	} else {
		return powFun(m*m, (n-1)/2) * m
	}
}
