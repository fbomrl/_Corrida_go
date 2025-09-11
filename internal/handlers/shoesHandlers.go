package handlers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/fbomrl/_Corrida_go/internal/services"
)

var temp = template.Must(template.ParseGlob("../templates/*.html"))

func Index(service *services.ShoesService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shoes, err := service.RepoShoes.FindAllShoes()
		if err != nil {
			log.Printf("Erro ao buscar calçados: %v", err)
			http.Error(w, "Erro ao buscar calçados", http.StatusInternalServerError)
			return
		}
		temp.ExecuteTemplate(w, "Index", shoes)
	}
}
