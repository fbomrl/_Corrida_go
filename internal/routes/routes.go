package routes

import (
	"net/http"

	"github.com/fbomrl/_Corrida_go/internal/handlers"
	"github.com/fbomrl/_Corrida_go/internal/services"
)

func Register(shoesService *services.ShoesService, runningService *services.RunningService) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Home())
	mux.HandleFunc("/shoes", handlers.Shoes(shoesService))
	mux.HandleFunc("/runnings", handlers.Runnings(runningService))
	mux.HandleFunc("/createrunnings", handlers.CreateRunnings(runningService))

	return mux
}
