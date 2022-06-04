package main

import "fmt"

func main() {
	mybill := newBill("Jack's bill")
	mybill.addItem("espresso", 1.49)
	mybill.addItem("apple pie", 3.99)
	mybill.addItem("cinnamon roll", 2.99)
	mybill.updateTip(2.53)
	fmt.Println(mybill.format())
}
