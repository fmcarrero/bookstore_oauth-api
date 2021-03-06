package ports

import (
	"github.com/fmcarrero/bookstore_oauth-api/src/domain/access_token"
)

type AccessTokenRepository interface {
	GetById(string) (*access_token.AccessToken, error)
	Create(token *access_token.AccessToken) error
	UpdateExpirationTime(token access_token.AccessToken) error
}
