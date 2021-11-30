package api

import (
	"skillbox/module30/skillbox-go-module-30-5/pkg/storage"
	"skillbox/module30/skillbox-go-module-30-5/pkg/user"
)

//DTO для объекта запроса
type RequestDTO struct {
	Source int `json:"source_id"`
	Target int `json:"target_id"`
	NewAge int `json:"new age"`
}

//DTO для объекта ответа
type ResponseDTO struct {
	Message string       `json:"message"`
	Items   []*user.User `json:"items"`
}

//DTO для объекта ответа с ошибкой
type ResponseErrorDTO struct {
	Message string `json:"message"`
}

type Service struct {
	Repo *storage.Repo
}
