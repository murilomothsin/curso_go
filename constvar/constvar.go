package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	// forma reduzida de criar um var
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	var e, f bool = true, false
	g, h, i := 2, false, "epa!"

	// obrigatório usar todas as variáveis criadas!
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
