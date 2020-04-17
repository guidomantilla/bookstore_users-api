package app

import (
	"github.com/gin-gonic/gin"

	. "github.com/guidomantilla/bookstore_common-lib/common/config"
	. "github.com/guidomantilla/bookstore_common-lib/common/db/impl"
	. "github.com/guidomantilla/bookstore_users-api/core/repository/impl"
	. "github.com/guidomantilla/bookstore_users-api/core/service/impl"
	. "github.com/guidomantilla/bookstore_users-api/core/ws"
	. "github.com/guidomantilla/bookstore_users-api/core/ws/impl"
	. "github.com/guidomantilla/bookstore_users-api/mgmt/ws"
	. "github.com/guidomantilla/bookstore_users-api/mgmt/ws/impl"
)

const (
	BOOKSTORE_USERS_DATASOURCE_URL      = "BOOKSTORE_USERS_DATASOURCE_URL"
	BOOKSTORE_USERS_DATASOURCE_USERNAME = "BOOKSTORE_USERS_DATASOURCE_USERNAME"
	BOOKSTORE_USERS_DATASOURCE_PASSWORD = "BOOKSTORE_USERS_DATASOURCE_PASSWORD"
	BOOKSTORE_USERS_ENVIRONMENT         = "BOOKSTORE_USERS_ENVIRONMENT"
)

func Init() {

	Config()

	managementWs, userWs := Wire()

	ZapLogger.Info("starting the app")
	router := gin.Default()

	router.GET("/mgmt/health", managementWs.Health)
	router.GET("/mgmt/env", managementWs.Env)
	router.GET("/mgmt/info", managementWs.Info)

	router.POST("/api/users", userWs.Create)
	router.GET("/api/users", userWs.Find)
	router.PUT("/api/users/:id", userWs.Update)
	router.DELETE("/api/users/:id", userWs.Delete)
	router.GET("/api/users/:id", userWs.FindById)

	router.Run(":8080")
}

func Config() {

	ConfigProperties()

	Properties[BOOKSTORE_USERS_DATASOURCE_URL] = ":username::password@tcp(localhost:3306)/bookstore-users?charset=utf8"
	Properties[BOOKSTORE_USERS_DATASOURCE_USERNAME] = "root"
	Properties[BOOKSTORE_USERS_DATASOURCE_PASSWORD] = "toolbox123*"
	Properties[BOOKSTORE_USERS_ENVIRONMENT] = "dev"

	ConfigZapLogger(Properties[BOOKSTORE_USERS_ENVIRONMENT])
}

func Wire() (ManagementWs, UserWs) {

	managementWs := NewDefaultManagementWs()

	url := Properties[BOOKSTORE_USERS_DATASOURCE_URL]
	username := Properties[BOOKSTORE_USERS_DATASOURCE_USERNAME]
	password := Properties[BOOKSTORE_USERS_DATASOURCE_PASSWORD]

	mysqlDataSource := NewMysqlDataSource(username, password, url)
	userRepository := NewDefaultUserRepository(mysqlDataSource)
	userService := NewDefaultUserService(userRepository)
	userWs := NewDefaultUserWs(userService)

	return managementWs, userWs
}
