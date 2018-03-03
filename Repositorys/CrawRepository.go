package Repositorys

import (
	"fmt"
	"os"
	"github.com/go-achist/DataSource/RemoteDataSource"
	"github.com/go-achist/Models"
	"github.com/go-achist/Factory"
)

func GetCrawRateData (name string) *Models.RateData {
	url := fmt.Sprintf("http://atcoder.jp/user/%s/history", name)
	rateSlise := RemoteDataSource.GetRateSlice(RemoteDataSource.GetDoc(url))
	if len(rateSlise) <= 0 {
		fmt.Println("そのユーザ名は知らん")
		os.Exit(0)
	}
	rateData := Factory.CreateRateData(rateSlise[0], rateSlise[1], rateSlise[2])
	return rateData
}

func GetCrawContestData () []Models.ContestData {
	url := "http://atcoder.jp/"
	contestSlice := RemoteDataSource.GetContestSlice(RemoteDataSource.GetDoc(url))
	var contestDataSlice = []Models.ContestData{}
	for i := 0; i < 20; i += 2 {
		cont := Models.ContestData{}
		cont.Date = contestSlice[i]
		cont.Name = contestSlice[i + 1]
		contestDataSlice = append(contestDataSlice, cont)
	}
	return contestDataSlice
}
