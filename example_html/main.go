package main

import (
	"net/http"

	m "github.com/DiegoSantosWS/gotoweb/example_html/models"
)

func main() {
	//rota index
	http.HandleFunc("/", m.IndexPage)
	//rota contato
	http.HandleFunc("/contato", m.ContactPage)
	//fmt.Println("Estou escutando comandante...")

	//Instanciando servidor na porta 1234
	http.ListenAndServe(":1234", nil)
}
