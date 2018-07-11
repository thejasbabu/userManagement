package main

import
(
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/thejasbabu/userManagement/handler"
    "github.com/thejasbabu/userManagement/config"
    "github.com/thejasbabu/userManagement/repository"
    "github.com/kelseyhightower/envconfig"
)

func main() {
    var config config.EnvConfig
    err := envconfig.Process("UM", &config)
    if err != nil {
      log.Fatal(err.Error())
    }
    router := mux.NewRouter()
    repo := repository.NewUserRepository(config.MgDbURL, config.MgDbName, config.MgUserCollection) 
    userHandler := handler.NewUserHandler(*repo)
    router.HandleFunc("/user/{page}", userHandler.GetUsers).Methods("GET")
    router.HandleFunc("/user", userHandler.CreateUser).Methods("POST")
    router.HandleFunc("/user/{id}", userHandler.GetUser).Methods("GET")
    router.HandleFunc("/user/{id}", userHandler.DeleteUser).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":" + config.AppPort, router))
}
