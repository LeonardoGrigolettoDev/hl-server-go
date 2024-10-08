package main

import (
	"fmt"
	"net/http"

	"github.com/LeonardoGrigolettoDev/fly-esp-server-go/config"
	services "github.com/LeonardoGrigolettoDev/fly-esp-server-go/services/device"
	"github.com/go-chi/chi/v5"
)

func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}
	r := chi.NewRouter()
	r.Post("/device/", services.Create)
	r.Put("/device/{id}", services.Update)
	r.Delete("/device/{id}", services.Delete)
	r.Get("/device/", services.List)
	r.Get("/device/{id}", services.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
