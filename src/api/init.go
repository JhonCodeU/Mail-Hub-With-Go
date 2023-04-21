package api

import (
	"github/jhoncodeu/mailbox-masive-go/src/api/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func RunApi() {
	// Crear el router
	router := chi.NewRouter()

	// habilitar CORS
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

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
