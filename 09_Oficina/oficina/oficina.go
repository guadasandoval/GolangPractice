package oficina

import "fmt"

/*En una oficina hay varios empleados con distintas categorias que realizan distintos tipos
de tareas:
● Administrativo
● Supervisor
● Dueño
De todos los empleados, incluyendo al dueño, me interesa conocer su nombre, edad
(sólo años) y sueldo (float64).
El supervisor además tiene bajo su responsabilidad un conjunto de administrativos a los
que supervisa.
a. Crear estructuras para los tipos de empleado y métodos para cada una de sus
tareas duplicando el código.
b. Todos pueden atender llamadas telefónicas y me interesa llevar un registro de
cuántas atiende cada uno.
c. Sólo un supervisor puede dar un porcentaje de aumento a un
administrativo.
d. El dueño puede felicitar a un administrativo o un supervisor imprimiendo un
mensaje por pantalla que diga "el dueño felicita a <nombre>".
e. En el main realizar todas las pruebas necesarias para demostrar que el código
funciona. Imprimiendo esta sentencia fmt.Printf("%+v%+v %+v \n", administrativo,
supervisor, duenio) al principio y luego al final para visualizar los cambios.*/

type Administrativo struct {
	Nombre string
	Edad int
	Sueldo float64
	LlamadasAtendidas int
}

type Duenio struct {
	Nombre string
	Edad int
	Sueldo float64
	LlamadasAtendidas int
}

type Supervisor struct {
	Nombre string
	Edad int
	Sueldo float64
	Administrativos []*Administrativo
	LlamadasAtendidas int
}

func (admin *Administrativo)AtenderLlamadas(){
	admin.LlamadasAtendidas++
}

func (duenio *Duenio)AtenderLlamadas(){
	duenio.LlamadasAtendidas++
}

func (sup *Supervisor)AtenderLlamadas(){
	sup.LlamadasAtendidas++
}

func(sup *Supervisor) AumentarSueldo(admin *Administrativo, porcentaje float64){
	porcentaje += 1
	admin.Sueldo *= porcentaje
}

func(duenio *Duenio) FelicitarAdmin(admin *Administrativo){
	 fmt.Printf("el dueño felicita a %s", admin.Nombre)
}

func(duenio *Duenio) FelicitarSuper(super *Supervisor){
	fmt.Printf("el dueño felicita a %s", super.Nombre)
}