package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"skillbox/module30/skillbox-go-module-30-5/pkg/storage"
	"skillbox/module30/skillbox-go-module-30-5/pkg/user"
)

//Тип для обработки прочих запросов
type Target struct {
	source int `json: "source_id"`
	target int `json: "target_id"`
}

//Сервис - имеет хранилище
type Service struct {
	St storage.Storage
}

func (s Service) NewService() Service {
	return Service{
		St: storage.Storage.NewStorage(),
	}
}

//Получение всех пользователей из хранилища
func (s *Service) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		response := ""
		for _, user := range s.St {
			response += user.ToString()
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

//Добавление пользователя в хранилище
func (s *Service) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		defer r.Body.Close()

		var u user.User
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		s.St.Add(&u)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Пользователь создан: " + string(u.Id) + " " + u.Name))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

//Удаление пользователя из хранилища
func (s *Service) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		defer r.Body.Close()

		var t Target
		if err := json.Unmarshal(content, &t); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if err := s.St.Delete(t.target); err != nil {
			//Возвращаем ошибку при удалении
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		} else {
			//Возвращаем имя удаленного пользователя
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Пользователь удален " + s.St[t.target].Name))
		}

		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

//Добавление в друзья одного пользователя другому
func (s *Service) MakeFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		response := ""
		for _, user := range s.St {
			response += user.ToString()
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

//Обновление возраста пользователя
func (s *Service) UpdateAge(w http.ResponseWriter, r *http.Request) {

}

//Получение списка друзей пользователя
func (s *Service) GetFriends(w http.ResponseWriter, r *http.Request) {

}
