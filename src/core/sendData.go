package core

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"sync"
	"time"
)

func SendRequestToZincSearch(baseURL string, pathFolder string) {
	// Obtener lista de archivos en la carpeta
	files, err := ioutil.ReadDir(pathFolder)
	if err != nil {
		panic(err)
	}

	fmt.Println("------------------------------------")
	fmt.Println("Total de archivos: ", len(files))

	// Crear WaitGroup para esperar a todas las goroutines
	var wg sync.WaitGroup
	wg.Add(len(files))

	// Crear canal para manejar errores de goroutines
	errCh := make(chan error)

	// Crear cliente HTTP reutilizable
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	for _, file := range files {
		// Incrementar el contador de WaitGroup para cada goroutine

		// Enviar la solicitud en una goroutine separada
		go func(fileName string) {
			defer wg.Done()

			// Leer el contenido del archivo
			content, err := ioutil.ReadFile(pathFolder + "/" + fileName)
			if err != nil {
				errCh <- err
				return
			}

			// Crear un buffer para el cuerpo de la solicitud
			var buf bytes.Buffer
			w := multipart.NewWriter(&buf)

			// Agregar el archivo al cuerpo de la solicitud
			fw, err := w.CreateFormFile("file", fileName)
			if err != nil {
				errCh <- err
				return
			}
			if _, err := fw.Write(content); err != nil {
				errCh <- err
				return
			}

			// Cerrar el cuerpo de la solicitud
			if err := w.Close(); err != nil {
				errCh <- err
				return
			}

			// Enviar la solicitud
			req, err := http.NewRequest("POST", baseURL+"/_bulk", &buf)
			if err != nil {
				errCh <- err
				return
			}
			req.Header.Set("Content-Type", w.FormDataContentType())
			req.SetBasicAuth("admin", "Complexpass#123")

			resqBulk, err := client.Do(req)
			if err != nil {
				errCh <- err
				return
			}
			defer resqBulk.Body.Close()

			fmt.Println("response Status Bulk:", resqBulk.Status)
		}(file.Name())
	}

	// Esperar a que todas las goroutines hayan terminado
	go func() {
		wg.Wait()
		close(errCh)
	}()

	// Manejar errores de las goroutines
	for err := range errCh {
		panic(err)
	}
}
