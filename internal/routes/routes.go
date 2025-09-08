package routes

import (
	"net/http"

	"github.com/fbomrl/_Corrida_go/internal/handlers"
	"github.com/fbomrl/_Corrida_go/internal/services"
)

func Register(service *services.ShoesService) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Index(service))

	return mux
}
