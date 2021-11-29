package main

import (
	"net/http"

	"skillbox/module30/skillbox-go-module-30-5/pkg/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	//Инциализация сервиса с хранилищем
	srv := service.Service.NewService()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//Инициализация обработчиков для ручек
	r.Get("/friends/{name}", srv.GetFriends)
	r.Get("/users", srv.GetAll)
	r.Post("/create", srv.Create)
	r.Post("/updateAge", srv.UpdateAge)
	r.Post("/makeFriends", srv.MakeFriends)
	r.Delete("/delete", srv.Delete)

	//Старт сервиса
	http.ListenAndServe(":3333", r)
}
