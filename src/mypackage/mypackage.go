package mypackage

import "fmt"

// CarPublic car con acceso público
// Generalmente cuando se quiere que algo sea público
// se poner la primera letra en mayúsucla, si se
// quiere como privado, se debe poner en minúscula.
// Cuando se habla de público y privado se habla de que
// es público y privado dentro del mismo paquete
// es por eso que dentro del archivo mypackage/pack.go
// se puede acceder al struct carPrivate, pero en main.go
// no deja importarlo
type CarPublic struct {
	Brand           string
	Year            int
	Color           string
	Type_           string
	Owner           string
	variablePrivada string
}

type carPrivate struct {
	brand           string
	year            int
	color           string
	type_           string
	owner           string
	variablePrivada string
}

func PrintMessage() {
	fmt.Println("Holandas desde un paquete")
}
