package Array

import "fmt"

func LinearSearchArray() {
	arr := []int{1, 7, 4, 9, 6, 12, 45, 2, 5}
	key := 12
	fmt.Println("Found at", linearSearch(arr, key))
	fmt.Println("Array Now:", arr)
}

func linearSearch(arr []int, key int) int {
	for i := 0; i < len(arr); i++ {
		if key == arr[i] {
			swap(&arr[i], &arr[i-1]) //Transposition
			return i - 1

			//swap(&arr[i], &arr[0]) //Move to Head/Front
			//return 0
		}
	}
	return -1
}

func swap(i, j *int) {
	temp := *i
	*i = *j
	*j = temp
}
