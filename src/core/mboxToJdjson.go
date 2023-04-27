package core

import (
	"fmt"
	"github/jhoncodeu/mailbox-masive-go/src/models"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func ConvertMboxToNdjson() error {
	mboxFolder := "src/data/output/enron.mbox"   // Carpeta donde se encuentran los archivos mbox
	jdJsonFile := "src/data/output/enron.ndjson" // Archivo ndjson de salida
	var wg sync.WaitGroup

	// Recorre todos los archivos mbox en la carpeta
	err := filepath.WalkDir(mboxFolder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(d.Name()) != ".mbox" {
			return nil
		}

		wg.Add(1)
		go func() {
			defer wg.Done()

			// Transforma el archivo mbox a ndjson
			ndjson, err := models.ConvertToJdjson(path)
			if err != nil {
				log.Printf("Error convirtiendo %s a ndjson: %s", path, err.Error())
				return
			}

			// Escribe el registro ndjson en el archivo de salida
			f, err := os.OpenFile(jdJsonFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Printf("Error abriendo el archivo %s: %s", jdJsonFile, err.Error())
				return
			}
			defer f.Close()

			if _, err := f.WriteString(ndjson); err != nil {
				log.Printf("Error escribiendo registro ndjson en el archivo %s: %s", jdJsonFile, err.Error())
				return
			}
		}()

		return nil
	})

	if err != nil {
		return err
	}

	wg.Wait()
	fmt.Printf("Transformación completada con éxito. El resultado se ha almacenado en el archivo %s\n", jdJsonFile)

	return nil
}
