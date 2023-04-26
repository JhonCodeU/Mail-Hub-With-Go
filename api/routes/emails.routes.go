package routes

import (
	"encoding/json"
	"github/jhoncodeu/mailbox-masive-go/config"
	"github/jhoncodeu/mailbox-masive-go/src/auth"
	"github/jhoncodeu/mailbox-masive-go/src/models"
	"io"
	"net/http"
)

var collection = "mailbox"

func EmailsHandlerSearch(w http.ResponseWriter, r *http.Request) {
	// Obtener el body de la peticion
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// Extraer el valor search, conditional y value del body
	err = json.Unmarshal(body, &models.RequestData)
	if err != nil {
		panic(err)
	}

	// Serializar RequestData a formato JSON
	jsonData, err := json.Marshal(models.RequestData)
	if err != nil {
		panic(err)
	}

	// Hacer la solicitud a _search
	resp, err := auth.SendRequest(config.UrlBase+"/"+collection+"/_search", "POST", jsonData, auth.HeaderHttpBasicAuth(config.AuthUser, config.AuthPass))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Devolver a la respuesta de la peticion
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, resp.Body)
}

func EmailsHandlerSearchAll(w http.ResponseWriter, r *http.Request) {
	// Obtener el body de la peticion
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// Extraer el valor search, conditional y value del body
	err = json.Unmarshal(body, &models.RequestData)
	if err != nil {
		panic(err)
	}

	// Serializar RequestData a formato JSON
	jsonData, err := json.Marshal(models.RequestData)
	if err != nil {
		panic(err)
	}

	// Hacer la solicitud a _search
	resp, err := auth.SendRequest(config.UrlBase+"/"+collection+"/_search", "POST", jsonData, auth.HeaderHttpBasicAuth(config.AuthUser, config.AuthPass))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Devolver a la respuesta de la peticion
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, resp.Body)
}
