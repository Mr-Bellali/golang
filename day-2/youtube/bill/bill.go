package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

//make new bill

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

//format bill

func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}
	total += b.tip
	//add tip 
	fs += fmt.Sprintf("%-25v ...$%.2f\n", "tip:", b.tip)


	//total 
	fs += fmt.Sprintf("%-25v ...$%.2f", "total:", total)
	return fs
}

//update tip

func (b *bill) updateTip(tip float64) {
	b.tip = tip 
}

//add item to bill

func (b *bill) addItem(name string , price float64 ) {
	b.items[name] = price 
}

// save bill

func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644 )

	if err != nil {
		panic(err)
	}

	fmt.Println("bill was saved to file")

}