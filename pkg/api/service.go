package api

import (
	"skillbox/module30/skillbox-go-module-30-5/pkg/storage"
	"skillbox/module30/skillbox-go-module-30-5/pkg/user"
)

type RequestDTO struct {
	Source int `json:"source_id"`
	Target int `json:"target_id"`
	NewAge int `json:"new age"`
}

type ResponseDTO struct {
	Message string       `json:"message"`
	Items   []*user.User `json:"items"`
}
type ResponseErrorDTO struct {
	Message string `json:"message"`
}

type Service struct {
	Repo *storage.Repo
}
