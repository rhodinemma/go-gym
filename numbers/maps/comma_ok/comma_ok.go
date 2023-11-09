package main

import "fmt"

func main() {
	prices := map[string]int{
		"apple":  5,
		"banana": 8,
		"orange": 4,
	}

	price, ok := prices["banana"]
	if ok {
		fmt.Printf("The price of banana is $%d\n", price)
	} else {
		fmt.Printf("We don't have bananas")
	}

	price, ok = prices["mango"]
	if ok {
		fmt.Printf("The price of mango is $%d\n", price)
	} else {
		fmt.Printf("We don't have mangoes")
	}
}
