package impl

import (
	"encoding/hex"
	"errors"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	. "github.com/guidomantilla/bookstore_common-lib/common/config"
	. "github.com/guidomantilla/bookstore_common-lib/common/exception"
	. "github.com/guidomantilla/bookstore_users-api/core/model"
	. "github.com/guidomantilla/bookstore_users-api/core/repository"
	. "github.com/guidomantilla/bookstore_users-api/core/service/validation"
)

const (
	CREATE_ERROR_TITLE     = "error creating the user"
	UPDATE_ERROR_TITLE     = "error updating the user"
	DELETE_ERROR_TITLE     = "error deleting the user"
	FIND_BY_ID_ERROR_TITLE = "error finding the user"
	FIND_ERROR_TITLE       = "error finding the users"
)

type DefaultUserService struct {
	userRepository UserRepository
}

func NewDefaultUserService(userRepository UserRepository) *DefaultUserService {
	return &DefaultUserService{
		userRepository: userRepository,
	}
}

func (userService *DefaultUserService) Create(user *User) *Exception {

	if err := Validate(user); err != nil {
		return BadRequestException(CREATE_ERROR_TITLE, err)
	}

	exists, err := userService.userRepository.ExistsById(user.Id)
	if err != nil {
		return InternalServerErrorException(CREATE_ERROR_TITLE, err)
	}

	if exists {
		return BadRequestException(CREATE_ERROR_TITLE, errors.New("user does exist"))
	}

	user.Status = STATUS_ACTIVE
	user.Date = time.Now().Format(time.RFC3339)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return InternalServerErrorException(CREATE_ERROR_TITLE, err)
	}
	user.Password = hex.EncodeToString(hashedPassword)

	if err := userService.userRepository.Create(user); err != nil {
		return InternalServerErrorException(CREATE_ERROR_TITLE, err)
	}

	return nil
}

func (userService *DefaultUserService) Update(id int64, user *User) *Exception {

	if err := Validate(user); err != nil {
		return BadRequestException(UPDATE_ERROR_TITLE, err)
	}

	exists, err := userService.userRepository.ExistsById(user.Id)
	if err != nil {
		return InternalServerErrorException(UPDATE_ERROR_TITLE, err)
	}

	if !exists {
		return NotFoundException(UPDATE_ERROR_TITLE, errors.New("user does not exist"))
	}

	user.Id = id
	if err := userService.userRepository.Update(user); err != nil {
		return InternalServerErrorException(UPDATE_ERROR_TITLE, err)
	}

	return nil
}

func (userService *DefaultUserService) Delete(id int64) *Exception {

	exists, err := userService.userRepository.ExistsById(id)
	if err != nil {
		return InternalServerErrorException(DELETE_ERROR_TITLE, err)
	}

	if !exists {
		return NotFoundException(DELETE_ERROR_TITLE, errors.New("user does not exist"))
	}
	if err := userService.userRepository.Delete(id); err != nil {
		return InternalServerErrorException(DELETE_ERROR_TITLE, err)
	}

	return nil
}

func (userService *DefaultUserService) FindById(id int64) (*User, *Exception) {

	user, err := userService.userRepository.FindById(id)
	if err != nil {

		if err.Error() == "sql: no rows in result set" {
			return nil, NotFoundException(FIND_BY_ID_ERROR_TITLE, errors.New("user does not exist"))
		}

		return nil, InternalServerErrorException(FIND_BY_ID_ERROR_TITLE, err)
	}

	return user, nil
}

func (userService *DefaultUserService) Find(paramMap map[string][]string) (*[]User, *Exception) {

	ZapLogger.Debug("method: Find", zap.Any("params", paramMap))

	users, err := userService.userRepository.Find(paramMap)
	if err != nil {
		return nil, InternalServerErrorException(FIND_ERROR_TITLE, err)
	}

	return users, nil
}
