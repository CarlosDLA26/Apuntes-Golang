// Se agrega el paquete al que pertenece el archivo
// en este caso como es el archivo donde está la función
// principal el package es main
package main

// Módulo para implementar entradas y salidas por consola
import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// Los structs son formas de crear clases dentro de Go
// La sintaxis es `type nombre_struct struct`
type car struct {
	brand string
	year  int
	color string
	type_ string
	owner string
}

func esPalindromo(text string) bool {
	var textReverse string
	text = strings.ReplaceAll(strings.ToLower(text), " ", "")
	// Recorremos el string en reversa
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	return text == textReverse
}

func normalFunction(message string) {
	fmt.Println(message)
}

// a y b son enteros, es una buena práctica
func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

// returns
func returnValue(a int) int {
	return a * 2
}

// returns multiples
// aquí estamos diciendo que se devuelven dos valores
// c y d que son enteros, cuando se hace el return
// se deben devolver en ese mismo orden
func doubleReturn(a int) (c, d int) {
	// a es c y a * 2 es d
	return a, a * 2
}

/*Reto de areas*/
func calculateCircleArea(r float64) float64 {
	return math.Pi * math.Pow(r, 2)
}

func calculateRectangleArea(a, b float64) float64 {
	return a * b
}

func calculateTrapezeArea(B, b, h float64) float64 {
	return ((B + b) * h) / 2
}

func esPar(n int) bool {
	return n%2 == 0
}

// la declaración defer apila las declaracinoes
// en este ejemplo primero se ejecuta Holandas 3
// luego holandas 2 y luego Holandas 1
// porque es como una estructura de datos llamda pila
func declaracion_defer() {
	defer fmt.Println("Holandas 1")
	defer fmt.Println("Holandas 2")
	fmt.Println("Holandas 3")
}

