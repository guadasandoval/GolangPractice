package saludo

import (
	"fmt"
)
/*En el main crear una función que reciba un nombre y devuelva un saludo con la forma
“hola <nombre>”. Si el nombre recibido es vacío debe devolver “hola extraño”. Incluir un
test para cada caso utilizando testify.*/

//Saludo recibe un nombre como parametro
func Saludo(nombre string) string {
if len(nombre) != 0{
	return fmt.Sprintf("Hola %s", nombre)
}
	return fmt.Sprint("Hola extraño")
}