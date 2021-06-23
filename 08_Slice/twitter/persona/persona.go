package persona

import "resoluciones/08_slice/twitter/datos"

//Persona struct
type Persona struct {
	Nombre   string
	Apellido string
	Tweets  []datos.Tweet
}