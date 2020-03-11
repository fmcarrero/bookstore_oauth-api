package app

import (
	"github.com/fmcarrero/bookstore_oauth-api/src/application/uses_cases"
	"github.com/fmcarrero/bookstore_oauth-api/src/domain/access_token/service"
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/adapters/repository"
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/controllers/access_token"
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/rest/users"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	accessTokenRepository := &repository.AccessTokenCassandraRepository{}
	userRepository := &users.UserRestRepository{}
	accessTokenService := &service.Service{
		AccessTokenRepository: accessTokenRepository,
		UserRepository:        userRepository,
	}
	getAccessTokenUseCase := &uses_cases.GetAccessTokenUseCase{AccessTokenService: accessTokenService}
	createAccessTokenUseCase := &uses_cases.CreateAccessTokenUseCase{AccessTokenService: accessTokenService}
	updateSetExpiresUseCase := &uses_cases.UpdateSetExpiresUseCase{AccessTokenService: accessTokenService}
	handler := access_token.Handler{
		GetAccessTokenUseCase:    getAccessTokenUseCase,
		CreateAccessTokenUseCase: createAccessTokenUseCase,
		UpdateTimeExpiresUseCase: updateSetExpiresUseCase,
	}

	mapUrls(handler)
	_ = router.Run(":8081")

}
