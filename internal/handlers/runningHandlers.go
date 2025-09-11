package handlers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/fbomrl/_Corrida_go/internal/services"
)

var tempRunning = template.Must(template.ParseGlob("templates/*.html"))

func Runnings(service *services.RunningService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		runnings, err := service.RepoRunning.FindAllRunnings()
		if err != nil {
			log.Printf("Erro ao buscar corridas: %v", err)
			http.Error(w, "Erro ao buscar corridas", http.StatusInternalServerError)
			return
		}
		tempRunning.ExecuteTemplate(w, "runnings", runnings)
	}
}
