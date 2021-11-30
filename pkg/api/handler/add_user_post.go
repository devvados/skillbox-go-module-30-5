package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"skillbox/module30/skillbox-go-module-30-5/pkg/api"
	"skillbox/module30/skillbox-go-module-30-5/pkg/storage"
	"skillbox/module30/skillbox-go-module-30-5/pkg/user"
)

func Add(s storage.Adder) http.HandlerFunc {
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
		var u user.User
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		s.Add(&u)

		//Формирование ответа
		data, _ := json.Marshal(api.ResponseDTO{
			Message: "Пользователь создан",
		})
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
