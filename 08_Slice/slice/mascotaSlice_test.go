package slice_test

import (
	"resoluciones/08_Slice/slice"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAgregarMascotaAUnaPersona(t *testing.T) {
	personaNueva := slice.Persona{Nombre: "Guadalupe", Apellido: "Sandoval", Edad: 27}
	mascotaA := slice.Mascota{Nombre: "India", Raza: "Siames"}
	mascotaB := slice.Mascota{Nombre: "Tina", Raza: "Callejera"}
	
	slice.AgregarMascota(&personaNueva, mascotaA)
	slice.AgregarMascota(&personaNueva, mascotaB)
	

	cantidadMascotas, _ := slice.ContarMascotasPersona(&personaNueva)

	assert.Equal(t, 2, cantidadMascotas)
}

func TestPersonaSinMascotas(t *testing.T) {
	personaNueva := slice.Persona{Nombre: "Guadalupe", Apellido: "Sandoval", Edad: 27}

	_ , err := slice.ContarMascotasPersona(&personaNueva)

	assert.Error(t, err, "Esta persona no tiene mascota")
}

func TestNombresMascotasOrdenadosAlfabeticamente(t *testing.T) {
	personaNueva := slice.Persona{Nombre: "Guadalupe", Apellido: "Sandoval", Edad: 27}
	mascotaA := slice.Mascota{Nombre: "India", Raza: "Siames"}
	mascotaB := slice.Mascota{Nombre: "Tina", Raza: "Callejera"}
	mascotaC := slice.Mascota{Nombre: "Fito", Raza: "Callejero"}
	mascotaD := slice.Mascota{Nombre: "Luli", Raza: "Cocker"}
	
	slice.AgregarMascota(&personaNueva, mascotaA)
	slice.AgregarMascota(&personaNueva, mascotaB)
	slice.AgregarMascota(&personaNueva, mascotaC)
	slice.AgregarMascota(&personaNueva, mascotaD)

	nombresOrdenados , _ := slice.NombreMascotasOrdenadosAlfabeticamente(&personaNueva)

	assert.Equal(t, nombresOrdenados, &[]string{"Fito","India","Luli","Tina"})
}

func TestNombresMascotasOrdenadosAlfabeticamentePersonaSinMascotas(t *testing.T) {
	personaNueva := slice.Persona{Nombre: "Guadalupe", Apellido: "Sandoval", Edad: 27}

	_ , err := slice.NombreMascotasOrdenadosAlfabeticamente(&personaNueva)

	assert.Error(t, err, "Esta persona no tiene mascota")
	assert.Nil(t, personaNueva.Mascota, "Esta persona no tiene mascota")
}