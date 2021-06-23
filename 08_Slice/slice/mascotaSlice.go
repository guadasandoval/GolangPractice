package slice

import (
	"fmt"
	"sort"
)

/*Del ejercicio de la semana anterior donde una persona tenía una mascota, asumir que
una persona puede tener de 0 a n mascotas.
a. Crear una función que me diga cuantas mascotas tiene una persona.
b. Crear una función que me traiga un slice con los nombres de todas las mascotas
que tenga ordenados alfabéticamente (en lugar de traer el nombre de la mascota
si es que tuviera).
c. Crear una función que me permita agregar una mascota a una persona (en lugar
de la asignación).*/

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
	Mascota  []Mascota
}

//ContarMascotasPersona func
func ContarMascotasPersona(persona *Persona) (int, error) {
	if persona.Mascota != nil {
		return len(persona.Mascota), nil
	}
	
	return 0 , fmt.Errorf("Esta persona no tiene mascota")
}

//NombreMascotasOrdenadosAlfabeticamente func
func NombreMascotasOrdenadosAlfabeticamente(persona *Persona) (*[]string, error) {

	nombresMascotas := make([]string, 0)
	if persona.Mascota != nil {
	for _, mascota := range persona.Mascota {
		nombresMascotas = append(nombresMascotas, mascota.Nombre)
		}
		sort.Strings(nombresMascotas)
		return &nombresMascotas, nil
	}
	
	return nil, fmt.Errorf("Esta persona no tiene mascota")
}

//AgregarMascota func
func AgregarMascota(persona *Persona, mascota Mascota) {

	persona.Mascota = append(persona.Mascota, mascota)
	
}
