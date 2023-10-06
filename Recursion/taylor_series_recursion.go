package Recursion

import "fmt"

func TaylorSeriesRecursion() {
	x := 3
	n := 10
	fmt.Println(e(x, n))
}

var p float64 = 1
var f float64 = 1
var r float64

func e(x, n int) float64 {
	if n == 0 {
		return float64(1)
	} else {
		r = e(x, n-1)
		p = p * float64(x)
		f = f * float64(n)
		return r + (p / f)
	}
}
