package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github/jhoncodeu/mailbox-masive-go/src/auth"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
)

func main() {
	authUser := auth.Login("admin", "Complexpass#123")
	url_base := "http://localhost:4080/api/user"

	jsonStr, err := json.Marshal(authUser)
	if err != nil {
		panic(err)
	}

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(authUser.Username+":"+authUser.Password)),
	}

	// Hacer la solicitud a cualquier Endpoint
	resp, err := auth.SendRequest(url_base, "GET", jsonStr, headers)

	defer resp.Body.Close()
	//fmt.Println("response Status:", resp.Status)

	// Ejecutar olímpicos de ejemplo
	loadOlympicsData(url_base+"/_bulk", headers)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func loadOlympicsData(url_base string, headers map[string]string) {
	file, err := os.Open("olympics.ndjson")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Crear un formulario para enviar el archivo
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "olympics.ndjson")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}
	err = writer.Close()
	if err != nil {
		panic(err)
	}

	// Hacer la solicitud POST para cargar el archivo en la ruta especificada
	resp, err := auth.SendRequest(url_base, "POST", body.Bytes(), headers)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
}
