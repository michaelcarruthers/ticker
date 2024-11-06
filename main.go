package main

import (
	"log"
	"net/http"

	"github.com/michaeldcarruthers/ticker/internal/helper"
	"github.com/michaeldcarruthers/ticker/internal/stock"
)

func handler(w http.ResponseWriter, r *http.Request) {
	asset := stock.New(stock.StockConfig{
		ApiKey:   helper.EnvLookup("APIKEY"),
		Days:     helper.EnvLookup("DAYS"),
		Provider: helper.EnvLookup("PROVIDER"),
		Symbol:   helper.EnvLookup("SYMBOL"),
	})

	avg, err := asset.ClosePricesAvg()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Write(*avg)
}

func main() {
	log.Println("Starting server on :8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
