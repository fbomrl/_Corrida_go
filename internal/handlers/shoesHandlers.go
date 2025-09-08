package handlers

import (
	"net/http"
	"text/template"

	"github.com/fbomrl/_Corrida_go/internal/services"
)

var temp = template.Must(template.ParseGlob("../templates/*.html"))

func Index(service *services.ShoesService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shoes, err := service.RepoShoes.FindAllShoes()
		if err != nil {
			http.Error(w, "Erro ao buscar calçados", http.StatusInternalServerError)
			return
		}
		temp.ExecuteTemplate(w, "Index", shoes)
	}
}
