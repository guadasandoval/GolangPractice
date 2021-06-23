package mascota

import "fmt"

/*Crear una estructura persona y mascota. Una persona puede tener una mascota (si no
tiene mascota, ese campo será nulo).
a. La persona tiene nombre, apellido, edad y mascota.
b. La mascota tiene nombre y raza (string).
c. Crear una función que me devuelva el nombre de la mascota que tiene una
persona y un error que indique “Esta persona no tiene mascota” cuando no
exista (en este caso, para el nombre devolveremos un string vacío).
d. Crear una función que me permita asignar una mascota a una persona.*/

//Mascota struct
type Mascota struct {
	Nombre string
	Raza   string
}

//Persona struct
type Persona struct {
	Nombre   string
	Apellido string
	Edad     int
	Mascota  *Mascota
}

//NombreMascota func
func NombreMascota(persona Persona) (string, error) {
	if persona.Mascota != nil {
		return persona.Mascota.Nombre, nil
	}
	return "", fmt.Errorf("Esta persona no tiene mascota")

}

//AsignarMascota func
func AsignarMascota(persona *Persona, mascota Mascota) {
	persona.Mascota = &mascota
}
