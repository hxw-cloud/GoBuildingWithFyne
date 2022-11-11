package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var currency = "USD"

type Gold struct {
	Items  []Price `json:"items"`
	Client *http.Client
}

type Price struct {
	Currency      string    `json:"curr"`
	Price         float64   `json:"xauPrice"`
	Change        float64   `json:"chgXau"`
	PreviousClose float64   `json:"xauClose"`
	Time          time.Time `json:"-"`
}

func (g *Gold) GetPrices() (*Price, error) {
	if g.Client == nil {
		g.Client = &http.Client{}
	}

	client := g.Client

	url := fmt.Sprintf("https://data-asg.goldprice.org/dbXRates/%s", currency)
	req, _ := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("error contacting goldprice.org ", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading json ", err)
		return nil, err
	}

	gold := Gold{}
	var previous, current, change float64
	err = json.Unmarshal(body, &gold)
	if err != nil {
		log.Println("error unmarshalling", err)
		return nil, err
	}
	previous, current, change = gold.Items[0].PreviousClose, gold.Items[0].Price, gold.Items[0].Change

	var currentInfo = Price{
		Currency:      currency,
		Price:         current,
		Change:        change,
		PreviousClose: previous,
		Time:          time.Now(),
	}

	return &currentInfo, nil
}
