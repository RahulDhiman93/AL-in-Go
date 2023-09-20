package main

import "fmt"

func recursion2() {
	x := 3
	fun2(x)
}

func fun2(n int) {
	if n > 0 {
		fun2(n - 1)
		fmt.Println(n)
	}
}
