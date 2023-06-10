package main

import (
	"fmt"

	"rsc.io/quote" // imports thte quote module https://pkg.go.dev/rsc.io/quote/v4
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println(quote.Go()) // using a function from an external module
}

//func main() {
//	fmt.Println(quote.Go())
//}
