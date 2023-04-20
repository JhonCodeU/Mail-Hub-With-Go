package models

type User struct {
	id       string `json:"_id"`
	name     string `json:"name"`
	password string `json:"password"`
	role     string `json:"role"`
}
