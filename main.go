package main

import "fmt"

func main() {
	maxCount := 50
	currentCount := 0

	for currentCount < maxCount {
		fmt.Print("#")
		currentCount++
	}
	fmt.Println(" Done!!")
}
