package main

import "fmt"

func main() {
	dataSet := []int{4, 6, 7, 8, 9, 12, 13, 14}

	var userInput int

	fmt.Scan(&userInput)

	answer := linear_search(dataSet, userInput)
	fmt.Println(answer)
}

func linear_search(heyStack []int, niddle int) int {
	var result int = -1
	for i := 0; i < len(heyStack); i++ {
		if heyStack[i] == niddle {
			result = i
		}
	}
	return result
}
