package impl

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	. "github.com/guidomantilla/bookstore_common-lib/common/exception"
	. "github.com/guidomantilla/bookstore_users-api/core/model"
	. "github.com/guidomantilla/bookstore_users-api/core/service"
	. "github.com/guidomantilla/bookstore_users-api/core/ws/marshalling"
)

type DefaultUserWs struct {
	userService UserService
}

func NewDefaultUserWs(userService UserService) *DefaultUserWs {
	return &DefaultUserWs{
		userService: userService,
	}
}

func (userWs *DefaultUserWs) Create(context *gin.Context) {

	var user User
	if err := context.ShouldBindJSON(&user); err != nil {
		exception := BadRequestException("error unmarshalling request json to object", err)
		context.JSON(exception.Code, exception)
		return
	}

	if exception := userWs.userService.Create(&user); exception != nil {
		context.JSON(exception.Code, exception)
		return
	}

	isPublic := context.GetHeader("X-Public") == "true"
	marshaledUser := MarshallUser(&user, isPublic)

	context.JSON(http.StatusCreated, marshaledUser)
}

func (userWs *DefaultUserWs) Update(context *gin.Context) {

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		exception := BadRequestException("url path has an invalid id", err)
		context.JSON(exception.Code, exception)
		return
	}

	var user User
	if err := context.ShouldBindJSON(&user); err != nil {
		exception := BadRequestException("error unmarshalling request json to object", err)
		context.JSON(exception.Code, exception)
		return
	}

	if exception := userWs.userService.Update(int64(id), &user); exception != nil {
		context.JSON(exception.Code, exception)
		return
	}

	context.JSON(http.StatusOK, map[string]string{})
}

func (userWs *DefaultUserWs) Delete(context *gin.Context) {

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		exception := BadRequestException("url path has an invalid id", err)
		context.JSON(exception.Code, exception)
		return
	}

	if exception := userWs.userService.Delete(int64(id)); exception != nil {
		context.JSON(exception.Code, exception)
		return
	}

	context.JSON(http.StatusOK, map[string]string{})
}

func (userWs *DefaultUserWs) FindById(context *gin.Context) {

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		exception := BadRequestException("url path has an invalid id", err)
		context.JSON(exception.Code, exception)
		return
	}

	user, exception := userWs.userService.FindById(int64(id))
	if exception != nil {
		context.JSON(exception.Code, exception)
		return
	}

	isPublic := context.GetHeader("X-Public") == "true"
	marshaledUser := MarshallUser(user, isPublic)

	context.JSON(http.StatusOK, marshaledUser)
}

func (userWs *DefaultUserWs) Find(context *gin.Context) {

	values := context.Request.URL.Query()

	users, exception := userWs.userService.Find(values)
	if exception != nil {
		context.JSON(exception.Code, exception)
		return
	}

	isPublic := context.GetHeader("X-Public") == "true"
	marshaledUsers := MarshallUsers(users, isPublic)

	context.JSON(http.StatusOK, marshaledUsers)
}
