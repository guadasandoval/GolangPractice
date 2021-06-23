package twitter_test

import (
	"fmt"
	"resoluciones/08_Slice/twitter"
	"resoluciones/08_slice/twitter/datos"
	"resoluciones/08_slice/twitter/persona"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAgregarTweetEnPersona(t *testing.T) {
	personaNueva := persona.Persona{Nombre: "Guadalupe", Apellido: "Sandoval"}
	tweetA := datos.Tweet{Mensaje: "Hola estoy en la capacitacion de Go"}
	tweetB := datos.Tweet{Mensaje: "Esto es un test"}
	
	
	twitter.AgregarTweetEnPersona(&personaNueva, tweetA)
	twitter.AgregarTweetEnPersona(&personaNueva, tweetB)
	
	tweetsPersona, _ := twitter.DevolverTweetsDeUnaPersona(&personaNueva)

	//Este falla
	assert.Equal(t, []string{"Hola estoy en la capacitacion de Go","Esto es un test"}, tweetsPersona)
}

func TestPersonaQueNoTieneTweets(t *testing.T){
	personaNueva := persona.Persona{Nombre: "Guadalupe", Apellido: "Sandoval"}
	
	tweetsPersona, err := twitter.DevolverTweetsDeUnaPersona(&personaNueva)
	assert.Error(t, err, "No hay tweets :(")
	assert.Nil(t,tweetsPersona, nil)
}

func TestPersonaQueMasTwitteo(t *testing.T){
	var personasTwitteras []persona.Persona

	personaA := persona.Persona{Nombre: "Guadalupe", Apellido: "Sandoval"}
	personaB := persona.Persona{Nombre: "Jesi", Apellido: "Vazquez"}
	personaC := persona.Persona{Nombre: "Vale", Apellido: "Benitez"}
 	
	personasTwitteras = append(personasTwitteras, personaA)
	personasTwitteras = append(personasTwitteras, personaB)
	personasTwitteras = append(personasTwitteras, personaC)

	tweetA := datos.Tweet{Mensaje: "Hola :)"}
	tweetB := datos.Tweet{Mensaje: "Esta para tomar una cerveza"}

	tweetC := datos.Tweet{Mensaje: "Hola!"}
	tweetD := datos.Tweet{Mensaje: "Amo los stickers de gatitxs"}
	tweetE := datos.Tweet{Mensaje: "Siempre hay que documentar"}
	tweetF := datos.Tweet{Mensaje: "ERR"}

	tweetG := datos.Tweet{Mensaje: "HOLA HOLA HOLA"}
	tweetH := datos.Tweet{Mensaje: "Vi una serie que esta buenisima"}
	tweetI := datos.Tweet{Mensaje: "cuidemosno del covid"}
	
	
	twitter.AgregarTweetEnPersona(&personaA, tweetA)
	twitter.AgregarTweetEnPersona(&personaA, tweetB)

	twitter.AgregarTweetEnPersona(&personaB, tweetC)
	twitter.AgregarTweetEnPersona(&personaB, tweetD)
	twitter.AgregarTweetEnPersona(&personaB, tweetE)
	twitter.AgregarTweetEnPersona(&personaB, tweetF)

	twitter.AgregarTweetEnPersona(&personaC, tweetG)
	twitter.AgregarTweetEnPersona(&personaC, tweetH)
	twitter.AgregarTweetEnPersona(&personaC, tweetI)
	
	personaMasTwitera := twitter.PersonaQueMasTwitteo(&personasTwitteras)
	fmt.Println(personaMasTwitera)
	assert.Equal(t, "Jesi", personaMasTwitera)
}
