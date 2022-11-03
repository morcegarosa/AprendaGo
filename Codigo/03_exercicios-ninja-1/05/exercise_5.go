package main

import "fmt"

type gopher int

var x gopher

var y int

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n\n", x)

	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%T\n", y)
}
