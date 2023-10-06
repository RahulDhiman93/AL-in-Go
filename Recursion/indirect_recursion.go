package Recursion

import "fmt"

func IndirectRecursion() {
	x := 20
	indirectFunA(x)
}

func indirectFunA(n int) {
	if n > 0 {
		fmt.Print(n, " ")
		indirectFunB(n - 1)
	}
}

func indirectFunB(n int) {
	if n > 1 {
		fmt.Print(n, " ")
		indirectFunA(n / 2)
	}
}
