package main

import (
	"testing"
)

func TestCold_GetPrices(t *testing.T) {

	g := Gold{
		Items:  nil,
		Client: client,
	}
	p, err := g.GetPrices()
	if err != nil {
		t.Error(err)
	}
	if p.Price != 1620.545 {
		t.Error("wrong price returned ", p.Price)
	}
}
