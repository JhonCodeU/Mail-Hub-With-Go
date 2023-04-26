package main

import (
	"expvar"
	"fmt"
	"github/jhoncodeu/mailbox-masive-go/api/routes"
	"net/http"
	"net/http/pprof"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	// Crear el router
	router := chi.NewRouter()

	// Habilitar CORS
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

	// Inicializar profiling
	router.HandleFunc("/pprof", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, r.RequestURI+"/", http.StatusMovedPermanently)
	})

	// Rutas de profiling
	router.HandleFunc("/pprof/*", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)
	router.HandleFunc("/vars", expVars)

	router.Handle("/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/pprof/mutex", pprof.Handler("mutex"))
	router.Handle("/pprof/heap", pprof.Handler("heap"))
	router.Handle("/pprof/block", pprof.Handler("block"))
	router.Handle("/pprof/allocs", pprof.Handler("allocs"))

	http.ListenAndServe(":8080", router)
}

func expVars(w http.ResponseWriter, r *http.Request) {
	first := true
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\n")
	expvar.Do(func(kv expvar.KeyValue) {
		if !first {
			fmt.Fprintf(w, ",\n")
		}
		first = false
		fmt.Fprintf(w, "%q: %s", kv.Key, kv.Value)
	})
	fmt.Fprintf(w, "\n}\n")
}
