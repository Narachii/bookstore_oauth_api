package app

import (
	"github.com/Narachii/bookstore_oauth_api/src/clients/cassandra"
	"github.com/Narachii/bookstore_oauth_api/src/domain/access_token"
	"github.com/Narachii/bookstore_oauth_api/src/http"
	"github.com/Narachii/bookstore_oauth_api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session := cassandra.GetSession()
	session.Close()
	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
