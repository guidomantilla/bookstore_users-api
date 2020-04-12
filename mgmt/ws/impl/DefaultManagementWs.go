package impl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultManagementWs struct {
}

func NewDefaultManagementWs() *DefaultManagementWs {
	return &DefaultManagementWs{}
}

func (managementWs *DefaultManagementWs) Health(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "It's Alive!!",
	})
}

func (managementWs *DefaultManagementWs) Env(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "env",
	})
}

func (managementWs *DefaultManagementWs) Info(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "info",
	})
}
