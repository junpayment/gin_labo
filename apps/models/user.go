package models

import "time"

type User struct {
  Id uint64 `json:"id" xorm:"'id'"`
  Name string `json:"name" xorm:"'name'"`
  Email string `json:"email" xorm:"'email'"`
  Address string `json:"address" xorm:"'address'"`
  Tel string `json:"tel" xorm:"'tel'"`
  Sex uint `json:"sex" xorm:"'sex'"`
  Created time.Time `json:"created" xorm:"created"`
  Updated time.Time `json:"updated" xorm:"updated"`
}

type UserRepository struct {
}

func NewUser( id uint64, name string, email string, address string, tel string, sex uint) User {
  return User{Id: id, Name: name, Email: email, Address: address, Tel: tel, Sex: sex}
}

func NewUserRepository() UserRepository {
  return UserRepository{}
}

func (m UserRepository) GetById(id uint64) *User {
  user := User{Id: id}
  has, _ := engine.Get(&user)
  if has {
    return &user
  }

  return nil
}

func (user *User)

