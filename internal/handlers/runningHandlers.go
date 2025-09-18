package handlers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/fbomrl/_Corrida_go/internal/model"
	"github.com/fbomrl/_Corrida_go/internal/services"
)

var tempRunning = template.Must(template.ParseGlob("templates/*.html"))

func Runnings(service *services.RunningService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			runnings, err := service.RepoRunning.FindAllRunnings()
			if err != nil {
				log.Printf("Erro ao buscar corridas: %v", err)
				http.Error(w, "Erro ao buscar corridas", http.StatusInternalServerError)
				return
			}
			tempRunning.ExecuteTemplate(w, "runnings", runnings)

		case http.MethodPost:
			name := r.FormValue("name")
			local := r.FormValue("local")
			dateStr := r.FormValue("date")
			distanceStr := r.FormValue("distance")
			hourStr := r.FormValue("hour")
			minuteStr := r.FormValue("minute")
			secondStr := r.FormValue("second")
			eventStr := r.FormValue("event")
			image := r.FormValue("image")
			shoesIdStr := r.FormValue("shoesid")

			date, err := time.Parse("02-01-2006", dateStr)
			if err != nil {
				http.Error(w, "Data inválida", http.StatusBadRequest)
			}
			distance, _ := strconv.ParseFloat(distanceStr, 64)
			hour, _ := strconv.Atoi(hourStr)
			minute, _ := strconv.Atoi(minuteStr)
			second, _ := strconv.Atoi(secondStr)
			event, _ := strconv.ParseBool(eventStr)
			shoesId, _ := strconv.Atoi(shoesIdStr)

			newRunning := model.Running{
				Name:     name,
				Local:    local,
				Date:     date,
				Distance: distance,
				Hour:     hour,
				Minute:   minute,
				Second:   second,
				Event:    event,
				Image:    image,
				ShoesId:  shoesId,
			}

			err = service.CreateRunningService(newRunning)
			if err != nil {
				http.Error(w, "Erro ao criar corrida", http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/runnings", http.StatusSeeOther)
		}

	}
}
