package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Data struct {
	MetaData struct {
		OneInformation     string `json:"1. Information"`
		TwoSymbol          string `json:"2. Symbol"`
		ThreeLastRefreshed string `json:"3. Last Refreshed"`
		FourInterval       string `json:"4. Interval"`
		FiveOutputSize     string `json:"5. Output Size"`
		SixTimeZone        string `json:"6. Time Zone"`
	} `json:"Meta Data"`
	TimeSeries5Min map[string]Attributes `json:"Time Series (5min)"`
}
type Attributes struct {
	OneOpen    string `json:"1. open"`
	TwoHigh    string `json:"2. high"`
	ThreeLow   string `json:"3. low"`
	FourClose  string `json:"4. close"`
	FiveVolume string `json:"5. volume"`
}

func main() {
	response, err := http.Get("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=IBM&interval=5min&apikey=BMJ0NM2DZIU3GYJG")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data Data
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		log.Println("cannot unmarshal data", err)
	}

	dataJSON, err := json.Marshal(data)

	resp, err := http.Post("http://localhost:9200/trading/_doc", "application/json", bytes.NewBuffer(dataJSON))
	if err != nil {
		fmt.Println("cannot post data to es", err)
	}
	fmt.Println("resp:", resp)
	fmt.Println("Elastic Insertion Successful")

}

