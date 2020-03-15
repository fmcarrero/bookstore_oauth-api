package app

import (
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/controllers/access_token"
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/controllers/ping"
)

func MapUrls(handler access_token.Handler) {

	Router.GET("/ping", ping.Ping)
	Router.GET("/oauth/access_token/:access_token_id", handler.GetById)
	Router.POST("/oauth/access_token", handler.Create)
}
