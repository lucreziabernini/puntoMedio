package pack1

type Punto struct {
	X, Y float64
}

func PuntoMedio(p1, p2 Punto) Punto {
	var media Punto
	media.X = (p1.X + p2.X) / 2
	media.Y = (p1.Y + p2.Y) / 2 
	return media 
}