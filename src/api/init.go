package api

import (
	"github/jhoncodeu/mailbox-masive-go/src/api/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RunApi() {
	// Crear el router
	router := chi.NewRouter()

	// Rutas
	router.Get("/", routes.HomeHandler)

	http.ListenAndServe(":8080", router)
}
