package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) string {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(input)
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	switch opt {
	case "a":
		name := getInput("Item name: ", reader)
		sprice := getInput("Item price: ", reader)
		price, err := strconv.ParseFloat(sprice, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, price)
		fmt.Printf("%v added\n", name)
		promptOptions(b)
	case "t":
		stip := getInput("Enter tip amount ($): ", reader)
		tip, err := strconv.ParseFloat(stip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(tip)
		fmt.Println("tip added - ", tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("you saved the file", b.name)
	default:
		fmt.Println("that was not a valid option...")
		promptOptions(b)
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
}
