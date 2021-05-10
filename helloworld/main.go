package main

import (
	"fmt"

	"github.com/jun1st/calculator"
	"rsc.io/quote"
)

func main() {

	total := calculator.Sum(3, 5)
	println(total)

	println("Version: ", calculator.Version)

	println(quote.Hello())

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	g(0)

	fmt.Println("Program finished successfully!")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic("panick in g() (major)")
	}

	defer fmt.Println("Defer in g()", i)

	fmt.Println("Printing in g()", i)

	g(i + 1)
}
