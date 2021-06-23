package persona_test

import (
	"resoluciones/07_Persona/persona"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModificoNombrePersona(t *testing.T) {

	personaOriginal := persona.Persona{Nombre: "Guadalupe", Apellido: "Sandoval"}

	personaNuevoNombre := persona.ModificaNombre(personaOriginal, "Lucia")

	assert.Equal(t, "Lucia", personaNuevoNombre.Nombre)
	assert.Equal(t, "Sandoval", personaNuevoNombre.Apellido)
}

func TestModificoNombrePersonaPuntero(t *testing.T) {

	personaOriginal := persona.Persona{Nombre: "Guadalupe", Apellido: "Sandoval"}

	persona.ModificaNombrePuntero(&personaOriginal, "Lucia")

	assert.Equal(t, "Lucia", personaOriginal.Nombre)
	assert.Equal(t, "Sandoval", personaOriginal.Apellido)
}
