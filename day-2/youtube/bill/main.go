package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error){
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill{
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	
	switch opt {
	case "a":
		name, _ := getInput("add an item: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("the price must be a number")
			promptOptions(b)
		}

		b.addItem(name, p)

		fmt.Println("item added", name, p)
		promptOptions(b)	

	case "t":
		tip, _ := getInput("add a tip: ", reader)

		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("the tip must be a number")
			promptOptions(b)
		}

		b.updateTip(t)

		fmt.Println("tip updated succefully..", tip)

		promptOptions(b)	

	case "s":
		b.save()
		fmt.Println("bill saved .. ", b.name)
	default: 
		fmt.Println("not a valid option ...")
		promptOptions(b)	
		
	}
}


func main() {
	mybill := createBill()
	promptOptions(mybill)

	fmt.Println(mybill)
}