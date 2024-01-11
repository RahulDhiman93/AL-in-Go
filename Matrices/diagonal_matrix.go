package Matrices

import (
	"fmt"
	"log"
)

//USING 1D ARRAY FOR 2D DIAGONAL MATRIX TO NOT DO REDUNDANT WORK FOR 0's

type matrix struct {
	arr []int
	n   int
}

func DiagonalMatrix() {
	var size int
	fmt.Print("Enter Matrix Size: ")
	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		log.Println(err)
	}
	mat := matrix{n: size}
	mat.arr = make([]int, size)
	i := 0
	for i < size {
		var x int
		fmt.Printf("Enter Value at (%d, %d): ", i, i)
		_, err := fmt.Scanf("%d", &x)
		if err != nil {
			log.Println(err)
		}
		mat.set(i, i, x)
		i++
	}
	if len(mat.arr) >= 3 {
		fmt.Println("Value at (2,2): ", mat.get(2, 2))
	}
	mat.display()
}

func (mat *matrix) set(i, j, x int) {
	if i == j {
		mat.arr[i] = x
	}
}

func (mat *matrix) get(i, j int) int {
	if i == j {
		return mat.arr[i]
	}
	return 0
}

func (mat *matrix) display() {
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
