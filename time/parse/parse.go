package main

import (
	"fmt"
	"time"
)

func main() {
	ts := "June 18, 1942"

	t, err := time.Parse("January 2, 2006", ts)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Println(t)
	}

	ds := "2700ms"
	d, err := time.ParseDuration(ds)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Println(d)
	}
}
