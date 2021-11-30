package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"skillbox/module30/skillbox-go-module-30-5/pkg/api/handler"
	"skillbox/module30/skillbox-go-module-30-5/pkg/storage"
)

func main() {
	//Инциализация сервиса с хранилищем
	storage := storage.New()

	//Инициализация маршрутизатора
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//Инициализация хендлов запросов
	r.Get("/users", handler.GetAll(storage))
	r.Post("/create", handler.Add(storage))
	r.Post("/makeFriends", handler.Link(storage))
	r.Delete("/delete", handler.Delete(storage))
	r.Put("/{user_id}", handler.Update(storage))
	r.Get("/friends/{user_id}", handler.GetFriends(storage))

	//Старт сервиса
	http.ListenAndServe(":3333", r)
}
