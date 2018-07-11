package handler

import (
 "net/http"
 "encoding/json"
 "io/ioutil"
 "io"
 "log"
 "strconv"
 "github.com/thejasbabu/userManagement/domain"
 "github.com/thejasbabu/userManagement/repository"
 "github.com/gorilla/mux"
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
  log.Println("User created successfully")
  w.WriteHeader(http.StatusCreated)
}

func (h UserHandler) GetUsers(w http.ResponseWriter, r * http.Request) {  
  vars := mux.Vars(r)
  page, err := strconv.Atoi(vars["page"])
  if(err != nil) {
    http.Error(w, err.Error(), 500)
    return 
  }

  users, err := h.Repository.PaginatedUsers(page)
  
  if(err != nil) {
    http.Error(w, err.Error(), 500)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(users)
  log.Println("Users data retrieved successfully")
}

func (h UserHandler) GetUser(w http.ResponseWriter, r * http.Request) {
  io.WriteString(w, "Getting Users...")
}

func (h UserHandler) DeleteUser(w http.ResponseWriter, r * http.Request) {
  io.WriteString(w, "Deleting Users")
}
