package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99, "espresso": 1.49},
		tip:   0,
	}
	return b
}

// format the bill
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0
	// list items
	for key, value := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", key+":", value)
		total += value
	}
	// total
	fs += fmt.Sprintf("\n%-25v ...$%0.2f", "total:", total)
	return fs
}
