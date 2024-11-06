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
		Days:     helper.EnvLookup("NDAYS"),
		Provider: helper.EnvLookup("PROVIDER"),
		Symbol:   helper.EnvLookup("SYMBOL"),
	})

	rsp, err := asset.ToJson()
	if rsp == nil || err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Write(rsp)
}

func main() {
	log.Println("Starting server on :8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
