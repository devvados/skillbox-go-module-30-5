package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"skillbox/module30/skillbox-go-module-30-5/pkg/api/handlers"
	"skillbox/module30/skillbox-go-module-30-5/pkg/storage/repositories"
)

func main() {
	//Инциализация сервиса с хранилищем
	storage := repositories.NewLocalRepository()

	//Инициализация маршрутизатора
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//Инициализация хендлов запросов
	r.Get("/users", handlers.GetAll(storage))
	r.Post("/create", handlers.Add(storage))
	r.Post("/makeFriends", handlers.Link(storage))
	r.Delete("/delete", handlers.Delete(storage))
	r.Put("/{user_id}", handlers.Update(storage))
	r.Get("/friends/{user_id}", handlers.GetFriends(storage))

	//Старт сервиса
	http.ListenAndServe(":3333", r)
}
