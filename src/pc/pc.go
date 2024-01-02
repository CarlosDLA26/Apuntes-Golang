package pc

import "fmt"

// Función del struct pc
func (pc_var Pc) Ping() {
	fmt.Println(pc_var.Brand, "Pong")
}

// Función para modificar datos en un objeto
// Se agrega el * para decirle que se va a acceder
// mediante punteros, si no se le agrega
// El siguiente código no funcionaría y no duplicaría la RAM
// Podría funcionar para getters y setters
func (pc_var *Pc) DuplicateRAM() {
	pc_var.Ram = pc_var.Ram * 2
}

// Para personalizar las salidas en consola de los objetos
// de una manera similar a __str__ en Python se usan los
// Stringers en Go
func (pc_var Pc) String() string {
	return fmt.Sprintf("PC Brand: %s - RAM: %d - Disk: %d", pc_var.Brand, pc_var.Ram, pc_var.Disk)
}

// Se deben poner la inicial en mayúscula también a las variables
// para que sean públicas, si no no se pueden acceder desde fuera
type Pc struct {
	Ram   int
	Disk  int
	Brand string
}
