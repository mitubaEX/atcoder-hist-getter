package Factory

import "github.com/mitubaEX/go-achist/Models"

func CreateRateData(performance int, newRate int, subset int) *Models.RateData {
	newRateData := new(Models.RateData)
	newRateData.Performance = performance
	newRateData.NewRate = newRate
	newRateData.OldRate = newRate - subset
	newRateData.Subset = subset
	return newRateData
}
