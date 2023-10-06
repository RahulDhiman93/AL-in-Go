package Recursion

import "fmt"

var memo map[int]int

func FibonacciSeries() {
	x := 999
	memo = make(map[int]int)
	fmt.Println("OUTPUT: ", fibFun(x))
}

func fibFun(n int) int {
	if n <= 1 {
		memo[n] = n
		return n
	}

	if memo[n] > 0 {
		return memo[n]
	}

	memo[n] = fibFun(n-2) + fibFun(n-1)
	return memo[n]
}
