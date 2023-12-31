package main

import (
	"fmt"
	"time"
)

func main() {
	chi, err := time.LoadLocation("America/Chicago")
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	chiTime := time.Date(2021, time.February, 28, 19, 30, 0, 0, chi)
	fmt.Println("Chicago: ", chiTime)

	nyc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	nycTime := chiTime.In(nyc)
	fmt.Println("NYC: ", nycTime)
}
