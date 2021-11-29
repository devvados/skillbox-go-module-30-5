package handler

import (
	"encoding/json"
	"net/http"
	"skillbox/module30/skillbox-go-module-30-5/pkg/api"
	"skillbox/module30/skillbox-go-module-30-5/pkg/storage"
	"skillbox/module30/skillbox-go-module-30-5/pkg/user"
)

func Get(storage *storage.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var items = make([]*user.User, 0)
		for _, val := range storage.Items {
			items = append(items, val)
		}

		//Формирование ответа
		data, _ := json.Marshal(api.ResponseDTO{
			Message: "Операция выполнена успешно",
			Items:   items,
		})
		w.WriteHeader(http.StatusCreated)
		w.Write(data)
	}
}
