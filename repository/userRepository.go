package repository

import (
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "github.com/thejasbabu/userManagement/domain"
  "errors"
  "log"
)

type UserRepository struct {}

func NewUserRepository() *UserRepository{
  return &UserRepository{}
}

func (r UserRepository) write (u domain.User) {
 dbUrl := "server1.example.com,server2.example.com"
 session, err := mgo.Dial(dbUrl)
 if err != nil {
    log.Fatal(err)
    return errors.New("Cannot talk to database")
  }
  defer session.Close()
  database := "test"
  userCollection := "people"
  collection := session.DB(database).C(userCollection)
  err := collection.Insert(u) 
  if err != nil {
     log.Fatal(err)
     return errors.New("Cannot insert user data") 
   }
   log.Println("Successfully added the user")
}
