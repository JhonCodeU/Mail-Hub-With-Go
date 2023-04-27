package models

type User struct {
	Id       string `json:"_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
