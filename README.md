# Go
Es un lenguaje de programación basado de código abierto creado por Google. Sus principales características son la simplicidad, la eficiencia y enfoque en la productividad del desarrollador. Go es un lenguaje de programación compilado.
## Ventajas
* Simplicidad: Go se inspira en lenguajes como C para el rendimiento, pero dejando de lado una sintaxis compleja.
* Eficiencia: Su compilador ofrece código nativo lo que permite un rendimiento excelente. go posee un recolector de basura, por lo que los programadores no tienen que preocuparse por la administración de memoria.
* Concurrencia y paralelismmo: Go proporciona goroutines que son unidades de ejecución livianas y canales para facilitar la comunicación entre ellas, lo cual es útil al momento de crear servicios web o servicios distribuidos.
## Programación
### Crear build de Go
Como Go es un lenguaje compilado debe crearse el ejecutable para poder usar el programa. Para ello usamos el comando `go build main.go`, esto creará un archivo `main.exe` podemos ejecutarlo dando doble click o desde la consola con el comando `./main.exe`.
Otra forma de ejecutarlo es usando el comando `go run main.go`, esto hará que se compile el archivo go en una carpeta temporal, se ejecute y después se borre, lo caul ayuda a temas de debug y pruebas. Este método es menos eficiente.
