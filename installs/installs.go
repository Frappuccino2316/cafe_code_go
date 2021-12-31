package installs

import (
	"fmt"
	"math/rand"
	"time"
)

func PrintInstall() {
	maxCount := 50
	currentCount := 0

	for currentCount < maxCount {
		sleepRate := time.Duration(rand.Intn(50))
		if sleepRate == 0 {
			continue
		}

		fmt.Print("#")
		currentCount++

		time.Sleep(time.Second / sleepRate)
	}
	fmt.Println(" Done!!")
}
