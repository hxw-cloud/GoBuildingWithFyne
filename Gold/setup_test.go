package main

import (
	"Gold/repository"
	"bytes"
	"fyne.io/fyne/v2/test"
	"io"
	"net/http"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {
	a := test.NewApp()
	testApp.App = a
	testApp.MainWindow = a.NewWindow("")
	testApp.HTTPClient = client
	testApp.DB = repository.NewTestRepository()
	os.Exit(m.Run())
}

var jsonToReturn = `{
    "ts": 1667471451830,
    "tsj": 1667471445504,
    "date": "Nov 3rd 2022, 06:30:45 am NY",
    "items": [
        {
            "curr": "USD",
            "xauPrice": 1620.545,
            "xagPrice": 18.9765,
            "chgXau": -13.645,
            "chgXag": -0.186,
            "pcXau": -0.835,
            "pcXag": -0.9706,
            "xauClose": 1634.19,
            "xagClose": 19.1625
        }
    ]
}`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var client = NewTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})
