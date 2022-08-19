package figuras

const pi = 3.1416

type Circulo struct {
	Radio     float64
	area      float64
	perimetro float64
}

func (circ *Circulo) calcularPerimetro() float64 {
	circ.perimetro =circ.Radio * 2 * pi
	return circ.perimetro
}

func (circ *Circulo) calcularArea() float64 {
	circ.area=circ.Radio * circ.Radio * pi
	return circ.area
}