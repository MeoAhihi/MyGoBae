package main

import "fmt"

func firstStroke(line, size int) int {
	return 0
}
func secondStroke(line, size int) int {
	return line
}
func thirdStroke(line, size int) int {
	return size - 1
}

func printN(size int) {
	for i := 0; i < size; i++ {
		list := map[int]bool{
			firstStroke(i, size):  true,
			secondStroke(i, size): true,
			thirdStroke(i, size):  true,
		}
		fmt.Println(list)
	}
}

func main() {
	printN(7)
}
