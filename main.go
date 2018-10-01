package main

import (
	"biogo"
	"fmt"
	"strings"
)

func main() {
	parsed, _ := biogo.ParseFasta("ctv_total.txt")
	gcskews, _ := biogo.CalculateGCSkew(parsed, 50, 5, 1000)

	for _, g := range gcskews {
		results := strings.Join(g.Values, ";")
		fmt.Println(results + ";1")
	}

	parsed2, _ := biogo.ParseFasta("si_total.txt")
	gcskews2, _ := biogo.CalculateGCSkew(parsed2, 50, 5, 1000)

	for _, g := range gcskews2 {
		results := strings.Join(g.Values, ";")
		fmt.Println(results + ";0")
	}
}
