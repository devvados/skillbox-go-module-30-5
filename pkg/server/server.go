package server

import (
	"github.com/go-chi/chi"
	"skillbox/module30/skillbox-go-module-30-5/cmd/utils"
	"skillbox/module30/skillbox-go-module-30-5/pkg/api/routers"
	"skillbox/module30/skillbox-go-module-30-5/pkg/client/mongodb"
	"skillbox/module30/skillbox-go-module-30-5/pkg/storage/repositories"
)

func Init(config utils.Configuration) *chi.Mux {
	client := mongodb.ConnectMongoDb(config.Database.Url)
	repo := repositories.NewMongoDbRepository(client, &config)
	r := routers.NewRouter(repo)

	return r
}
