package handlers

import (
	"net/http"
	"text/template"

	"github.com/alex-evans/allowance_app/src/domain"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(
		"src/templates/index.html",
		"src/templates/head.html",
		"src/templates/header.html",
		"src/templates/footer.html",
	)

	kidData := domain.KidData{
		Name:          "Kennedy",
		CurrentAmount: 10,
	}

	t.Execute(w, kidData)
}
