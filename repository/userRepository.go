package repository

import (
  "gopkg.in/mgo.v2"
  "github.com/thejasbabu/userManagement/domain"
  "errors"
  "log"
)

type UserRepository struct {
    DbUrl           string
    DbName          string
    DbCollection    string
}

func NewUserRepository(dbUrl string, dbName string, dbCollection string) *UserRepository{
  return &UserRepository{DbUrl: dbUrl, DbName: dbName, DbCollection: dbCollection}
}

func (r UserRepository) Write (u domain.User) error {
 session, err := mgo.Dial(r.DbUrl)
 if err != nil {
    log.Fatal(err)
    return errors.New("Cannot talk to database")
  }
  defer session.Close()
  
  collection := session.DB(r.DbName).C(r.DbCollection)
  err = collection.Insert(u) 
  if err != nil {
     log.Fatal(err)
     return errors.New("Cannot insert user data") 
   }
   log.Println("Successfully added the user")
   return nil
}

func (r UserRepository) PaginatedUsers(page int) ([]*domain.User, error) {
  session, err := mgo.Dial(r.DbUrl)
  if err != nil {
     log.Fatal(err)
     return nil, errors.New("Cannot talk to database")
   }
  defer session.Close()
  
  collection := session.DB(r.DbName).C(r.DbCollection)
  var users []*domain.User
  err = collection.Find(nil).Skip((page-1) * 20).Limit(20).All(&users)
  if err != nil {
    log.Fatal(err)
    return nil, errors.New("Cannot retrieve user data") 
  }

  return users, nil
}
