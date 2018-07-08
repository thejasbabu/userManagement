package handler

import (
 "net/http"
 "log"
 "encoding/json"
 "io/ioutil"
 "io"
 "github.com/thejasbabu/userManagement/domain"
) 

type UserHandler struct {}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
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
  log.Println(user)
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
