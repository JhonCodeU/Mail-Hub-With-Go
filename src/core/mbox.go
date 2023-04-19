package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Directorio donde se almacenarán los archivos mbox
	root := "src/data/enron_mail_20110402/maildir" // Ruta del directorio raíz
	mboxRoot := "src/data/output/enron.mbox"       // Ruta del directorio de salida

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Si el archivo es un archivo de correo electrónico
		if !info.IsDir() && strings.Contains(info.Name(), "._") == false {

			// Obtener la ruta relativa del archivo de correo electrónico
			relPath, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}

			// Obtener el nombre del archivo mbox correspondiente
			mboxPath := filepath.Join(mboxRoot, relPath)
			mboxPath = strings.TrimSuffix(mboxPath, filepath.Ext(mboxPath)) + ".mbox"

			// Leer el contenido del archivo de correo electrónico
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			// Crear el archivo mbox si no existe
			if _, err := os.Stat(mboxPath); os.IsNotExist(err) {
				err = os.MkdirAll(filepath.Dir(mboxPath), os.ModePerm)
				if err != nil {
					return err
				}

				_, err = os.Create(mboxPath)
				if err != nil {
					return err
				}
			}

			// Agregar el contenido al archivo mbox
			err = ioutil.WriteFile(mboxPath, []byte(content), os.ModePerm)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}
