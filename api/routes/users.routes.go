package routes

import (
	"encoding/json"
	"github/jhoncodeu/mailbox-masive-go/config"
	"github/jhoncodeu/mailbox-masive-go/src/auth"
	"github/jhoncodeu/mailbox-masive-go/src/models"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := auth.SendRequest(config.UrlBase+"/user", "GET", nil, headers)
	if err != nil {
		panic(err)
	}
	// Devolver a la respuesta de la peticion
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, resp.Body)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var user models.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}

	// Hacer la solicitud a cualquier Endpoin
	resp, err := auth.SendRequest(config.UrlBase+"/user", "POST", body, headers)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Devolver a la respuesta de la peticion
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	io.Copy(w, resp.Body)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if userID := chi.URLParam(r, "userID"); userID != "" {
		// Hacer la solicitud a cualquier Endpoin
		resp, err := auth.SendRequest(config.UrlBase+"/user/"+userID, "DELETE", nil, headers)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// Devolver a la respuesta de la peticion
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		io.Copy(w, resp.Body)

	} else {
		w.Write([]byte("UserID no encontrado"))
		return
	}
}
