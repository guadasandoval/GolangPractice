package password_test

import(
	"testing"
	"github.com/stretchr/testify/assert"
	"resoluciones/05_Password/password"

)
func TestVerificarPasswordLongitud(t *testing.T) {
	password := password.VerificarPassword("hola123")
	assert.Equal(t, false, password)
}

func TestVerificarPasswordNumeros(t *testing.T) {
	password := password.VerificarPassword("holamundo")
	assert.Equal(t, false, password)
}

func TestVerificarPasswordMayuscula(t *testing.T) {
	password := password.VerificarPassword("holamundo123")
	assert.Equal(t, false, password)
}

func TestVerificarPasswordCorrecta(t *testing.T) {
	password := password.VerificarPassword("holaMundo123")
	assert.Equal(t, true, password)
}
