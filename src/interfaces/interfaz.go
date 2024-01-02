package interfaces

import (
	"fmt"
	"math"
)

/*Una interfaz es un conjunto de métodos con nombre en común que
usan varios structs, la forma de crear solo un punto de entrada,
son las interfaces
*/
/*La interfaz Figuras2D es la que se usa para calcular el area de las figuras
En reemplazo de las funciones de Area(), separadas lo que hace la interfaz es
tomar ambas funciones y ponerlas como un único punto de entrada
*/
type Figuras2D interface {
	area() float64
}

type Cuadrado struct {
	Lado float64
}

type Rectangulo struct {
	Base   float64
	Altura float64
}

type Circulo struct {
	Radio float64
}

/*En este caso podemos dejar las funciones area como privadas
porque no es necesario exponerlos, para calcular el area de las figuras
se debe usar el método que usa la interfaz CalcularArea(),
este si debe ser público
*/
func (c Cuadrado) area() float64 {
	return c.Lado * c.Lado
}

func (r Rectangulo) area() float64 {
	return r.Base * r.Altura
}

func (c Circulo) area() float64 {
	return math.Pi * c.Radio * c.Radio
}

/*Esto solo ejecuta la función Area a partir de la interfaz Figuras2D
*/
func CalcularArea(f Figuras2D) float64 {
	fmt.Println("Area:", f.area())
	return f.area()
}
