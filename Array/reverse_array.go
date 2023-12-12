package Array

import "fmt"

func ReverseArray() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Reverse string:", reverseArr(arr))
}

func reverseArr(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
	}
	return arr
}
