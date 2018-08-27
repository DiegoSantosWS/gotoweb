package renderHtml

import "html/template"

var (
	ModelosIndex = template.Must(template.ParseFiles("view/index.html"))

	ModelosContato = template.Must(template.ParseFiles("view/contato.html"))
)
