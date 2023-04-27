package core

import (
	"fmt"
	"github/jhoncodeu/mailbox-masive-go/src/models"
	"os"
	"path/filepath"
	"strings"
)

func ConvertMboxToNdjson() error {
	mboxFolder := "src/data/output/enron.mbox"     // Carpeta donde se encuentran los archivos mbox
	jdJsonFolder := "src/data/output/enron.jdjson" // Carpeta donde se almacenarán los archivos ndjson

	// Comprobar si existe la carpeta enron.json
	if _, err := os.Stat(jdJsonFolder); os.IsNotExist(err) {
		// Crear la carpeta enron.json si no existe
		err := os.Mkdir(jdJsonFolder, 0777)
		if err != nil {
			return err
		}
	}

	// Recorre todos los archivos mbox en la carpeta
	err := filepath.Walk(mboxFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(info.Name()) != ".mbox" {
			return nil
		}

		// Transforma el archivo mbox a ndjson
		ndjson, err := models.ConvertToJdjson(path)
		if err != nil {
			return err
		}

		// Almacenar el archivo ndjson en un archivo nuevo
		ndjsonFilename := strings.TrimSuffix(info.Name(), ".mbox") + ".ndjson"
		ndjsonFile, err := os.Create(filepath.Join(jdJsonFolder, ndjsonFilename))
		if err != nil {
			return err
		}

		defer ndjsonFile.Close()

		_, err = ndjsonFile.WriteString(ndjson)
		if err != nil {
			return err
		}

		fmt.Printf("Archivo %s transformado a %s\n", info.Name(), ndjsonFilename)

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
