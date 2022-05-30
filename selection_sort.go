package main

import "fmt"

func selection_sort(arr []int) {
	cursor_position := 0

	for j := 0; j < len(arr)-1; j++ {

		for i := cursor_position; i < len(arr)-1; i++ {
			if arr[i+1] <= arr[cursor_position] {
				temp := arr[cursor_position]
				arr[cursor_position] = arr[i+1]
				arr[i+1] = temp
				cursor_position += 1
			}
			fmt.Println(arr)
		}
		cursor_position = j
	}

}

func main() {
	dataSet := []int{10, 5, 2, 8, 7, 7, 9, 67}
	selection_sort(dataSet)
	fmt.Println("Final result : ", dataSet)

}
