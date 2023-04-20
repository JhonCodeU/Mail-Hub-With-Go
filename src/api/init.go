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
	// Rutas para los usuarios
	router.Get("/users", routes.GetUsersHandler)
	router.Post("/user", routes.CreateUserHandler)
	router.Put("/user", routes.CreateUserHandler)
	router.Route("/user/{userID}", func(r chi.Router) {
		r.Delete("/", routes.DeleteUserHandler)
	})

	// Rutas para los correos
	router.Post("/emails/search", routes.EmailsHandlerSearch)
	router.Post("/emails/search_all", routes.EmailsHandlerSearchAll)

	http.ListenAndServe(":8080", router)
}
