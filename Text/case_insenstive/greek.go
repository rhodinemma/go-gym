package main

import (
	"fmt"
	"strings"
)

// Letter in Greek
type Letter struct {
	Symbol  string
	English string
}

var letters = []Letter{
	{"α", "alpha"},
	{"β", "beta"},
	{"γ", "gamma"},
	{"δ", "delta"},
	{"ε", "epsilon"},
	{"ζ", "zeta"},
	{"η", "eta"},
	{"θ", "theta"},
	{"ι", "iota"},
	{"κ", "kappa"},
	{"λ", "lambda"},
	{"μ", "mu"},
	{"ν", "nu"},
	{"ξ", "xi"},
	{"ο", "omicron"},
	{"π", "pi"},
	{"ρ", "rho"},
	{"σ", "sigma"},
	{"τ", "tau"},
	{"υ", "upsilon"},
	{"Σ", "Epsilon"},
}

// englishFor return the English name for a greek letter
func englishFor(greek string) (string, error) {
	for _, letter := range letters {
		if strings.EqualFold(greek, letter.Symbol) {
			return letter.English, nil
		}
	}

	return "", fmt.Errorf("unknown letter: %s", greek)
}

func main() {
	fmt.Println(englishFor("π"))
	fmt.Println(englishFor("λ"))
	fmt.Println(englishFor("δ"))
}
