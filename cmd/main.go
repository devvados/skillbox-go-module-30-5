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

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/users", handler.Get(storage))
	r.Post("/create", handler.Add(storage))
	r.Post("/makeFriends", handler.Link(storage))

	http.ListenAndServe(":3333", r)
}
