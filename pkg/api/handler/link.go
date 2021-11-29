package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"skillbox/module30/skillbox-go-module-30-5/pkg/api"
	"skillbox/module30/skillbox-go-module-30-5/pkg/storage"
)

func Link(s storage.Linker) http.HandlerFunc {
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
		var t api.RequestDTO
		if err := json.Unmarshal(content, &t); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		//Формирование ответа
		var message string
		var status int
		if err := s.Link(t.Source, t.Target); err != nil {
			message = err.Error()
			status = http.StatusInternalServerError
		} else {
			message = "Пользователи добавлены в друзья друг к другу"
			status = http.StatusCreated
		}
		data, _ := json.Marshal(api.ResponseDTO{
			Message: message,
		})
		w.WriteHeader(status)
		w.Write(data)
	}
}
