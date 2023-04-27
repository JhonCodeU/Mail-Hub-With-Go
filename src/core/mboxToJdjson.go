package core

import (
	"fmt"
	"github/jhoncodeu/mailbox-masive-go/src/models"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ConvertMboxToNdjson() error {
	mboxFolder := "src/data/output/enron.mbox"     // Carpeta donde se encuentran los archivos mbox
	ndjsonFolder := "src/data/output/enron-ndjson" // Carpeta donde se guardarán los archivos ndjson

	// Crear la carpeta ndjson si no existe
	if _, err := os.Stat(ndjsonFolder); os.IsNotExist(err) {
		if err := os.Mkdir(ndjsonFolder, 0755); err != nil {
			return fmt.Errorf("error creando la carpeta %s: %s", ndjsonFolder, err.Error())
		}
	}

	fmt.Println("Iniciando transformación de archivos a ndjson...")
	// Recorre todos los archivos mbox en la carpeta
	err := filepath.WalkDir(mboxFolder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(d.Name()) != ".mbox" {
			return nil
		}

		// Transforma el archivo mbox a ndjson
		ndjson, err := models.ConvertToJdjson(path)
		if err != nil {
			log.Printf("Error convirtiendo %s a ndjson: %s", path, err.Error())
			return nil
		}

		// Obtener el nombre del archivo sin la extensión
		fileName := strings.TrimSuffix(d.Name(), filepath.Ext(d.Name()))

		// Crear el archivo ndjson
		f, err := os.OpenFile(ndjsonFolder+"/"+fileName+".ndjson", os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Printf("Error creando el archivo %s: %s", fileName+".ndjson", err.Error())
			return nil
		}
		defer f.Close()

		// Escribir el registro ndjson en el archivo
		if _, err := f.WriteString(ndjson); err != nil {
			log.Printf("Error escribiendo registro ndjson en el archivo %s: %s", fileName+".ndjson", err.Error())
			return nil
		}

		return nil
	})

	if err != nil {
		return err
	}

	fmt.Printf("Transformación completada con éxito. Los archivos ndjson se han almacenado en la carpeta %s\n", ndjsonFolder)

	return nil
}
