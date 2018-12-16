package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // o tipo é definido pelo compilador

	//dessa forma vc declara e já inicializa a desgraçada
	area := PI * math.Pow(raio, 2)
	fmt.Println("Area:", area) //go te obriga a usar a variavel

	const (
		a = 1
		b = 2
	)

	var (
		c = 1
		d = 2
	)

	fmt.Println(a, b, c, d)

	var e, f = 2, false

	fmt.Println(e, f)

}
