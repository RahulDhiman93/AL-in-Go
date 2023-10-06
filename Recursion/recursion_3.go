package Recursion

import "fmt"

var x = 0

func Recursion3() {
	var r int
	r = fun3(5)
	fmt.Print(r, " ")
}

func fun3(n int) int {
	if n > 0 {
		x += 1
		return fun3(n-1) + x
	}
	return 0
}
