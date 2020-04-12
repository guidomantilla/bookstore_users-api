package ws

import (
	"github.com/gin-gonic/gin"
)

type UserWs interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	FindById(context *gin.Context)
	Find(context *gin.Context)
}
