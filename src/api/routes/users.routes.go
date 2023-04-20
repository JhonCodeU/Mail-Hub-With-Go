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
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	// Convertir el body de la peticion a json
	jsonStr, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	// Hacer la solicitud a cualquier Endpoin
	resp, err := auth.SendRequest(config.UrlBase+"/user", "POST", jsonStr, headers)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Devolver a la respuesta de la peticion
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, resp.Body)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	// Convertir el body de la peticion a json
	jsonStr, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	// Hacer la solicitud a cualquier Endpoin
	resp, err := auth.SendRequest(config.UrlBase+"/user", "PUT", jsonStr, headers)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Devolver a la respuesta de la peticion
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
