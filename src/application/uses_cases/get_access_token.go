package uses_cases

import (
	"github.com/fmcarrero/bookstore_oauth-api/src/domain/access_token"
	"github.com/fmcarrero/bookstore_oauth-api/src/domain/access_token/service"
)

type GetAccessTokenPort interface {
	Handler(id string) (*access_token.AccessToken, error)
}
type GetAccessTokenUseCase struct {
	AccessTokenService service.AccessTokenServicePort
}

func (h *GetAccessTokenUseCase) Handler(id string) (*access_token.AccessToken, error) {
	accessToken, err := h.AccessTokenService.GetById(id)
	return accessToken, err
}
