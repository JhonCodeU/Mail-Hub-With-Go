package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github/jhoncodeu/mailbox-masive-go/src/auth"
	"github/jhoncodeu/mailbox-masive-go/src/core"
	"io"
	"mime/multipart"
	"os"
)

func init() {
	core.Exec()
}

func run() {
	authUser := auth.Login("admin", "Complexpass#123")
	url_base := "http://localhost:4080/api"

	jsonStr, err := json.Marshal(authUser)
	if err != nil {
		panic(err)
	}

	headers := auth.HeaderHttpBasicAuth(authUser.Username, authUser.Password)

	// Hacer la solicitud a cualquier Endpoin
	resp, err := auth.SendRequest(url_base+"/user", "GET", jsonStr, headers)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp != nil {
		fmt.Println("response Status:", resp.Status)
	} else {
		fmt.Println("Response is nil!")
	}

	// Ejecutar olímpicos de ejemplo
	//loadOlympicsData(url_base+"/_bulk", headers)

	// transformar un archivo de texto a json
	/* 	jdjson, err := models.ConvertToJdjson()
	   	if err != nil {
	   		panic(err)
	   	}
	   	resqBulk, err := auth.SendRequest(url_base+"/_bulk", "POST", []byte(jdjson), headers)
	   	if err != nil {
	   		panic(err)
	   	}

	   	defer resqBulk.Body.Close()
	   	fmt.Println("response Status:", resqBulk.Status) */
}

func loadOlympicsData(url_base string, headers map[string]string) {
	fileJdjson := "example/olympics.ndjson"
	file, err := os.Open(fileJdjson)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Crear un formulario para enviar el archivo
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", fileJdjson)
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
