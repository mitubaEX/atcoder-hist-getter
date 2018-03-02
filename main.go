package main

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/go-achist/Models"
	"github.com/go-achist/Services"
	"net/http"
	"os"
)

func GetResultString(rateData *Models.RateData) string {
	var transition = "Highest"
	if rateData.Subset < 0 {
		transition = "Lowest"
	}
	return fmt.Sprintf("%d->%d(%d) %s\nPerformance %d\n",
		rateData.OldRate, rateData.NewRate, rateData.Subset, transition, rateData.Performance)
}

func handler(w rest.ResponseWriter, r *rest.Request) {
	name := r.PathParam("name")
	fmt.Println(name)
	result := GetResultString(Services.CrawRateData(name))
	w.WriteJson(map[string]string{"Body": result})
}

func challenge(w rest.ResponseWriter, r *rest.Request) {
	val := Models.RequestBody{}
	err := r.DecodeJsonPayload(&val)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if val.Challenge == "" {
		rest.Error(w, "Not Challenge", 400)
	}
	if val.Token == "" {
		rest.Error(w, "Not Challenge", 400)
	}
	if val.Type == "" {
		rest.Error(w, "Not Challenge", 400)
	}
	w.WriteJson(map[string]string{"challenge": val.Challenge})
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("ユーザ名を入力してちょ")
	} else {
		fmt.Println(GetResultString(Services.CrawRateData(os.Args[1])))
	}
}