func main() {

	/* ---------------------------BÁSICO--------------------------- */
	fmt.Println("-----------BÁSICO-------------")
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
	fmt.Println("-----------OPERACIONES ARITMÉTICAS-------------")
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
	fmt.Println("-----------TIPOS DE DATOS PRIMITIVOS-------------")
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
	fmt.Println("-----------PAQUETE FMT-------------")
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
	fmt.Printf("Tipo de dato helloMessage: %T\n", helloMessage)

	/*-------------FUNCIONES Y FUNCIONES ANÓNIMAS---------------*/
	fmt.Println("-----------FUNCIONES Y FUNCIONES ANÓNIMAS-------------")
	normalFunction("Holandas")
	tripleArgument(20, 1, "c")
	value := returnValue(20)
	fmt.Println("Value", value)
	value_1, value_2 := doubleReturn(20)
	fmt.Println("Value 1:", value_1, "Value 2:", value_2)
	// otra forma de tomar solo un valor
	value_aux_1, _ := doubleReturn(4)
	fmt.Println("Value aux 1:", value_aux_1)
	// Println también devuelve cosas
	resultadoPrint_1, resultadoPrint_2 := fmt.Println("A ver")
	fmt.Println(resultadoPrint_1, " --- ", resultadoPrint_2)
	// Reto de las areas
	areaCircle := calculateCircleArea(20)
	areaRectangle := calculateRectangleArea(10, 4)
	areaTrapeze := calculateTrapezeArea(20, 10, 5)
	fmt.Println("Area circulo", areaCircle)
	fmt.Println("Area Rectangulo", areaRectangle)
	fmt.Println("Area Trapecio", areaTrapeze)

	/*-------------CICLOS---------------*/
	fmt.Println("-----------CICLOS-------------")
	// Go solo tiene ciclos for :0
	// for condicional
	// i empieza en cero hasta que 9 y se suma de a uno
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// for while
	fmt.Println("---------")
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}
	// for forever ciclo hasta la eternidad
	fmt.Println("---------")
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		// Agregamos el break para que no se ejecute hasta el infinito
		break
	}

	fmt.Println("---------")
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}

	// for each
	fmt.Println("---------")
	listaNumeros := []int{2, 4, 6, 8, 10, 12, 14, 16}
	// normal solo indices
	for i := range listaNumeros {
		fmt.Println(i)
	}
	fmt.Println("---------")
	for _, numero := range listaNumeros {
		fmt.Println(numero)
	}
	fmt.Println("---------")
	// como enumerate de python
	for i, numero := range listaNumeros {
		fmt.Printf("índice %d número %d\n", i, numero)
	}

	/*-------------CONDICIONALES---------------*/
	fmt.Println("-----------CONDICIONALES-------------")
	fmt.Println("and -> &&")
	fmt.Println("or -> ||")
	fmt.Println("not -> !")
	// Convertir texto a número
	value, err := strconv.Atoi("20")
	if err != nil {
		log.Fatal(err) // imprimir el error en caso de que lo haya
	}
	fmt.Println("Value:", value)
	fmt.Println("--------------")
	// número par o impar
	fmt.Println(esPar(20))
	fmt.Println("--------------")
	// user, password := "", ""
	// fmt.Println("Usuario: ")
	// fmt.Scanf("%s\n", &user)
	// fmt.Println("Contraseña: ")
	// fmt.Scanf("%s\n", &password)
	// if user == "Carlos" && password == "123456" {
	// 	fmt.Println("Ingreso correcto")
	// } else {
	// 	fmt.Println("Ingreso NO otorgado")
	// }
	// Switch normal
	modulo := 5 % 2
	switch modulo {
	case 0:
		fmt.Println("Es par")
	case 1:
		fmt.Println("Es impar")
	default:
		fmt.Println("Ni idea que será esa mierda")
	}

	// Switch con la variable dentro
	switch modulo = 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	case 1:
		fmt.Println("Es impar")
	default:
		fmt.Println("Ni idea que será esa mierda")
	}

	// Switch sin condición
	value_switch := 200
	switch {
	case value_switch > 150:
		fmt.Println("Mayor que 150")
	case value < 0:
		fmt.Println("Menor que 0")
	default:
		fmt.Println("Sin condición")
	}

	// defer
	// La palabra clave defer sirve para decirle a Go
	// que antes de que el programa termine, ejecute las
	// instrucciones con ella
	declaracion_defer()

	/*-------------SLICES Y LISTAS---------------*/
	fmt.Println("-----------LISTAS-------------")
	/*Un arreglo contiene un tamaño fijo mientras que los
	slices pueden tener un tamaño variable
	*/
	var arreglo_1 [4]int // arreglo de 4 enteros
	arreglo_1[0] = 2
	arreglo_1[1] = 3
	arreglo_2 := [5]int{0, 1, 2, 3, 4}
	arreglo_3 := [3][2]int{} // [[0 0] [0 0] [0 0]]
	fmt.Println(arreglo_1, len(arreglo_1), cap(arreglo_1))
	fmt.Println(arreglo_2, len(arreglo_2), cap(arreglo_2))
	fmt.Println(arreglo_3, len(arreglo_3), cap(arreglo_3))
	fmt.Println("-----------SLICES-------------")
	/*Los slices parten de un arreglo*/
	arreglo_aux := [5]string{"Carlos", "Ana", "Valentina", "Carolina", "Jessica"}
	slice_1 := arreglo_aux[0:4]
	slice_1 = append(slice_1, "Miguel")
	fmt.Println(slice_1, len(slice_1), cap(slice_1)) // len=5, cap=5
	// Cuando un slice se le agrega más de la capacidad que tiene
	// en este ejemplo capacidad 5, esa capacidad se duplica
	slice_1 = append(slice_1, "Stefany")
	fmt.Println(slice_1, len(slice_1), cap(slice_1)) // len=6, cap=10
	// Otra forma de declarar slices
	var slice_2 []float64
	fmt.Println(slice_2, len(slice_2), cap(slice_2)) // len=0, cap=0
	slice_2 = append(slice_2, 1.65)
	fmt.Println(slice_2, len(slice_2), cap(slice_2)) // len=1, cap=1
	slice_2 = append(slice_2, 2.4)
	fmt.Println(slice_2, len(slice_2), cap(slice_2)) // len=2, cap=2
	slice_2 = append(slice_2, 84)
	fmt.Println(slice_2, len(slice_2), cap(slice_2)) // len=4, cap=4
	// append a nueva lista
	newSlice := []float64{8, 9, 10}
	slice_2 = append(slice_2, newSlice...) // los ... significa que los desempaqueta
	fmt.Println(slice_2, len(slice_2), cap(slice_2))

	/*-------------PALINDROMO CON CICLOS---------------*/
	fmt.Println("-----------PALINDROMO CON CICLOS-------------")
	textoPrueba := "Carlos Loaiza"
	fmt.Println(textoPrueba)
	fmt.Println(esPalindromo("amor a roma"))
	fmt.Println(esPalindromo("ama"))
	fmt.Println(esPalindromo("Ama"))
	fmt.Println(esPalindromo("A m  a"))
	fmt.Println(esPalindromo(textoPrueba))
	// Go no cambia el valor de textoPrueba cuando está en la función
	fmt.Println(textoPrueba)

	/*-------------Llave valor---------------*/
	fmt.Println("-----------Llave valor-------------")
	// En go se llaman maps, en python diccionarios
	map_1 := make(map[string]int) // Declarar un map vacío
	var map_2 map[string]int
	map_3 := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}

	map_1["José"] = 20
	map_1["Carlos"] = 30
	fmt.Println(map_1)
	fmt.Println(map_2)
	fmt.Println(map_3)
	// Recorrer un map
	for llave, valor := range map_1 {
		fmt.Println(llave, valor)
	}

	// tomar datos de diccionario
	edad1 := map_1["José"]
	fmt.Println(edad1)
	// Si se toman datos de una llave inexistente devuelve el cero valor
	// del valor de diccionario, en este caso el cero valor de int,
	// es decir 0
	edad2 := map_1["Messi"]
	fmt.Println(edad2)
	// Verificar si hay una llave en el map
	// La variable ok nos dice true si está, de lo contrario false
	valor_1, esta_1 := map_1["Messi"]
	valor_2, esta_2 := map_1["Carlos"]
	fmt.Println(valor_1, esta_1)
	fmt.Println(valor_2, esta_2)

	/*-------------Structs---------------*/
	fmt.Println("-----------Structs-------------")
	// Instanciar con los datos dados
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)
	// Instanciar sin datos
	var otherCar car
	otherCar.brand = "Chevrolet"
	fmt.Println(otherCar)

}
