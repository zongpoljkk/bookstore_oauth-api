package app

import (
	"github.com/gin-gonic/gin"
	"github.com/zongpoljkk/bookstore_oauth-api/src/domain/access_token"
	"github.com/zongpoljkk/bookstore_oauth-api/src/http"
	"github.com/zongpoljkk/bookstore_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(access_token.NewService(db.NewRepo()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run(":8080")
}
