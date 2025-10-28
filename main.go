// Function a set of a instruccion a name is optinal
// Example:
//
//	set name value
//	set value
package main

import (
	"fmt"
	// se indenta solo hacia abajo cuando hay multiples paquetes
	"math"
)

// Namin is optional
func mathExample() float64 {
	float := 3.14
	newValue := math.Round(float)
	// fmt.Println(newValue)
	return newValue
}

func main() {
	fmt.Println(mathExample())
}
