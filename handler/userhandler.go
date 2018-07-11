package handler

import (
 "net/http"
 "encoding/json"
 "io/ioutil"
 "io"
 "github.com/thejasbabu/userManagement/domain"
 "github.com/thejasbabu/userManagement/repository"
) 

type UserHandler struct {
    Repository repository.UserRepository
}

func NewUserHandler(repository repository.UserRepository) *UserHandler {
	return &UserHandler{Repository: repository}
}

func (h UserHandler) CreateUser(w http.ResponseWriter, r * http.Request) {
  body, err := ioutil.ReadAll(r.Body)
  if(err != nil) {
    http.Error(w, err.Error(), 500)
    return 
  }
  var user domain.User
  err = json.Unmarshal(body, &user)
  if(err != nil) { 
    http.Error(w, err.Error(), 500)
    return
  }
  err = h.Repository.Write(user)
  if(err != nil) {
    http.Error(w, err.Error(), 500)
    return
  }
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
