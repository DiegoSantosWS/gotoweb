package models

import (
	"log"
	"net/http"

	ctr "github.com/DiegoSantosWS/gotoweb/example_html/renderHtml"
)

//IndexPage carrega template de index
func IndexPage(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "My frist page using GOLANG",
		"Agre":  "My frist page using GOLANG",
	}
	if err := ctr.ModelosIndex.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, "[CONTENT ERRO] Erro in the execute template",
			http.StatusInternalServerError)
		log.Fatal(err.Error())
	}
}
