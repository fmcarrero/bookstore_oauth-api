package uses_cases

import (
	"github.com/fmcarrero/bookstore_oauth-api/src/domain/access_token"
	"github.com/fmcarrero/bookstore_oauth-api/src/domain/access_token/service"
	"github.com/fmcarrero/bookstore_oauth-api/src/domain/access_token_request"
)

type CreateAccessTokenPort interface {
	Handler(request access_token_request.AccessTokenRequest) (*access_token.AccessToken, error)
}
type CreateAccessTokenUseCase struct {
	AccessTokenService service.AccessTokenServicePort
}

func (h *CreateAccessTokenUseCase) Handler(request access_token_request.AccessTokenRequest) (*access_token.AccessToken, error) {
	return h.AccessTokenService.Create(request)

}
