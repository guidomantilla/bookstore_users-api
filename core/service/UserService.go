package service

import (
	. "github.com/guidomantilla/bookstore_users-api/common/exception"
	. "github.com/guidomantilla/bookstore_users-api/core/model"
)

type UserService interface {
	Create(user *User) *Exception
	Update(id int64, user *User) *Exception
	Delete(id int64) *Exception
	FindById(id int64) (*User, *Exception)
	Find(paramMap map[string][]string) (*[]User, *Exception)
}
