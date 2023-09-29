package models

type Response struct {
	Model      GenericModel `json:"model"`
	GenericKey string       `json:"generic_key"`
	StatusCode int          `json:"status_code"`
}
