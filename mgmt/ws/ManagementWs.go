package ws

import "github.com/gin-gonic/gin"

type ManagementWs interface {
	Health(context *gin.Context)
	Env(context *gin.Context)
	Info(context *gin.Context)
}
