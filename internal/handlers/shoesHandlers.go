package handlers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/fbomrl/_Corrida_go/internal/services"
)

var temp = template.Must(template.ParseGlob("../templates/*.html"))

func Shoes(service *services.ShoesService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shoes, err := service.FindAllShoesService()
		if err != nil {
			log.Printf("Erro ao buscar calçados: %v", err)
			http.Error(w, "Erro ao buscar calçados", http.StatusInternalServerError)
			return
		}

		log.Printf("Número de calçados encontrados: %d", len(shoes))
		for i, shoe := range shoes {
			log.Printf("Calçado %d: ID=%d, Name=%s, TotalKm=%.2f",
				i, shoe.Id, shoe.Name, shoe.TotalKm)
		}

		if err := temp.ExecuteTemplate(w, "shoes", shoes); err != nil {
			log.Printf("Erro ao renderizar template: %v", err)
			http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
			return
		}
	}
}
