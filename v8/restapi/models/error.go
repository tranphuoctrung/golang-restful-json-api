package models

type JsonErr struct {
	Code int `json:"code"`

	Message string `json:"message"`
}
