package correo

/*Programa que simule el envío de un mail imprimiendo por pantalla el mensaje
enviando correo...
destinatario: <destinatario>
asunto: <asunto>
mensaje: <mensaje>
Para los distintos casos que se van a plantear a continuación la idea es reutilizar la
función de “simulación de envio de mail”
Escribir el código dentro de un paquete llamado correo y llamar desde el main
a. Crear función no exportada llamada simulaEnvio que reciba por parámetros un
destinatario, asunto y mensaje del tipo string e imprima por pantalla “enviando
correo...” y debajo el destinatario, el asunto y el mensaje. En el caso de que el
parámetro sea vacío imprimir un guión “-”
b. Función exportada que me permita enviar un mail con un destinatario (sin asunto
y sin mensaje)
c. Función exportada que me permita enviar un mail con un destinatario y asunto
(sin mensaje)
d. Función exportada que me permita enviar un mail con un destinatario, asunto y
mensaje*/

func longitudDato(s string) string{
	if len(s) == 0 {
		return "-"
	} else 
		return s
}

func simulaEnvio(destinatario, asunto, mensaje string) string {
	return "enviando correo...\n destinatario:" + longitudDato(destinatario) + "\n asunto:" + longitudDato(asunto) + "\n mensaje:" + longitudDato(mensaje) 
}

//CorreoDestinatario envia un correo solo con destinatario
func CorreoDestinatario(destinatario string) string {
	simulaEnvio(destinatario, "", "")
}

//CorreoDestinatarioYAsunto envia un correo con destinatario y asunto
func CorreoDestinatarioYAsunto(destinatario, asunto string) string {
	simulaEnvio(destinatario, asunto, "")
}

//CorreoCompleto envia un correo completo
func CorreoCompleto(destinatario, asunto, mensaje string) string {
	simulaEnvio(destinatario, asunto, mensaje)
}