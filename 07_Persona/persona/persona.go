package persona

/*
Crear una estructura persona y crear una función que permita cambiar nombre, una que
reciba puntero y otra que no reciba puntero y mostrar la diferencia entre ambas.
*/

//Persona struct
type Persona struct {
	Nombre   string
	Apellido string
}

//ModificaNombre sin puntero
func ModificaNombre(persona Persona, nuevoNombre string) Persona {

	persona.Nombre = nuevoNombre
	return persona
}

//ModificaNombrePuntero con puntero
func ModificaNombrePuntero(persona *Persona, nuevoNombre string) {

	persona.Nombre = nuevoNombre
}
