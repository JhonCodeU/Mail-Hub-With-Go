package core

import (
	"bytes"
	"fmt"
	"github/jhoncodeu/mailbox-masive-go/src/auth"
	"io/ioutil"
	"mime/multipart"
	"sync"
)

func SendRequestToZincsearch(base_url string, pathFolder string) {
	// Obtener lista de archivos en la carpeta
	files, err := ioutil.ReadDir(pathFolder)
	if err != nil {
		panic(err)
	}

	fmt.Println("------------------------------------")
	fmt.Println("Total de archivos: ", len(files))

	// Crear WaitGroup para esperar a todas las goroutines
	var wg sync.WaitGroup

	for _, file := range files {
		// Leer el contenido del archivo
		content, err := ioutil.ReadFile(pathFolder + "/" + file.Name())
		if err != nil {
			panic(err)
		}

		// Incrementar el contador de WaitGroup para cada goroutine
		wg.Add(1)

		// Enviar la solicitud en una goroutine separada
		go func(content []byte) {
			defer wg.Done()

			headers := auth.HeaderHttpBasicAuth("admin", "Complexpass#123")

			// Crear un buffer para el cuerpo de la solicitud
			var buf bytes.Buffer
			w := multipart.NewWriter(&buf)

			// Agregar el archivo al cuerpo de la solicitud
			fw, err := w.CreateFormFile("file", file.Name())
			if err != nil {
				panic(err)
			}
			if _, err := fw.Write(content); err != nil {
				panic(err)
			}

			// Cerrar el cuerpo de la solicitud
			if err := w.Close(); err != nil {
				panic(err)
			}

			// Enviar la solicitud
			resqBulk, err := auth.SendRequest(base_url+"/_bulk", "POST", buf.Bytes(), headers)
			if err != nil {
				panic(err)
			}
			defer resqBulk.Body.Close()

			fmt.Println("response Status Bulk:", resqBulk.Status)
		}(content)
	}

	// Esperar a que todas las goroutines hayan terminado
	wg.Wait()
}
