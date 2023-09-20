package main

import "fmt"

func recursion1() {
	x := 3
	fun(x)
}

func fun(n int) {
	if n > 0 {
		fmt.Println(n)
		fun(n - 1)
	}
}
