package service

import (
	"github.com/fmcarrero/bookstore_oauth-api/src/domain/access_token"
	"github.com/fmcarrero/bookstore_oauth-api/src/domain/access_token_request"
	"github.com/fmcarrero/bookstore_utils-go/logger"
)
import "github.com/fmcarrero/bookstore_oauth-api/src/domain/ports"

type AccessTokenServicePort interface {
	GetById(string) (*access_token.AccessToken, error)
	Create(request access_token_request.AccessTokenRequest) (*access_token.AccessToken, error)
	UpdateExpirationTime(token access_token.AccessToken) error
}

type Service struct {
	AccessTokenRepository ports.AccessTokenRepository
	UserRepository        ports.UserRepository
}

func (s *Service) GetById(id string) (*access_token.AccessToken, error) {
	accessToken, err := s.AccessTokenRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return accessToken, err
}
func (s *Service) Create(request access_token_request.AccessTokenRequest) (*access_token.AccessToken, error) {

	if err := request.Validate(); err != nil {
		return nil, err
	}
	user, errLogin := s.UserRepository.LoginUser(request.Username, request.Password)
	if errLogin != nil {
		return nil, errLogin
	}

	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()
	if err := s.AccessTokenRepository.Create(&at); err != nil {
		logger.Error(err.Error(), err)
		return nil, err
	}
	return &at, nil

}

func (s *Service) UpdateExpirationTime(token access_token.AccessToken) error {
	if err := token.Validate(); err != nil {
		return err
	}
	return s.AccessTokenRepository.UpdateExpirationTime(token)

}
