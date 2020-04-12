package repository

import . "github.com/guidomantilla/bookstore_users-api/core/model"

type UserRepository interface {
	Create(user *User) error
	CreateBulk(users []*User) error
	Update(user *User) error
	Delete(id int64) error
	FindById(id int64) (*User, error)
	Find(paramMap map[string][]string) (*[]User, error)
	ExistsById(id int64) (bool, error)
}
