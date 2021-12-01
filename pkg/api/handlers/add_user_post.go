package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"skillbox/module30/skillbox-go-module-30-5/pkg/api"
	"skillbox/module30/skillbox-go-module-30-5/pkg/model"
	"skillbox/module30/skillbox-go-module-30-5/pkg/storage/interfaces"
)

func Add(repo interfaces.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Чтение запроса
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		//Парсинг запроса
		var u model.User
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		repo.AddUser(&u)

		//Формирование ответа
		data, _ := json.Marshal(api.ResponseDTO{
			Message: "Пользователь создан",
		})
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
