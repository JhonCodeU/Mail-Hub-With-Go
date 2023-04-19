package auth

import (
	"bytes"
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
