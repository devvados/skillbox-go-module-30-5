package handler

import (
	"encoding/json"
	"net/http"
	"skillbox/module30/skillbox-go-module-30-5/pkg/api"
	"skillbox/module30/skillbox-go-module-30-5/pkg/storage"
	"strconv"
	"strings"
)

func GetFriends(s storage.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Формирование ответа
		var status int
		var data []byte
		userId, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/friends/"))

		users, err := s.GetFriends(userId)
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
