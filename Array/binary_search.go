package Array

import "fmt"

func BinarySearchArray() {
	arr := []int{4, 8, 10, 15, 18, 21, 24, 27, 29, 33, 34, 37, 39, 41, 43} //This needs to be a sorted array
	key := 34
	fmt.Printf("Found %d at %dth Index\n", key, binarySearch(arr, key))
}

func binarySearch(arr []int, key int) int {
	l := 0
	h := len(arr) - 1

	for l <= h {
		mid := (l + h) / 2

		if key == arr[mid] {
			return mid
		} else if key < arr[mid] {
			h--
		} else if key > arr[mid] {
			l++
		}
	}

	return -1
}
