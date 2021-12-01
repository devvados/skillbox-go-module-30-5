package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"skillbox/module30/skillbox-go-module-30-5/pkg/api/handlers"
	"skillbox/module30/skillbox-go-module-30-5/pkg/storage/repositories"
)

func main() {
	port := ":3333"

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

	fmt.Printf("Started listening on port %s", port)
	//Старт сервиса
	http.ListenAndServe(port, r)
}
