package app

import (
	"fmt"
	"github.com/fmcarrero/bookstore_oauth-api/src/application/uses_cases"
	"github.com/fmcarrero/bookstore_oauth-api/src/domain/access_token/service"
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/adapters/repository"
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/cassandra"
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/controllers/access_token"
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/rest/users"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

var (
	Router = gin.Default()
)

func StartApplication() {

	_ = godotenv.Load()
	cassandra.Run()
	accessTokenRepository := &repository.AccessTokenCassandraRepository{}
	userRepository := &users.UserRestRepository{}
	accessTokenService := CreateAccessTokenService(accessTokenRepository, userRepository)

	getAccessTokenUseCase := &uses_cases.GetAccessTokenUseCase{AccessTokenService: accessTokenService}
	createAccessTokenUseCase := CreateAccessTokenUseCase(accessTokenService)
	updateSetExpiresUseCase := &uses_cases.UpdateSetExpiresUseCase{AccessTokenService: accessTokenService}
	handler := access_token.Handler{
		GetAccessTokenUseCase:    getAccessTokenUseCase,
		CreateAccessTokenUseCase: createAccessTokenUseCase,
		UpdateTimeExpiresUseCase: updateSetExpiresUseCase,
	}

	MapUrls(handler)
	port := os.Getenv("PORT")
	if port!= "" {
		port = "8080"
	}
	_ = Router.Run(fmt.Sprintf(":%s",port))

}

func CreateAccessTokenService(accessTokenRepository *repository.AccessTokenCassandraRepository, userRepository *users.UserRestRepository) *service.Service {
	return &service.Service{
		AccessTokenRepository: accessTokenRepository,
		UserRepository:        userRepository,
	}
}

func CreateAccessTokenUseCase(accessTokenService *service.Service) *uses_cases.CreateAccessTokenUseCase {
	return &uses_cases.CreateAccessTokenUseCase{AccessTokenService: accessTokenService}
}
