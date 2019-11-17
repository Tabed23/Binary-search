package main

import (
	"fmt"
)

func main() {
	array := make([]int, 6)
	array = append(array, 0)
	array = append(array, 1)
	array = append(array, 3)
	array = append(array, 4)
	array = append(array, 5)
	array = append(array, 6)
	if binarySearch(array, 3) {
		fmt.Println(" found the target")
	} else {
		fmt.Println("target is not in the slice")
	}

}
func binarySearch(a []int, target int) bool {

	start := 0
	end := len(a) - 1
	flag := false
	for start < end {
		mid := (start + end) / 2
		if a[mid] == target {
			flag = true
		}
		if a[mid] > target {
			end = mid + 1
		}
		if a[mid] < target {
			start = mid - 1
		}
	}
	return flag
}
