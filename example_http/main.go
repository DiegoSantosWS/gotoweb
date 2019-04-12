package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://fliper.app/")
	if err != nil {
		log.Fatal("Erro no GET  da url.", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Erro na leitura do body...", err.Error())
	}

	fmt.Println("Retorno:\r\n", string(body))
}
