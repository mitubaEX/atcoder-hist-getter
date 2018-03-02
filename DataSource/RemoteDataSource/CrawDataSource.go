package RemoteDataSource

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"strconv"
)

func GetRateSlice(doc *goquery.Document) []int{
	var rateSlice = []int{}
	doc.Find("td").Each(func(i int, s *goquery.Selection) {
		if i >= 3 && i <= 5{

			// oh my god
			replaceStr := strings.Replace(strings.Replace(s.Text(), "\t", "", -1), "\n", "", -1)

			num, err := strconv.Atoi(replaceStr)
			if err != nil {
				log.Fatal(err)
			}
			rateSlice = append(rateSlice, num)
		}
	})
	return rateSlice
}

func GetDoc(url string) *goquery.Document{
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}
