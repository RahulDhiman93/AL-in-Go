package Recursion

import "fmt"

func TreeRecursion() {
	x := 3
	treeFun(x)
}

func treeFun(n int) {
	if n > 0 {
		fmt.Print(n, " ")
		treeFun(n - 1)
		treeFun(n - 1)
	}
}
