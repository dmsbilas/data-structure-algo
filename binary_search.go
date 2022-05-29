package main

import (
	"fmt"
	"math"
)

func binary_search(haystack []int, niddle int) int {

	left := 0
	right := len(haystack) - 1

	var result int

	for left <= right {
		mid := int(math.Ceil(float64(left+right) / 2))

		if haystack[mid] == niddle {
			result = mid
		}
		if haystack[mid] < niddle {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

func main() {
	dataArray := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}
	fmt.Println(binary_search(dataArray, 12))

}
