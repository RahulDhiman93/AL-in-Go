package main

import "fmt"

func SumOfNRecursion() {
	x := 15
	fmt.Println(sumOfNFun(x))
}

func sumOfNFun(n int) int {
	if n == 0 {
		return 0
	} else {
		return sumOfNFun(n-1) + n
	}
}
