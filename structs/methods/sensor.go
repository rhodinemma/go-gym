package main

import "fmt"

// A thermostat measures and controls the temperature
type thermostat struct {
	ID    string
	value float64
}

// Value returns the current temperature in Celsius
func (t *thermostat) Value() float64 {
	return t.value
}

// Set tells the thermostat to set the temperature
func (t *thermostat) Set(value float64) {
	t.value = value
}

// Kind returns the device kind
func (t *thermostat) Kind() string {
	return "thermostat"
}

func main() {
	t := thermostat{"Living Room", 16.2}
	fmt.Printf("%s before: %.2f\n", t.ID, t.Value())
	// Living Room before: 16.20

	t.Set(18)
	fmt.Printf("%s after: %.2f\n", t.ID, t.Value())
	// Living Room after:  18.00
}
