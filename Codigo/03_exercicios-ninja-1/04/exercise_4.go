package main

import "fmt"

type gopher int

var x gopher

func main() {
	fmt.Printf("%v %T\n", x, x)

	x = 42
	fmt.Println(x)
}

/*

// Outra forma de fazer separando os valores

package main

import "fmt"

type gopher int

var x gopher

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
}

*/
