package validation

import (
	"errors"
	"strings"

	. "github.com/guidomantilla/bookstore_users-api/core/model"
)

func Validate(user *User) error {

	if user.Id == 0 {
		return errors.New("user has an invalid id")
	}

	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if "" == user.Email {
		return errors.New("user has an invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if "" == user.Password {
		return errors.New("user has an invalid password")
	}
	return nil
}
