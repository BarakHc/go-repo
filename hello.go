package main

import "fmt"
import "rsc.io/quote"

func main() {
	fmt.Printf("hello V23, world\n")
}

func Hello() string {
    return quote.Hello()
}
