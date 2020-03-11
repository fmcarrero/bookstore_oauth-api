package app

import (
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/controllers"
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/controllers/access_token"
)

func mapUrls(handler access_token.Handler) {

	router.GET("/ping", controllers.Ping)
	router.GET("/oauth/access_token/:access_token_id", handler.GetById)
}
