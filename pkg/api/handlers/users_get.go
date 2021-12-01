package handlers

import (
	"encoding/json"
	"net/http"
	"skillbox/module30/skillbox-go-module-30-5/pkg/api"
	"skillbox/module30/skillbox-go-module-30-5/pkg/storage/interfaces"
)

func GetAll(repo interfaces.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := repo.GetAll()

		//Формирование ответа
		data, _ := json.Marshal(api.ResponseDTO{
			Message: "Операция выполнена успешно",
			Items:   users,
		})
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
