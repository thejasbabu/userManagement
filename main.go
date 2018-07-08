package main

import
(
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/thejasbabu/userManagement/handler"
)

func main() {
    router := mux.NewRouter()
    userHandler := handler.NewUserHandler()
    router.HandleFunc("/user", userHandler.GetUsers).Methods("GET")
    router.HandleFunc("/user/{id}", userHandler.GetUser).Methods("GET")
    router.HandleFunc("/user", userHandler.CreateUser).Methods("POST")
    router.HandleFunc("/user/{id}", userHandler.DeleteUser).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router))
}
