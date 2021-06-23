package prueba

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func pruebaClienteHttp(URL string, ID int) {

	response, err := http.Get("http://front/site/seleccion/datos-candidatx/1236012360")
	if err != nil {
		fmt.Println("error Get", err)
		return
	}

	fmt.Println("status code", response.StatusCode)
	if response.Body == nil {
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error readAll", err)
		return
	}

	fmt.Println("body", string(body))
}