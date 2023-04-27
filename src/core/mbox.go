package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func Mbox() {
	// Directorio donde se almacenarán los archivos mbox
	root := "src/data/enron_mail_20110402/maildir" // Ruta del directorio raíz
	mboxRoot := "src/data/output/enron.mbox"       // Ruta del directorio de salida

	// Obtener la cantidad total de archivos de correo electrónico
	fileCount := 0
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && !strings.HasPrefix(info.Name(), "._") {
			fileCount++
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Crear la barra de progreso
	bar := progressbar.Default(int64(fileCount))

	// Crear un canal de correos electrónicos para procesarlos concurrentemente
	mboxChan := make(chan string)

	// Iniciar goroutines para procesar los correos electrónicos
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for path := range mboxChan {
				// Obtener la ruta relativa del archivo de correo electrónico
				relPath, err := filepath.Rel(root, path)
				if err != nil {
					fmt.Println(err)
					continue
				}

				// Obtener el nombre del archivo mbox correspondiente
				mboxPath := filepath.Join(mboxRoot, relPath)
				mboxPath = strings.TrimSuffix(mboxPath, filepath.Ext(mboxPath)) + ".mbox"

				// Leer el contenido del archivo de correo electrónico
				content, err := ioutil.ReadFile(path)
				if err != nil {
					fmt.Println(err)
					continue
				}

				// Crear el archivo mbox si no existe
				if _, err := os.Stat(mboxPath); os.IsNotExist(err) {
					err = os.MkdirAll(filepath.Dir(mboxPath), os.ModePerm)
					if err != nil {
						fmt.Println(err)
						continue
					}

					_, err = os.Create(mboxPath)
					if err != nil {
						fmt.Println(err)
						continue
					}
				}

				// Agregar el contenido al archivo mbox
				err = ioutil.WriteFile(mboxPath, []byte(content), os.ModePerm)
				if err != nil {
					fmt.Println(err)
					continue
				}

				// Actualizar la barra de progreso
				bar.Add(1)
			}
		}()
	}

	// Procesar los archivos de correo electrónico y enviarlos al canal
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && !strings.HasPrefix(info.Name(), "._") {
			mboxChan <- path
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	// Cerrar el canal de correos electrónicos
	close(mboxChan)
}
