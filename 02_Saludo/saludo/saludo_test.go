package saludo_test

import (
	"testing"
	"resoluciones/02_Saludo/saludo"
	"github.com/stretchr/testify/assert"
)

func TestSaludo(t *testing.T){
 saludo := saludo.Saludo("Guada")
 assert.Equal(t, "Hola Guada", saludo)
}