package main

import (
	"fmt"
	"github.com/lucreziabernini/puntoMedio"
)

func main() {
	var p1, p2 pack1.Punto
	p1.X = 55.0
	p2.X = 3.0
	p1.Y = 89.0
	p2.Y = 36.0
	
	fmt.Println(pack1.PuntoMedio(p1, p2))

}
