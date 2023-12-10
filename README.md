# Go
Es un lenguaje de programación basado de código abierto creado por Google. Sus principales características son la simplicidad, la eficiencia y enfoque en la productividad del desarrollador. Go es un lenguaje de programación compilado.
## Ventajas
* Simplicidad: Go se inspira en lenguajes como C para el rendimiento, pero dejando de lado una sintaxis compleja.
* Eficiencia: Su compilador ofrece código nativo lo que permite un rendimiento excelente. go posee un recolector de basura, por lo que los programadores no tienen que preocuparse por la administración de memoria.
* Concurrencia y paralelismmo: Go proporciona goroutines que son unidades de ejecución livianas y canales para facilitar la comunicación entre ellas, lo cual es útil al momento de crear servicios web o servicios distribuidos.
## Goroutines
Las goroutines son lo que los lenguajes de programación denominan hilos. En el ecosistema de Go existe al menos una principal, llamada goruutina principal y puede controlar las demás goroutines. También existen los canales, que son formas de comunicarnos entre goroutines.
## Programación
### Crear build de Go
Como Go es un lenguaje compilado debe crearse el ejecutable para poder usar el programa. Para ello usamos el comando `go build main.go`, esto creará un archivo `main.exe` podemos ejecutarlo dando doble click o desde la consola con el comando `./main.exe`.
Otra forma de ejecutarlo es usando el comando `go run main.go`, esto hará que se compile el archivo go en una carpeta temporal, se ejecute y después se borre, lo cual ayuda a temas de debug y pruebas. Este método es menos eficiente.
### Básico
Ver la sección BÁSICO y OPERACIONES ARITMÉTICAS del archivo `main.go`. Para la sección de operaciones aritméticas es posible usar este complemento [Cómo hacer cálculos matemáticos en Go con operadores](https://www.digitalocean.com/community/tutorials/how-to-do-math-in-go-with-operators-es).
### Paquete fmt
El paquete fmt es una implementación I/O formateada, básicamente realizar impresiones en consola. Un ejemplo se muestra en el archivo `main.go` en la sección "PAQUETE FMT". también se puede consultar la documentación oficial en [fmt](https://pkg.go.dev/fmt).

