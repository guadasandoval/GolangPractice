package password

import (
	
	"regexp"
)
/*Crear un método que permita validar la seguridad de una contraseña retornando un
error si no cumple con algún criterio especificado. La contraseña debe tener al menos 8
caracteres, un número y una mayúscula.
TIP: usar regex para simplificar la detección de las reestricciones ;)
regexp.MatchString(".*\\d.*", password)
regexp.MatchString(".*[A-Z].*", password)
a. Test que falle si tiene menos de 8 caracteres
b. Test que falle si no tiene un número
c. Test que falle si no tiene una mayúscula
d. Test que pase la validación y que la contraseña tenga más de un número
*/

//func MatchString(pattern string, s string) (matched bool, err error)

//VerificarPassword func 
func VerificarPassword(pass string) bool{

	//8 caracteres
	if (len(pass) < 8) {
		return false
	}
	
	//contiene algun caracter numerico
 	contiene, _ := regexp.MatchString(".*\\d.*", pass)
	if(!contiene){
		return false
	}
	
	//contiene algun caracter en mayuscula
	contiene, _ = regexp.MatchString(".*[A-Z].*", pass)
	if(!contiene){
		return false
	}

	return true
}