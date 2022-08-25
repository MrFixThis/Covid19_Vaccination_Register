package model

type Administrator struct {
	Person
	PasswordHash string `json:"-"`
}
