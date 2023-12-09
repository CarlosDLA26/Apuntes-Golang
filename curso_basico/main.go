// Se agrega el paquete al que pertenece el archivo
// en este caso como es el archivo donde está la función
// principal el package es main
package main

// Módulo para implementar entradas y salidas por consola
import "fmt"

func main() {

	/* ---------------------------BÁSICO--------------------------- */
	// Declaración de constantes
	const pi, pi2 = 3.1416, 3.15
	const pi3 float64 = 3.1416

	// Imprimir en consola
	fmt.Println("Hola mundo")
	fmt.Println("pi", pi)
	fmt.Printf("pi2 %f", pi2)

	// Al agregar := se dice que la variable no fue declarada
	// antes, es decir, se declara por primera vez, ya después
	// de que se declaró ya se puede usar solo el =
	base := 20
	fmt.Println(base)
	base = 30
	fmt.Println(base)

	// Declaración de variables (me gusta más con el var)
	base_1 := 20
	var altura int = 34
	var area int
	var area_2 = 70
	fmt.Println(base_1, altura, area, area_2)

	// Zero values, Go asigna variables por defecto a tipos de variables
	// específicas cuando no se inciializan con un valor preciso
	var a int  // por defecto a es 0
	var b float64  // por defecto es 0
	var c string  // por defecto ""
	var d bool  // por defecto false
	fmt.Println(a, b, c, d)

	// Ejemplo
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area del cuadrado", areaCuadrado)

	/*---------------------------OPERACIONES ARITMÉTICAS---------------------------*/
	var result = 0
	// Suma
	result = 20 + 10
	fmt.Println("Suma:", result)
	// Resta
	result = 20 + 10
	fmt.Println("Resta:", result)
	// Multiplicación
	result = 20 * 10
	fmt.Println("Multiplicación:", result)
	// Division
	var result_1 float64 = 20 / 10.2
	fmt.Println("División:", result_1)
	// Módulo
	var result_2 int = 21 % 2
	fmt.Println("Módulo:", result_2)
	// Incremental
	var x = 0
	x++
	x += 2
	fmt.Println("x:", x)
	// Decremental
	x--
	x -= 3
	fmt.Println("x:", x)
	// Reto de areas
	// Area rectangulo
	h_rec := 20
	b_rec := 10
	fmt.Println("Area Rectángulo", h_rec * b_rec)
	// Area de círculo
	var r_cir float64 = 10
	fmt.Println("Area círuclo", 3.1416 * r_cir * r_cir)  // Se debe decir que r_cir es float64
	// Area Trapecio
	var B_trap float64 = 20
	var b_trap float64 = 10
	var h_trap float64 = 2
	fmt.Println("Area trapecio", ((b_trap * B_trap) / 2) * h_trap)

	var xxx int = 10
	yyy := 2.5 + 10
	fmt.Println(xxx, yyy)

}
