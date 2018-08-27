package routers

import (
	"net/http"

	"github.com/DiegoSantosWS/gotoweb/projeto_api/models"
	"github.com/gorilla/mux"
)

func Routers() {
	r := mux.NewRouter() //Abre a instrancia para criar as rotas
	r.HandleFunc("/v1/products/", models.Products)
	r.HandleFunc("/v1/products/{id}", models.Products)

	r.HandleFunc("/v1/clients/", models.Products)
	r.HandleFunc("/v1/clients/{id}", models.Products)

	http.ListenAndServe(":1234", r)
}
