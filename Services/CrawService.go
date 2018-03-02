package Services

import (
	"github.com/go-achist/Repositorys"
	"github.com/go-achist/Models"
)

func CrawRateData () *Models.RateData{
	return Repositorys.GetCrawRateData()
}
