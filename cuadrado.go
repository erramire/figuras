package figuras

type Cuadrado struct {
	Lado     float64	
	area      float64
	perimetro float64
}

func (cua *Cuadrado) calcularPerimetro() float64 {
	cua.perimetro=4*cua.Lado
	return cua.perimetro
}

func (cua *Cuadrado) calcularArea() float64 {
	cua.area=cua.Lado * cua.Lado
	return cua.area
}