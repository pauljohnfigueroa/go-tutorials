package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{
			"ketchup": 100.23,
			"chicken": 200.44,
		},
		tip: 0,
	}

	return b
}

func (b bill) format() string {
	fs := "Bill breakdown \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-15v ... $%v\n", k+":", v)
		total += v
	}
	fs += fmt.Sprintf("%-15v ... $%0.2f \n", "total:", total)
	return fs
}
