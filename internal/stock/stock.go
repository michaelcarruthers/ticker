package stock

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/michaeldcarruthers/ticker/internal/helper"
	"github.com/michaeldcarruthers/ticker/internal/timeseries"
)

type Stock struct {
	// Alpha Vantage API key
	apiKey string

	// Number of days to average the stock closing price
	Days string

	// Stock data provider
	Provider string

	// Stock ticker symbol
	Symbol string
}

// StockConfig represents the input parameters of a stock
type StockConfig struct {
	ApiKey   string
	Days     string
	Provider string
	Symbol   string
}

// StockResponse Provides the data structure to unmarshal the stock response
type StockResponse struct {
	TimeSeries map[string]timeseries.Index `json:"Time Series (Daily)"`
	Metadata   struct {
		Information   string `json:"1. Information"`
		LastRefreshed string `json:"3. Last Refreshed"`
		OutputSize    string `json:"4. Output Size"`
		Symbol        string `json:"2. Symbol"`
		TimeZone      string `json:"5. Time Zone"`
	} `json:"Meta Data"`
}

/*
ClosePrices returns a collection of closing prices of the stock
for the number of days specified
*/
func (s *Stock) ClosePrices() (*[]float64, error) {
	timeseries, err := s.TimeSeries()
	if err != nil {
		return nil, err
	}

	var closePrices []float64
	for _, ts := range *timeseries {
		price, err := strconv.ParseFloat(ts.Close, 64)
		if err != nil {
			return nil, err
		}
		closePrices = append(closePrices, price)
	}
	return &closePrices, nil
}

/*
ClosePricesAvg returns the average of the closing prices of the stock
for the number of days specified
*/
func (s *Stock) ClosePricesAvg() (*float64, error) {
	prices, err := s.ClosePrices()
	if err != nil {
		return nil, err
	}

	var total float64
	for _, price := range *prices {
		total += price
	}
	avg := total / float64(len(*prices))

	return &avg, nil
}

// Data requests and returns the stock data from the API
func (s *Stock) Data() (*StockResponse, error) {
	rsp, err := http.Get(s.Url())
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return nil, err
	}

	data, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err.Error())
	}

	var stockData StockResponse
	err = json.Unmarshal(data, &stockData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal stock response: %s", err.Error())
	}

	return &stockData, nil
}

// TimeSeries returns a collection of time series data associated with the stock
func (s *Stock) TimeSeries() (*[]timeseries.TimeSeries, error) {
	data, err := s.Data()
	if err != nil {
		return nil, err
	}

	dates, err := s.TimeSeriesDays()
	if err != nil {
		return nil, nil
	}

	var entries []timeseries.TimeSeries
	for _, date := range dates {
		ts := *timeseries.New(date, data.TimeSeries[date])
		entries = append(entries, ts)
	}

	return &entries, nil
}

// TimeSeriesDays a collection of dates which comprise the provided number of days
func (s *Stock) TimeSeriesDays() ([]string, error) {
	data, err := s.Data()
	if err != nil {
		return nil, err
	}

	var dates []string
	for date := range data.TimeSeries {
		dates = append(dates, date)
	}

	sortedDates, err := helper.SortByDate(dates)
	if err != nil {
		return nil, err
	}
	slices.Reverse(sortedDates)

	numDays, err := strconv.Atoi(s.Days)
	if err != nil {
		return nil, err
	}

	var entries []string
	for _, date := range sortedDates[:numDays] {
		entries = append(entries, date.Format(time.DateOnly))
	}

	return entries, nil
}

func (s Stock) ToJson() ([]byte, error) {
	data, err := s.Data()
	if err != nil {
		return nil, err
	}

	avg, err := s.ClosePricesAvg()
	if err != nil {
		return nil, err
	}

	ts, err := s.TimeSeries()
	if err != nil {
		return nil, err
	}

	content := map[string]interface{}{
		"Average":    avg,
		"MetaData":   data.Metadata,
		"TimeSeries": ts,
	}

	rsp, err := json.Marshal(&content)
	if err != nil {
		return nil, err
	}

	return rsp, err
}

// Url returns the URL of the Alpha Vantage API
func (s *Stock) Url() string {
	if s.Provider == "local" {
		return "http://localhost:9090/response.json"
	}

	base := "https://www.alphavantage.co"
	query := fmt.Sprintf(
		"function=TIME_SERIES_DAILY&symbol=%s&apikey=%s",
		s.Symbol,
		s.apiKey,
	)
	return fmt.Sprintf("%s/query?%s", base, query)
}

func New(cfg StockConfig) *Stock {
	asset := &Stock{
		apiKey:   cfg.ApiKey,
		Days:     cfg.Days,
		Provider: cfg.Provider,
		Symbol:   cfg.Symbol,
	}
	return asset
}
