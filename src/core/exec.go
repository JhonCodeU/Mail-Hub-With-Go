package core

import (
	"fmt"
	"github/jhoncodeu/mailbox-masive-go/config"
	"os"
)

func Exec() {

	// Pedir al usuario que seleccione el tipo a comvertir
	fmt.Println("Seleccione el tipo de archivo:")
	fmt.Println("1 - TRANSFORMAR A MBOX")
	fmt.Println("2 - TRANSFORMAR DE MBOX A JDJSON")
	fmt.Println("3 - INDEXAR LOS CORREOS A ZINCSEARCH API")
	fmt.Println("4 - SALIR")

	var option int
	fmt.Scanln(&option)

	// Leer el archivo seleccionado
	switch option {
	case 1:
		// Comprobar si existe la carpeta enron.mbox
		if _, err := os.Stat("src/data/output/enron.mbox"); os.IsNotExist(err) {
			fmt.Println("Comvertiendo los archivos de correos electrónicos a formato mbox...")
			Mbox()
			return
		} else {
			fmt.Println("Ya los correos estan en formato mbox")
			return
		}
	case 2:
		// comprobar si existe la carpeta enron.mbox
		if _, err := os.Stat("src/data/output/enron.mbox"); os.IsNotExist(err) {
			fmt.Println("No existe la carpeta enron.mbox")
			fmt.Println("Tienes que convertir los archivos de correos electrónicos a formato mbox")
			return
		} else {
			// crear la carpeta enron.json si no existe
			pathFolder := "src/data/output/enron.jdjson"
			if _, err := os.Stat(pathFolder); os.IsNotExist(err) {
				os.Mkdir(pathFolder, 0777)
			}
			// validar si hay archivos en la carpeta si hay eliminarlos
			if _, err := os.Stat(pathFolder + "/enron.json"); !os.IsNotExist(err) {
				os.Remove(pathFolder + "/enron.json")
			}

			fmt.Println("Comvertiendo los archivos de correos electrónicos a formato jdjson...")
			ConvertMboxToNdjson()
		}
	case 3:
		// enviar los correos a la api _bulk
		fmt.Println("Enviando los correos a la api _bulk...")
		pathFolder := "src/data/output/enron.jdjson"
		SendRequestToZincsearch(config.UrlBase, pathFolder)

	case 4:
		fmt.Println("Saliendo...")
		break

	default:
		fmt.Println("Opción no válida.")
		return
	}
}
