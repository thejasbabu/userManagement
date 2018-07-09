package main

import
(
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/thejasbabu/userManagement/handler"
    "github.com/kelseyhightower/envconfig"
)

type Config struct {
    AppPort           int
    MgDbURL           string
    MgDbName          string
    UserMgCollection  string
}

func main() {
    var config Config
    err := envconfig.Process("UM", &config)
    if err != nil {
      log.Fatal(err.Error())
    }
    router := mux.NewRouter()
    userHandler := handler.NewUserHandler()
    router.HandleFunc("/user", userHandler.GetUsers).Methods("GET")
    router.HandleFunc("/user/{id}", userHandler.GetUser).Methods("GET")
    router.HandleFunc("/user", userHandler.CreateUser).Methods("POST")
    router.HandleFunc("/user/{id}", userHandler.DeleteUser).Methods("DELETE")
    log.Fatal(http.ListenAndServe(config.AppPort, router))
}
