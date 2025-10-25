package main

import "fmt"

func main() {
	arr := [5]int{1, 3, 4, 5, 7}

	left, right := 0, len(arr)-1

	target := 7
	found := false

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			fmt.Println("target found at index :", mid)
			found = true
			break
		} else if target > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if !found {
		fmt.Println("target not found")
	}

}
