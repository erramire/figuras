package figuras

import "fmt"

type Figura interface {
	calcularPerimetro() float64
	calcularArea() float64
}

func CalcularPerimetroFigura(figura Figura) {
	fmt.Println("El Perimetro es :", figura.calcularPerimetro())
}

func CalcularAreaFigura(figura Figura) {
	fmt.Println("El Area es :", figura.calcularArea())
}