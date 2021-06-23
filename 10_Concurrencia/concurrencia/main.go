package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"
)


func LeerArchivo(nombreArchivo string, canal chan string, cancelarLectura context.CancelFunc) {
	//time.Sleep(3*time.Second)
	file, err := ioutil.ReadFile(nombreArchivo)
	if(err != nil){
		fmt.Println(err)
		cancelarLectura()
	}

	canal <- string(file)
}

func main(){
	const nombreArchivo = "holaMundo.txt"
	// canal sin buffer
	canal := make(chan string)
	ctxVacio := context.Background()
	ctxTimeout, cancelarCtxTimeout := context.WithTimeout(ctxVacio, 2*time.Second)
	go LeerArchivo(nombreArchivo, canal, cancelarCtxTimeout)
	select{
	case contenido := <-canal:
		fmt.Print(contenido)
	case <-ctxTimeout.Done():
		fmt.Println("No se pudo leer el archivo", ctxTimeout.Err())
	}
	
}