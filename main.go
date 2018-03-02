package main

import (
	"fmt"
	"os"
	"github.com/go-achist/Services"
	"github.com/go-achist/Models"
)

func PrintRate(rateData *Models.RateData) {
	var transition = "Highest"
	if rateData.Subset < 0 {
		transition = "Lowest"
	}
	fmt.Printf("%d->%d(%d) %s\nPerformance %d\n",
		rateData.OldRate, rateData.NewRate, rateData.Subset, transition, rateData.Performance)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("ユーザ名を入力してちょ")
	} else {
		PrintRate(Services.CrawRateData())
	}
}
