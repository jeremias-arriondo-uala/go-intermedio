package models

type Request struct {
	GenericKey   string   `json:"generic_key"`
	GenericArray []string `json:"objects"`
}
