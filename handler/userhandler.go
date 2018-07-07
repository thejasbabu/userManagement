package handler

import (
 "net/http"
 "io"
) 

type UserHandler struct {}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h UserHandler) CreateUser(w http.ResponseWriter, r * http.Request) {
  io.WriteString(w, "Creating User...")
}

func (h UserHandler) GetUser(w http.ResponseWriter, r * http.Request) {
  io.WriteString(w, "Getting User...")
}

func (h UserHandler) GetUsers(w http.ResponseWriter, r * http.Request) {
  io.WriteString(w, "Getting Users...")
}

func (h UserHandler) DeleteUser(w http.ResponseWriter, r * http.Request) {
  io.WriteString(w, "Deleting Users")
}
