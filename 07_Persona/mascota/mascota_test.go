package mascota_test

import (
	"resoluciones/07_Persona/mascota"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNombreMascota(t *testing.T) {
	persona := mascota.Persona{Nombre: "Guada", Apellido: "Sandoval", Edad: 27}
	mascotaNueva := mascota.Mascota{Nombre: "TINA", Raza: "Callejera"}

	mascota.AsignarMascota(&persona, mascotaNueva)

	nombreMascota, error := mascota.NombreMascota(persona)

	assert.Equal(t, "TINA", nombreMascota)
	assert.NoError(t, error)
	assert.NotNil(t, persona.Mascota)
}

func TestPersonaSinMascota(t *testing.T) {
	persona := mascota.Persona{Nombre: "Guada", Apellido: "Sandoval", Edad: 27}

	nombreMascota, err := mascota.NombreMascota(persona)

	assert.Equal(t, "", nombreMascota)
	assert.Error(t, err, "Esta persona no tiene mascota")
}
