package service

import (
	"github.com/fmcarrero/bookstore_oauth-api/src/domain/access_token"
)
import "github.com/fmcarrero/bookstore_oauth-api/src/domain/ports"

type Port interface {
	GetById(string) (*access_token.AccessToken, error)
}

type Service struct {
	AccessTokenRepository ports.AccessTokenRepository
}

func (s *Service) GetById(id string) (*access_token.AccessToken, error) {
	accessToken, err := s.AccessTokenRepository.GetById(id)
	if err != nil {

	}
	return accessToken, err
}
