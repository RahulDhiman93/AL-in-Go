package main

import "fmt"

func FactorialRecursion() {
	x := 5
	fmt.Println(facFun(x))
}

func facFun(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	} else {
		return facFun(n-1) * n
	}
}
