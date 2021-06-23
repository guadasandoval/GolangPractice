package dosretornos_test

import(
	"testing"
	"resoluciones/03_dosRetornos/dosretornos"
	"github.com/stretchr/testify/assert"
)

func TestSumaCorrecta(t *testing.T) {
	suma, _ := dosretornos.SumaYResta(4,6) 
	assert.Equal(t, 10, suma)
}