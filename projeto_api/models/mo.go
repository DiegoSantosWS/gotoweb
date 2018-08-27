package models

import (
	"encoding/json"
	"log"
	"net/http"
)

type Foo struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Slug        string  `json:"slug"`
	Price       float64 `json:"price"`
}

const N = 10

func Products(w http.ResponseWriter, r *http.Request) {

	datas2 := make([]Foo, N)
	for i := 0; i < 10; i++ {
		datas2[i] = Foo{ID: i + 1, Title: "test ", Description: "Teste Produto ", Slug: "teste-produto", Price: 12.33}
	}
	j, err := json.Marshal(datas2)
	if err != nil {
		log.Fatal("ERROR: json produtos", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)

}

func ProductsByID(w http.ResponseWriter, r *http.Request) {

}

func Clients(w http.ResponseWriter, r *http.Request) {

}

func ClientsByID(w http.ResponseWriter, r *http.Request) {

}
