package main

import "fmt"

func TowerOfHanoi() {
	numberOfDisc := 5
	tohFun(numberOfDisc, 1, 2, 3)
}

func tohFun(n, a, b, c int) {
	if n > 0 {
		tohFun(n-1, a, c, b)
		fmt.Printf("From %d to %d \n", a, c)
		tohFun(n-1, b, a, c)
	}
}
