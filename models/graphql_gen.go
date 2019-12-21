// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type Account struct {
	ID         string   `json:"id"`
	Email      string   `json:"email"`
	EmailToken string   `json:"emailToken"`
	Payload    *Payload `json:"payload"`
}

type Asset struct {
	Name string `json:"name"`
}

type Payload struct {
	Data string `json:"data"`
	Salt string `json:"salt"`
}
