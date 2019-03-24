package controllers

import (
	"dev/models"
)

// User is
type User struct {
}

// NweUser ...
func NewUser() User {
	return User{}
}

// Get ...
func (c User) Get(n int) interface{} {
	repo := models.NewUserRepository()
	user := repo.GetByID(n)
	return user
}