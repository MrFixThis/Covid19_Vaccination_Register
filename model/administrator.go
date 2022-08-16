package model

type Administrator struct {
	Person
	Password string `json:"password"`
}
