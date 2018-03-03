package Services

import (
	"github.com/mitubaEX/go-achist/Repositorys"
	"github.com/mitubaEX/go-achist/Models"
)

func CrawRateData (name string) *Models.RateData{
	return Repositorys.GetCrawRateData(name)
}

func CrawContestData () []Models.ContestData{
	return Repositorys.GetCrawContestData()
}
