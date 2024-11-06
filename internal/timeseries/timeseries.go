package timeseries

type TimeSeries struct {
	// Date of the time series data
	Date string

	// Raw time series data from the stock provider
	data Index

	// Closing price of the time series data
	Close string
}

// Index represents the raw time series data from the stock provider
type Index struct {
	Close  string `json:"4. close"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Open   string `json:"1. open"`
	Volume string `json:"5. volume"`
}

// SetClose sets the closing price of the time series data
func (ts *TimeSeries) SetClose(closePrice string) {
	ts.Close = closePrice
}

func New(date string, data Index) *TimeSeries {
	ts := &TimeSeries{
		data: data,
		Date: date,
	}
	ts.SetClose(data.Close)
	return ts
}
