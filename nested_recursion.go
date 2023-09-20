package main

import "fmt"

func NestedRecursion() {
	x := 95
	fmt.Println(nestedFun(x))
}

func nestedFun(n int) int {
	if n > 100 {
		return n - 10
	} else {
		return nestedFun(nestedFun(n + 11))
	}
}
