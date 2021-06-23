package oficina_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"resoluciones/09_Oficina/oficina"

) 

func TestAdminAtiendeLlamadaIncremetaContador(t *testing.T){
	admin := oficina.Administrativo{Nombre: "Guadalupe", Edad: 27, Sueldo: 500}

	admin.AtenderLlamadas()
	assert.Equal(t, 1, admin.LlamadasAtendidas)
}

func TestSupervisorAumentaSueldoAdmin(t *testing.T){
	admin := oficina.Administrativo{Nombre: "Guadalupe", Edad: 27, Sueldo: 500}
	sup := oficina.Supervisor{Nombre: "Lore", Edad: 30, Sueldo: 100, Administrativos: []*oficina.Administrativo{&admin}}
	sup.AumentarSueldo(&admin, 1)

	assert.Equal(t, float64(1000), admin.Sueldo)
}