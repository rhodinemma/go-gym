package main

import (
	"fmt"
	"time"
)

func main() {
	lennon := time.Date(1940, time.October, 9, 18, 30, 0, 0, time.UTC)
	fmt.Println(lennon)

	fmt.Println(lennon.Format("2006-01-02"))
	fmt.Println(lennon.Format("Mon, Jan 02"))

	fmt.Println(lennon.Format(time.RFC3339Nano))

	d := 3500 * time.Millisecond
	fmt.Println(d)
}
