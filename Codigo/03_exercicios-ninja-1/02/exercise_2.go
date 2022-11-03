// Retorna o valor e o tipo atribuido as variaveis
package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", z, z)

}

/* Retorna somente o valor atribuido as variaveis
package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)

}

*/
