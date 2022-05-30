package main

import "fmt"

func main() {
	var length int

	fmt.Scan(&length)

	start := 0
	fmt.Println(start)
	end := 1

	var sum int
	i := 0

	for i < length {
		sum = start + end
		fmt.Println(sum)
		start = end
		end = sum
		i++
	}

}
