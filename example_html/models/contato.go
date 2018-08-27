package models

import (
	"log"
	"net/http"

	ctr "github.com/DiegoSantosWS/gotoweb/example_html/renderHtml"
)

//IndexPage carrega template de index
func ContactPage(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title":  "My frist page using GOLANG",
		"Title2": "Enter in contact to more informations",
	}
	if err := ctr.ModelosContato.ExecuteTemplate(w, "contato.html", data); err != nil {
		http.Error(w, "[CONTENT ERRO] Erro in the execute template",
			http.StatusInternalServerError)
		log.Fatal(err.Error())
	}
}
