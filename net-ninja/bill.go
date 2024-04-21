package main

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
