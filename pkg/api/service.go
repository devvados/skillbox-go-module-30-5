package api

import (
	"skillbox/module30/skillbox-go-module-30-5/pkg/storage"
	"skillbox/module30/skillbox-go-module-30-5/pkg/user"
)

type RequestDTO struct {
	Source int `json: "source_id"`
	Target int `json: "target_id"`
}

type ResponseDTO struct {
	Error   int          `json: "error"`
	Message string       `json: "message"`
	Items   []*user.User `json: "items"`
}

type Service struct {
	Repo *storage.Repo
}
