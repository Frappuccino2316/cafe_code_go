package installs

import "fmt"

func PrintInstall() {
	maxCount := 50
	currentCount := 0

	for currentCount < maxCount {
		fmt.Print("#")
		currentCount++
	}
	fmt.Println(" Done!!")
}
