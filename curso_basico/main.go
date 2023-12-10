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
	var a int     // por defecto a es 0
	var b float64 // por defecto es 0
	var c string  // por defecto ""
	var d bool    // por defecto false
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
	fmt.Println("Area Rectángulo", h_rec*b_rec)
	// Area de círculo
	var r_cir float64 = 10
	fmt.Println("Area círuclo", 3.1416*r_cir*r_cir) // Se debe decir que r_cir es float64
	// Area Trapecio
	var B_trap float64 = 20
	var b_trap float64 = 10
	var h_trap float64 = 2
	fmt.Println("Area trapecio", ((b_trap*B_trap)/2)*h_trap)

	var xxx int = 10
	yyy := 2.5 + 10
	fmt.Println(xxx, yyy)

	/*---------------------------TIPOS DE DATOS PRIMITIVOS---------------------------*/
	// Numeros enteros
	// int = Depende del OS (32 o 64 bits)
	// int8 = 8bits = -128 a 127
	// int16 = 16bits = -2^15 a 2^15-1
	// int32 = 32bits = -2^31 a 2^31-1
	// int64 = 64bits = -2^63 a 2^63-1

	// Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	// u significa unsigned
	// uint = Depende del OS (32 o 64 bits)
	// uint8 = 8bits = 0 a 127
	// uint16 = 16bits = 0 a 2^16-1
	// uint32 = 32bits = 0 a 2^32-1
	// uint64 = 64bits = 0 a 2^64-1

	// Números decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	// Textos y booleanos
	// string = ""
	// bool = true or false

	// numeros complejos
	// Complex64 = Real e Imaginario float32
	// Complex128 = Real e Imaginario float64
	// Ejemplo: c := 10 + 8i
	complejo := 10 + 8i
	fmt.Println(complejo)

	/*---------------------------PAQUETE FMT---------------------------*/
	helloMessage := "Hello"
	worldMessage := "world"
	// Println imprime y al final imrpime un salto de linea
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf imprime variables dentro del string
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	// %s para strings
	// %d enteros
	// %f flotantes
	// %v tipo de dato no determinado
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf no imprimen nada, guardan un string en una variable
	// es una forma de crear strings, lo mismo con Sprintln
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Saber el tipo de dato
	fmt.Printf("%T\n", complejo)
	fmt.Printf("Tipo de dato helloMessage: %T", helloMessage)


}
