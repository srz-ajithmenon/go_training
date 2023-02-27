package main

import (
	"Author/global"
	"Author/handlers"
	"Author/models"
	"Author/services"
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	dbSession := services.GetNewDbSession()

	dbSession.AddUser(&models.User{
		ID:    1,
		Name:  "Ajith",
		Email: "ajithmenon@gmail.com",
	})

	global.GlobalCtx = context.WithValue(global.GlobalCtx, "db", dbSession)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(handlers.Middleware)
	r.Get("/user", handlers.GetUser)
	r.Post("/adduser", handlers.AddUser)
	r.Put("/updateuser", handlers.UpdateUser)
	r.Delete("/deleteuser", handlers.DeleteUser)

	http.ListenAndServe(":3000", r)
}
