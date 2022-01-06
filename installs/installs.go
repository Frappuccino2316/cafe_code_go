package installs

import (
	"fmt"
	"math/rand"
	"time"
)

func PrintInstall() {
	// 時刻からランダムなシード値を設定
	rand.Seed(time.Now().UnixNano())

	// #を50文字出力する
	maxCount := 50
	// #をカウントする
	currentCount := 0
	// 出力用のダミーパッケージ名
	packageList := [...]string{
		"hoge",
		"fuga",
		"coffee-code",
		"go-package",
		"sample",
		"json-parset",
		"yokozuna",
		"kakamimochi",
		"ramen",
		"mac",
	}
	// ダミーのアラートメッセージ
	alertMsg := `Warning: This package is old version!!
	You should install newer version!!`

	// パッケージのダウンロードを繰り返す
	for {
		// 0~9でランダムに生成
		randomNum := rand.Intn(10)

		// パッケージ名をランダムに出力
		fmt.Printf("go: downloading %s", packageList[randomNum])
		fmt.Println("")

		// 出力している#が50未満なら#を出力する
		for currentCount < maxCount {
			// ランダムに時間をかけているように見せるための数値を生成
			sleepRate := time.Duration(rand.Intn(50))

			fmt.Print("#")
			currentCount++

			// sleepRateが0ならsleepしない
			if sleepRate == 0 {
				continue
			}
			time.Sleep(time.Second / sleepRate)
		}

		fmt.Println(" Done!!")

		// randomNumが0ならダミーのアラートを出す
		if randomNum == 0 || randomNum == 1 || randomNum == 2 {
			fmt.Printf("%s", alertMsg)
			fmt.Println("")
		}

		// 最後に#のカウントを0に戻す
		currentCount = 0
	}
}
