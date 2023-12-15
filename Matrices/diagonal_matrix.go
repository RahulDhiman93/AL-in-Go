package Matrices

import (
	"fmt"
)

//USING 1D ARRAY FOR 2D DIAGONAL MATRIX TO NOT DO REDUNDANT WORK FOR 0's

type matrix struct {
	arr []int
	n   int
}

func DiagonalMatrix() {
	size := 4
	mat := matrix{n: size}
	mat.arr = make([]int, size)
	set(&mat, 0, 0, 5)
	set(&mat, 1, 1, 7)
	set(&mat, 2, 2, 9)
	set(&mat, 3, 3, 8)
	fmt.Println("Value at (2,2): ", get(&mat, 2, 2))
	display(mat)
}

func set(mat *matrix, i, j, x int) {
	if i == j {
		mat.arr[i] = x
	}
}

func get(mat *matrix, i, j int) int {
	if i == j {
		return mat.arr[i]
	}
	return 0
}

func display(mat matrix) {
	for i := 0; i < mat.n; i++ {
		for j := 0; j < mat.n; j++ {
			if i == j {
				fmt.Print(mat.arr[i], " ")
			} else {
				fmt.Print("0", " ")
			}
		}
		fmt.Print("\n")
	}
}
