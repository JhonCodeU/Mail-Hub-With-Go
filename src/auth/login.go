package auth

import (
	"bytes"
	"encoding/base64"
	"net/http"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(Username string, Password string) Auth {
	return Auth{
		Username: Username,
		Password: Password,
	}
}

func HeaderHttpBasicAuth(Username string, Password string) map[string]string {
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(Username+":"+Password)),
	}

	return headers
}

func SendRequest(url string, method string, body []byte, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}
	// Añadir las cabeceras a la solicitud
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Crear el cliente HTTP y hacer la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
