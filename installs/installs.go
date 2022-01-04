package installs

import (
	"fmt"
	"math/rand"
	"time"
)

func PrintInstall() {
	maxCount := 50
	currentCount := 0
	packageList := [...]string{"hoge", "fuga", "coffee-code", "go-package", "sample", "json-parset", "yokozuna", "kakamimochi", "ramen", "mac"}
	alertMsg := `Warning: This package is old version!!
	You should install newer version!!
	`

	for {
		randomNum := rand.Intn(10)

		fmt.Printf("go: downloading %s", packageList[randomNum])
		fmt.Println("")

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

		if randomNum == 0 {
			fmt.Printf("%s", alertMsg)
		}

		if currentCount == 50 {
			currentCount = 0
		}
	}
}
