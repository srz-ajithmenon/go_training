package services

import (
	"Author/models"
)

type DbSession struct {
	Data map[int]*models.User
}

type Db interface {
	GetUser(id int) *models.User
	AddUser(user *models.User)
	UpdateUser(id int, user *models.User)
	DeleteUser(id int)
}

func (db *DbSession) GetUser(id int) *models.User {
	return db.Data[id]
}

func (db *DbSession) AddUser(user *models.User) {
	db.Data[user.ID] = user
}

func (db *DbSession) UpdateUser(id int, user *models.User) {
	user.ID = id
	db.Data[id] = user
}

func (db *DbSession) DeleteUser(id int) {
	// var user *models.User
	db.Data[id] = nil
}

func GetNewDbSession() Db {
	return &DbSession{Data: make(map[int]*models.User)}
}
