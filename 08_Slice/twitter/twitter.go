package twitter

import (
	"fmt"
	"resoluciones/08_slice/twitter/datos"
	"resoluciones/08_slice/twitter/persona"
)

/*Escribir un programa que permita a una persona generar tweets.
a. Crear una estructura Tweet en un paquete llamado datos.
b. Crear una estructura Persona en un paquete llamado persona.
c. Crear una función que permita a una persona twittear algo, dicho tweet se
agrega a la colección de tweets de la persona.
d. Dadas varias personas encontrar a la persona que más twitteó.
e. Función que devuelva todos los tweets de una persona.
f. Dada una persona, contar la cantidad de veces que se repite cada palabra en
sus tweets. Devolver palabra y cantidad de veces que se repite (maps) (range).
g. Dada una persona, contar las primeras 3 palabras que más se repiten en sus
tweets. Devolver palabra y cantidad de veces que se repite (maps) (range).*/

//AgregarTweetEnPersona func
func AgregarTweetEnPersona(persona *persona.Persona, tweet datos.Tweet) {

	persona.Tweets = append(persona.Tweets, tweet)

}

//PersonaQueMasTwitteo func 
func PersonaQueMasTwitteo(personas *[]persona.Persona) string{
	var personaMasTwitera persona.Persona
	masTweets := 0
	for _, persona := range *personas{
		if(len(persona.Tweets) > masTweets){
			personaMasTwitera = persona
		}
	}
	return personaMasTwitera.Nombre
}

//DevolverTweetsDeUnaPersona func
func DevolverTweetsDeUnaPersona(persona *persona.Persona) ([]string, error) {

	var tweets []string
	if persona.Tweets != nil {
		for _, tweet := range persona.Tweets {
			tweets = append(tweets, tweet.Mensaje)
			return tweets, nil
		}
	}
	return nil, fmt.Errorf("No hay tweets :(")
}
