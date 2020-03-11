package app

import (
	"github.com/fmcarrero/bookstore_oauth-api/src/application/uses_cases"
	"github.com/fmcarrero/bookstore_oauth-api/src/domain/access_token/service"
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/adapters/repository"
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/adapters/repository/cassandra"
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/controllers/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()
	accessTokenRepository := &repository.AccessTokenCassandraRepository{}
	accessTokenService := &service.Service{
		AccessTokenRepository: accessTokenRepository,
	}
	getAccessTokenUseCase := &uses_cases.GetAccessTokenUseCase{AccessTokenService: accessTokenService}
	handler := access_token.Handler{
		GetAccessTokenUseCase: getAccessTokenUseCase,
	}

	mapUrls(handler)
	_ = router.Run(":8081")

}
