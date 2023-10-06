package Recursion

import "fmt"

func Recursion1() {
	x := 3
	fun(x)
}

func fun(n int) {
	if n > 0 {
		fmt.Print(n, " ")
		fun(n - 1)
	}
}
