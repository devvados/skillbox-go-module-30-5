package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"skillbox/module30/skillbox-go-module-30-5/pkg/api"
	"skillbox/module30/skillbox-go-module-30-5/pkg/storage"
)

func Delete(s storage.Deleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Чтение запроса
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			data, _ := json.Marshal(api.ResponseErrorDTO{
				Message: err.Error(),
			})
			w.Write(data)
			return
		}
		defer r.Body.Close()

		//Парсинг запроса
		var t api.RequestDTO
		if err := json.Unmarshal(content, &t); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			data, _ := json.Marshal(api.ResponseErrorDTO{
				Message: err.Error(),
			})
			w.Write(data)
			return
		}

		//Формирование ответа
		var status int
		var data []byte
		if err := s.DeleteUser(t.Source); err != nil {
			data, _ = json.Marshal(api.ResponseErrorDTO{
				Message: err.Error(),
			})
			status = http.StatusInternalServerError
		} else {
			data, _ = json.Marshal(api.ResponseDTO{
				Message: "Пользователь удален",
			})
			status = http.StatusOK
		}
		w.WriteHeader(status)
		w.Write(data)
	}
}
