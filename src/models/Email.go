package models

import (
	"bufio"
	"bytes"
	"os"
	"strings"

	"github.com/olivere/ndjson"
)

type Mailbox struct {
	MessageID    string `json:"Message-ID"`
	Date         string `json:"Date"`
	From         string `json:"From"`
	To           string `json:"To"`
	Subject      string `json:"Subject"`
	MimeVersion  string `json:"Mime-Version"`
	ContentType  string `json:"Content-Type"`
	ContentTrans string `json:"Content-Transfer-Encoding"`
	XFrom        string `json:"X-From"`
	XTo          string `json:"X-To"`
	XFolder      string `json:"X-Folder"`
	XOrigin      string `json:"X-Origin"`
	XFileName    string `json:"X-FileName"`
	Body         string `json:"Body"`
}

type table struct {
	Index map[string]interface{} `json:"index"`
}

// leer y comvertir un archivo de texto a json

func ConvertToJdjson(path string) (string, error) {
	// Abrir el archivo de texto
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	var mailboxList []Mailbox

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var fileLines []string

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}
	defer file.Close()

	mailbox := Mailbox{}

	for _, line := range fileLines {
		// Si esta vacio el campo un break
		parts := strings.Split(line, ": ")

		if len(parts) < 2 {
			//Guardar estas partes en el campo body
			mailbox.Body += line
			continue
		}

		//fmt.Println(parts)
		key := parts[0]
		value := parts[1]

		switch key {
		case "Message-ID":
			mailbox.MessageID = value
		case "Date":
			mailbox.Date = value
		case "From":
			mailbox.From = value
		case "To":
			mailbox.To = value
		case "Subject":
			mailbox.Subject = value
		case "Mime-Version":
			mailbox.MimeVersion = value
		case "Content-Type":
			mailbox.ContentType = value
		case "Content-Transfer-Encoding":
			mailbox.ContentTrans = value
		case "X-From":
			mailbox.XFrom = value
		case "X-To":
			mailbox.XTo = value
		case "X-Folder":
			mailbox.XFolder = value
		case "X-Origin":
			mailbox.XOrigin = value
		case "X-FileName":
			mailbox.XFileName = value
		default:
			continue
		}
	}

	mailboxList = append(mailboxList, mailbox)

	// convertir a ndjson
	index := map[string]interface{}{"_index": "mailbox"}
	indexs := []table{{Index: index}}

	var buf bytes.Buffer
	r := ndjson.NewWriter(&buf)

	for _, index := range indexs {
		err := r.Encode(index)
		if err != nil {
			return "", err
		}
	}

	// agregar mailbox
	for _, mailbox := range mailboxList {
		err := r.Encode(mailbox)
		if err != nil {
			return "", err
		}
	}

	return buf.String(), nil
}
