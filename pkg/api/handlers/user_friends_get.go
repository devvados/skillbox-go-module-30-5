package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"skillbox/module30/skillbox-go-module-30-5/pkg/api"
	"skillbox/module30/skillbox-go-module-30-5/pkg/storage/interfaces"
	"strconv"
	"strings"
)

func GetFriends(repo interfaces.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Формирование ответа
		var status int
		var data []byte
		userId, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/friends/"))

		users, err := repo.GetFriends(context.TODO(), userId)
		if err != nil {
			data, _ = json.Marshal(api.ResponseErrorDTO{
				Message: err.Error(),
			})
			status = http.StatusInternalServerError
		} else {
			data, _ = json.Marshal(api.ResponseDTO{
				Message: "Операция выполнена успешно",
				Items:   users,
			})
			status = http.StatusOK
		}
		w.WriteHeader(status)
		w.Write(data)
	}
}
