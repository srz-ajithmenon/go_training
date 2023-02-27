package handlers

import (
	"Author/global"
	"Author/models"
	"Author/services"
	"context"
	"encoding/json"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	var user *models.User
	json.NewDecoder(r.Body).Decode(&user)

	ctx := r.Context()
	dbsession := ctx.Value("db")
	db := dbsession.(services.Db)
	db.GetUser(user.ID)
	data, err := json.Marshal(db.GetUser(user.ID))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func AddUser(w http.ResponseWriter, r *http.Request) {

	var user *models.User
	json.NewDecoder(r.Body).Decode(&user)

	ctx := r.Context()
	dbsession := ctx.Value("db")
	db := dbsession.(services.Db)
	db.AddUser(user)

	data, err := json.Marshal(db.GetUser(2))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	var user *models.User
	json.NewDecoder(r.Body).Decode(&user)

	ctx := r.Context()
	dbsession := ctx.Value("db")
	db := dbsession.(services.Db)
	db.UpdateUser(1, user)

	data, err := json.Marshal(db.GetUser(1))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	dbsession := ctx.Value("db")
	db := dbsession.(services.Db)
	db.DeleteUser(1)

	data, err := json.Marshal(db.GetUser(1))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)

}

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		dbsession := global.GlobalCtx.Value("db")

		handler.ServeHTTP(writer, request.WithContext(context.WithValue(request.Context(), "db", dbsession.(services.Db))))
	})
}
