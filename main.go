package main

import (
	"bytes"
	"fmt"
	"github/jhoncodeu/mailbox-masive-go/src/auth"
	"github/jhoncodeu/mailbox-masive-go/src/core"
	"io"
	"mime/multipart"
	"os"
)

func main() {

	// CpuProfiler
	CpuProfile()

	//code
	//core.ExecAll()
	core.ConvertMboxToNdjson()

	/* 	authUser := auth.Login(config.AuthUser, config.AuthPass)

	   	jsonStr, err := json.Marshal(authUser)
	   	if err != nil {
	   		panic(err)
	   	}

	   	headers := auth.HeaderHttpBasicAuth(authUser.Username, authUser.Password)

	   	// Hacer la solicitud a cualquier Endpoin
	   	resp, err := auth.SendRequest(config.UrlBase+"/user", "GET", jsonStr, headers)
	   	if err != nil {
	   		panic(err)
	   	}
	   	defer resp.Body.Close()

	   	if resp != nil {
	   		fmt.Println("Conexion a DB:", resp.Status)
	   	} else {
	   		fmt.Println("No Conexion a DB!")
	   	} */

	// MemProfiler
	MemProfile()

	// Ejecutar olímpicos de ejemplo
	//loadOlympicsData(config.UrlBase+"/_bulk", headers)
}

func loadOlympicsData(url_base string, headers map[string]string) {
	fileJdjson := "src/example/olympics.ndjson"
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
